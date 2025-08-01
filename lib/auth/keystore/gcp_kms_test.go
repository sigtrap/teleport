/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package keystore

import (
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	mathrandv1 "math/rand" //nolint:depguard // only used for deterministic output
	"net"
	"strings"
	"sync"
	"testing"
	"time"

	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/gravitational/trace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ssh"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"github.com/gravitational/teleport/api/types"
	apiutils "github.com/gravitational/teleport/api/utils"
	"github.com/gravitational/teleport/api/utils/grpc/interceptors"
	"github.com/gravitational/teleport/api/utils/keys"
	"github.com/gravitational/teleport/lib/auth/keystore/internal/faketime"
	"github.com/gravitational/teleport/lib/cryptosuites"
	"github.com/gravitational/teleport/lib/jwt"
	"github.com/gravitational/teleport/lib/service/servicecfg"
	"github.com/gravitational/teleport/lib/services"
	"github.com/gravitational/teleport/lib/tlsca"
)

// fakeGCPKMSServer is a GRPC service implementation which fakes the real GCP
// KMS service, to be used in unit tests.
type fakeGCPKMSServer struct {
	kmspb.UnimplementedKeyManagementServiceServer

	mu          sync.RWMutex
	keyVersions map[string]*keyState

	initialKeyState kmspb.CryptoKeyVersion_CryptoKeyVersionState
}

func newFakeGCPKMSServer(opts ...fakeGCPKMSServerOption) *fakeGCPKMSServer {
	f := &fakeGCPKMSServer{
		keyVersions: make(map[string]*keyState),
	}
	for _, opt := range opts {
		opt(f)
	}
	if f.initialKeyState == kmspb.CryptoKeyVersion_CRYPTO_KEY_VERSION_STATE_UNSPECIFIED {
		f.initialKeyState = kmspb.CryptoKeyVersion_ENABLED
	}
	return f
}

type fakeGCPKMSServerOption func(*fakeGCPKMSServer)

func withInitialKeyState(state kmspb.CryptoKeyVersion_CryptoKeyVersionState) fakeGCPKMSServerOption {
	return func(f *fakeGCPKMSServer) {
		f.initialKeyState = state
	}
}

type keyState struct {
	pem              []byte
	cryptoKey        *kmspb.CryptoKey
	cryptoKeyVersion *kmspb.CryptoKeyVersion
}

func (f *fakeGCPKMSServer) CreateCryptoKey(ctx context.Context, req *kmspb.CreateCryptoKeyRequest) (*kmspb.CryptoKey, error) {
	keyName := req.Parent + "/cryptoKeys/" + req.CryptoKeyId
	keyVersionName := keyName + "/cryptoKeyVersions/1"

	cryptoKey := apiutils.CloneProtoMsg(req.CryptoKey)
	cryptoKey.Name = keyName

	cryptoKeyVersion := &kmspb.CryptoKeyVersion{
		Name:            keyVersionName,
		State:           f.initialKeyState,
		ProtectionLevel: cryptoKey.VersionTemplate.ProtectionLevel,
		Algorithm:       cryptoKey.VersionTemplate.Algorithm,
	}

	var pem []byte
	switch cryptoKey.VersionTemplate.Algorithm {
	case kmspb.CryptoKeyVersion_RSA_SIGN_PKCS1_2048_SHA256, kmspb.CryptoKeyVersion_RSA_SIGN_PKCS1_4096_SHA512:
		pem = testRSA2048PrivateKeyPEM
	case kmspb.CryptoKeyVersion_RSA_DECRYPT_OAEP_4096_SHA256:
		pem = testRSA4096PrivateKeyPEM
	case kmspb.CryptoKeyVersion_EC_SIGN_P256_SHA256:
		signer, err := cryptosuites.GenerateKeyWithAlgorithm(cryptosuites.ECDSAP256)
		if err != nil {
			return nil, trace.Wrap(err)
		}
		pem, err = keys.MarshalPrivateKey(signer)
		if err != nil {
			return nil, trace.Wrap(err)
		}
	default:
		return nil, trace.BadParameter("unsupported algorithm %v", cryptoKey.VersionTemplate.Algorithm)
	}

	f.mu.Lock()
	f.keyVersions[keyVersionName] = &keyState{
		pem:              pem,
		cryptoKey:        cryptoKey,
		cryptoKeyVersion: cryptoKeyVersion,
	}
	f.mu.Unlock()

	return cryptoKey, nil
}

func (f *fakeGCPKMSServer) GetPublicKey(ctx context.Context, req *kmspb.GetPublicKeyRequest) (*kmspb.PublicKey, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	keyState, ok := f.keyVersions[req.Name]
	if !ok {
		return nil, trace.NotFound("no such public key")
	}
	if keyState.cryptoKeyVersion.State != kmspb.CryptoKeyVersion_ENABLED {
		return nil, trace.BadParameter("cannot fetch public key, state has value %s", keyState.cryptoKeyVersion.State)
	}

	signer, err := keys.ParsePrivateKey(keyState.pem)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	// Not using [keys.MarshalPublicKey] here because GCP always encodes RSA keys in PKIX format, not PKCS#1.
	pubKeyDER, err := x509.MarshalPKIXPublicKey(signer.Public())
	if err != nil {
		return nil, trace.Wrap(err)
	}
	pubKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  keys.PKIXPublicKeyType,
		Bytes: pubKeyDER,
	})

	return &kmspb.PublicKey{
		Pem: string(pubKeyPEM),
	}, nil
}

func (f *fakeGCPKMSServer) AsymmetricSign(ctx context.Context, req *kmspb.AsymmetricSignRequest) (*kmspb.AsymmetricSignResponse, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	keyState, ok := f.keyVersions[req.Name]
	if !ok {
		return nil, trace.NotFound("no such key")
	}
	if keyState.cryptoKeyVersion.State != kmspb.CryptoKeyVersion_ENABLED {
		return nil, trace.BadParameter("cannot fetch key, state has value %s", keyState.cryptoKeyVersion.State)
	}

	signer, err := keys.ParsePrivateKey(keyState.pem)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	var digest []byte
	var alg crypto.Hash
	switch typedDigest := req.Digest.Digest.(type) {
	case *kmspb.Digest_Sha256:
		switch keyState.cryptoKeyVersion.Algorithm {
		case kmspb.CryptoKeyVersion_RSA_SIGN_PKCS1_2048_SHA256, kmspb.CryptoKeyVersion_EC_SIGN_P256_SHA256:
		default:
			return nil, trace.BadParameter("requested key uses algorithm %s which cannot handle a 256 bit digest",
				keyState.cryptoKeyVersion.Algorithm)
		}
		digest = typedDigest.Sha256
		alg = crypto.SHA256
	case *kmspb.Digest_Sha512:
		if keyState.cryptoKeyVersion.Algorithm != kmspb.CryptoKeyVersion_RSA_SIGN_PKCS1_4096_SHA512 {
			return nil, trace.BadParameter(
				"requested key uses algorithm %s which cannot handle a 512 bit digest",
				keyState.cryptoKeyVersion.Algorithm)
		}
		digest = typedDigest.Sha512
		alg = crypto.SHA512
	default:
		return nil, trace.BadParameter("unsupported digest type %T", typedDigest)
	}

	testRand := mathrandv1.New(mathrandv1.NewSource(0))
	sig, err := signer.Sign(testRand, digest, alg)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	resp := &kmspb.AsymmetricSignResponse{
		Signature: sig,
	}

	return resp, nil
}

func (f *fakeGCPKMSServer) AsymmetricDecrypt(ctx context.Context, req *kmspb.AsymmetricDecryptRequest) (*kmspb.AsymmetricDecryptResponse, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	keyState, ok := f.keyVersions[req.Name]
	if !ok {
		return nil, trace.NotFound("no such key")
	}
	if keyState.cryptoKeyVersion.State != kmspb.CryptoKeyVersion_ENABLED {
		return nil, trace.BadParameter("cannot fetch key, state has value %s", keyState.cryptoKeyVersion.State)
	}

	signer, err := keys.ParsePrivateKey(keyState.pem)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	decrypter, ok := signer.Signer.(crypto.Decrypter)
	if !ok {
		return nil, trace.Errorf("private key is not a valid decrypter")
	}

	testRand := mathrandv1.New(mathrandv1.NewSource(0))
	plaintext, err := decrypter.Decrypt(testRand, req.Ciphertext, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return nil, trace.Wrap(err)
	}

	resp := &kmspb.AsymmetricDecryptResponse{
		Plaintext: plaintext,
	}

	return resp, nil
}

func (f *fakeGCPKMSServer) ListCryptoKeys(ctx context.Context, req *kmspb.ListCryptoKeysRequest) (*kmspb.ListCryptoKeysResponse, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	var cryptoKeys []*kmspb.CryptoKey
	for keyVersionName, keyState := range f.keyVersions {
		if !strings.HasPrefix(keyVersionName, req.Parent) {
			continue
		}
		cryptoKeys = append(cryptoKeys, keyState.cryptoKey)
	}
	resp := &kmspb.ListCryptoKeysResponse{
		CryptoKeys: cryptoKeys,
		TotalSize:  int32(len(cryptoKeys)),
	}
	return resp, nil
}

func (f *fakeGCPKMSServer) DestroyCryptoKeyVersion(ctx context.Context, req *kmspb.DestroyCryptoKeyVersionRequest) (*kmspb.CryptoKeyVersion, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	keyState, ok := f.keyVersions[req.Name]
	if !ok {
		return nil, trace.NotFound("no such key")
	}

	keyState.cryptoKeyVersion.State = kmspb.CryptoKeyVersion_DESTROY_SCHEDULED
	return keyState.cryptoKeyVersion, nil
}

// deleteKey is a test helper to delete a key by the raw ID which would be
// stored in the teleport CA.
func (f *fakeGCPKMSServer) deleteKey(raw []byte) error {
	keyID, err := parseGCPKMSKeyID(raw)
	if err != nil {
		return trace.Wrap(err)
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	delete(f.keyVersions, keyID.keyVersionName)
	return nil
}

// activateAllKeys is a test helper to set the state of all currently held keys
// to ENABLED.
func (f *fakeGCPKMSServer) activateAllKeys() {
	f.mu.Lock()
	defer f.mu.Unlock()
	for _, keyVersion := range f.keyVersions {
		keyVersion.cryptoKeyVersion.State = kmspb.CryptoKeyVersion_ENABLED
	}
}

// newTestGRPCServer returns a basic gRPC server with no services set and some
// helpful interceptors set up.
func newTestGRPCServer() *grpc.Server {
	// Set up some helpful interceptors so that we return compliant error types.
	return grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.GRPCServerUnaryErrorInterceptor),
		grpc.StreamInterceptor(interceptors.GRPCServerStreamErrorInterceptor),
	)
}

type contextDialer func(context.Context, string) (net.Conn, error)

// newTestGCPKMSService creates a new gRPC server and sets it up with an
// in-memory (bufconn) listener and a fakeGCPKMSServer ready to respond to
// requests. Returns a pointer to the fakeGCPKMSServer for the test to
// manipulate, and a dialer function which can be used to connect a gRPC client
// to the bufconn.
func newTestGCPKMSService(t *testing.T, opts ...fakeGCPKMSServerOption) (*fakeGCPKMSServer, contextDialer) {
	grpcServer := newTestGRPCServer()

	fakeKMSServer := newFakeGCPKMSServer(opts...)

	kmspb.RegisterKeyManagementServiceServer(grpcServer, fakeKMSServer)

	l := bufconn.Listen(1024 * 4)
	bufConnDialer := func(ctx context.Context, _ string) (net.Conn, error) {
		return l.DialContext(ctx)
	}

	grpcServeErr := make(chan error)
	go func() {
		grpcServeErr <- grpcServer.Serve(l)
	}()
	t.Cleanup(func() {
		require.NoError(t, <-grpcServeErr)
	})
	t.Cleanup(grpcServer.Stop)

	return fakeKMSServer, bufConnDialer
}

// newTestGPCKMSClient accepts a dial function and creates a test gRPC client
// connected over a bufconn to a KeyManagement service.
func newTestGCPKMSClient(t *testing.T, dialer contextDialer) *kms.KeyManagementClient {
	ctx := context.Background()

	conn, err := grpc.Dial("bufconn",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	kmsClient, err := kms.NewKeyManagementClient(ctx, option.WithGRPCConn(conn))
	require.NoError(t, err)

	return kmsClient
}

// TestGCPKMSKeystore tests GCP KMS keystore operation in the presence of
// delays, timeouts, and errors specific to GCP KMS.
func TestGCPKMSKeystore(t *testing.T) {
	clusterName, err := services.NewClusterNameWithRandomID(types.ClusterNameSpecV2{
		ClusterName: "test-cluster",
	})
	require.NoError(t, err)
	for _, tc := range []struct {
		desc                  string
		initialKeyState       kmspb.CryptoKeyVersion_CryptoKeyVersionState
		doActivateKeys        bool
		doDeleteKey           bool
		expectNewKeyPairError bool
		expectSignError       bool
	}{
		{
			// Tests the basic case where the key is immediately ready to use.
			desc:            "key enabled",
			initialKeyState: kmspb.CryptoKeyVersion_ENABLED,
		},
		{
			// Tests the case where the key is temporarily pending and the
			// keystore needs to wait for it to be ready.
			desc:            "key pending temporarily",
			initialKeyState: kmspb.CryptoKeyVersion_PENDING_GENERATION,
			doActivateKeys:  true,
		},
		{
			// Tests the case where the key never becomes ready.
			desc:                  "key pending forever",
			initialKeyState:       kmspb.CryptoKeyVersion_PENDING_GENERATION,
			doActivateKeys:        false,
			expectNewKeyPairError: true,
		},
		{
			// Tests what happens when the key is externally deleted from the
			// KMS while a signer is held.
			desc:            "deleted externally",
			initialKeyState: kmspb.CryptoKeyVersion_ENABLED,
			doDeleteKey:     true,
			expectSignError: true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			testCtx, cancelTestCtx := context.WithCancel(context.Background())
			clientContext, cancelClientContext := context.WithCancel(testCtx)
			var wg sync.WaitGroup
			tickerErr := make(chan error)

			defer func() {
				cancelTestCtx()
				err := <-tickerErr
				require.ErrorIs(t, err, context.Canceled, "expected context.Canceled err from ticker goroutine")
				wg.Wait()
			}()

			// Create a fake clock and a channel which will receive pointers to
			// all newly created tickers.
			tickerCh := make(chan *faketime.FakeTicker)
			clock := faketime.NewFakeClock(tickerCh)

			fakeKMSServer, dialer := newTestGCPKMSService(t,
				withInitialKeyState(tc.initialKeyState),
			)
			kmsClient := newTestGCPKMSClient(t, dialer)
			manager, err := NewManager(testCtx, &servicecfg.KeystoreConfig{
				GCPKMS: servicecfg.GCPKMSConfig{
					ProtectionLevel: "HSM",
					KeyRing:         "test-keyring",
				},
			}, &Options{
				ClusterName:          clusterName,
				HostUUID:             "test-host-id",
				AuthPreferenceGetter: &fakeAuthPreferenceGetter{types.SignatureAlgorithmSuite_SIGNATURE_ALGORITHM_SUITE_HSM_V1},
				kmsClient:            kmsClient,
				faketimeOverride:     clock,
			})
			require.NoError(t, err, "error while creating test keystore manager")

			// Spin up a goroutine to fake all tickers created during execution
			// of the test. The keystore creates a ticker when waiting for a key
			// that is currently pending.
			handleTicker := func(ticker *faketime.FakeTicker) error {
				if ticker.Tag != (pendingRetryTag{}) {
					return trace.BadParameter("unknown ticker tag %v, test currently only handles tickers tagged with pendingRetryTag", ticker.Tag)
				}
				if tc.initialKeyState == kmspb.CryptoKeyVersion_ENABLED {
					// No need to tick if key is initially enabled.  Never
					// ticking is a form of assertion that the test doesn't halt
					// when the keystore reads the active key on the first try.
					return nil
				}
				if tc.initialKeyState != kmspb.CryptoKeyVersion_PENDING_GENERATION {
					return trace.BadParameter("unknown initial key state %v, test currently only handles ENABLED or PENDING_GENERATION", tc.initialKeyState)
				}
				// At this point we know the key is pending. Tick once to
				// guarantee the pending key is seen by the keystore.
				ticker.Tick(testCtx)
				if tc.doActivateKeys {
					fakeKMSServer.activateAllKeys()
					// Tick again so the keystore can see the active key. If the
					// active key was seen immediately after the previous tick
					// then the tick will never be received and this will hang
					// until testCtx is canceled, so run it in a background
					// goroutine.
					wg.Add(1)
					go func() {
						defer wg.Done()
						ticker.Tick(testCtx)
					}()
					return nil
				}
				// If we make it here the key will never be activated. Allow a
				// couple more ticks so that the key will be seen in a pending
				// state multiple times, then cancel the client context rather
				// than waiting for it to expire (injecting fake times and fake
				// contexts would be more complicated than it's worth).
				ticker.Tick(testCtx)
				ticker.Tick(testCtx)
				cancelClientContext()
				return nil
			}
			handleTickers := func() error {
				for {
					select {
					case <-testCtx.Done():
						return testCtx.Err()
					case ticker := <-tickerCh:
						if err := handleTicker(ticker); err != nil {
							return trace.Wrap(err)
						}
					}
				}
			}
			go func() {
				tickerErr <- handleTickers()
			}()

			// Test key creation.
			sshKeyPair, err := manager.NewSSHKeyPair(clientContext, cryptosuites.UserCASSH)
			if tc.expectNewKeyPairError {
				require.Error(t, err, "expected to get error generating SSH keypair, got nil")
				return
			}
			require.NoError(t, err, "unexpected error while generating SSH keypair")

			tlsKeyPair, err := manager.NewTLSKeyPair(clientContext, "test-cluster", cryptosuites.UserCATLS)
			require.NoError(t, err, "unexpected error creating TLS keypair")

			jwtKeyPair, err := manager.NewJWTKeyPair(clientContext, cryptosuites.JWTCAJWT)
			require.NoError(t, err, "unexpected error creating JWT keypair")

			// Put all the keys into a "CA" so that the keystore manager can
			// select them and we can test the public API.
			ca, err := types.NewCertAuthority(types.CertAuthoritySpecV2{
				Type:        types.HostCA,
				ClusterName: clusterName.GetClusterName(),
				ActiveKeys: types.CAKeySet{
					SSH: []*types.SSHKeyPair{sshKeyPair},
					TLS: []*types.TLSKeyPair{tlsKeyPair},
					JWT: []*types.JWTKeyPair{jwtKeyPair},
				},
			})
			require.NoError(t, err, "unexpected error creating CA")

			// Client private key that will be the basis of test certs to be signed.
			clientPrivKey, err := keys.ParsePrivateKey(testRSA2048PrivateKeyPEM)
			require.NoError(t, err)

			// Test signing an SSH certificate.
			t.Run("ssh", func(t *testing.T) {
				sshSigner, err := manager.GetSSHSigner(clientContext, ca)
				require.NoError(t, err, "unexpected error getting SSH signer")

				if tc.doDeleteKey {
					err = fakeKMSServer.deleteKey(sshKeyPair.PrivateKey)
					require.NoError(t, err, "unexpected error deleting SSH key")
				}

				// Make sure we can sign an SSH certificate, meaning we created the
				// correct type of key which can support the default 512 byte digest.
				cert := &ssh.Certificate{
					ValidPrincipals: []string{"root"},
					Key:             clientPrivKey.SSHPublicKey(),
					CertType:        ssh.HostCert,
				}
				err = cert.SignCert(mathrandv1.New(mathrandv1.NewSource(0)), sshSigner)
				if tc.expectSignError {
					require.Error(t, err, "expected to get error signing SSH cert")
					return
				}
				require.NoError(t, err, trace.DebugReport(err))
			})

			// Test signing a TLS certificate.
			t.Run("tls", func(t *testing.T) {
				tlsCert, tlsSigner, err := manager.GetTLSCertAndSigner(clientContext, ca)
				require.NoError(t, err, "unexpected error getting TLS cert")

				tlsCA, err := tlsca.FromCertAndSigner(tlsCert, tlsSigner)
				require.NoError(t, err, "unexpected error creating TLS CA")

				if tc.doDeleteKey {
					err = fakeKMSServer.deleteKey(tlsKeyPair.Key)
					require.NoError(t, err, "unexpected error deleting TLS key")
				}

				template := &x509.Certificate{
					SerialNumber: big.NewInt(1),
					Subject: pkix.Name{
						CommonName: "example.com",
					},
				}
				_, err = x509.CreateCertificate(
					mathrandv1.New(mathrandv1.NewSource(0)),
					template,
					tlsCA.Cert,
					clientPrivKey.Public(),
					tlsCA.Signer,
				)
				if tc.expectSignError {
					require.Error(t, err, "expected to get error signing TLS cert")
					return
				}
				require.NoError(t, err, "unexpected error signing TLS cert")
			})

			// Test signing a JWT.
			t.Run("jwt", func(t *testing.T) {
				jwtSigner, err := manager.GetJWTSigner(clientContext, ca)
				require.NoError(t, err, "unexpected error getting JWT signer")

				servicesJWTSigner, err := services.GetJWTSigner(jwtSigner, "test-cluster", nil)
				require.NoError(t, err, "unexpected error from services.GetJWTSigner")

				if tc.doDeleteKey {
					err = fakeKMSServer.deleteKey(jwtKeyPair.PrivateKey)
					require.NoError(t, err, "unexpected error deleting JWT key")
				}

				_, err = servicesJWTSigner.Sign(jwt.SignParams{
					Username: "root",
					Roles:    []string{"access"},
					URI:      "example.com",
					Expires:  time.Now().Add(time.Hour),
				})
				if tc.expectSignError {
					require.Error(t, err, "expected to get error signing JWT")
					return
				}
				require.NoError(t, err, "unexpected error signing JWT: %s", trace.DebugReport(err))
			})
		})
	}
}

func TestGCPKMSDeleteUnusedKeys(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	clusterName, err := services.NewClusterNameWithRandomID(types.ClusterNameSpecV2{
		ClusterName: "test-cluster",
	})
	require.NoError(t, err)

	const (
		localHostID  = "local-host-id"
		otherHostID  = "other-host-id"
		localKeyring = "local-keyring"
		otherKeyring = "other-keyring"
	)

	for _, tc := range []struct {
		desc            string
		existingKeys    []keySpec
		activeKeys      []keySpec
		expectDestroyed []keySpec
	}{
		{
			// Only inactive keys should be destroyed.
			desc: "active and inactive",
			existingKeys: []keySpec{
				{keyring: localKeyring, id: "id_active", host: localHostID},
				{keyring: localKeyring, id: "id_inactive", host: localHostID},
			},
			activeKeys: []keySpec{
				{keyring: localKeyring, id: "id_active", host: localHostID},
			},
			expectDestroyed: []keySpec{
				{keyring: localKeyring, id: "id_inactive", host: localHostID},
			},
		},
		{
			// Inactive key from other host should not be destroyed, it may
			// be recently created and just not added to Teleport CA yet, or the
			// other Auth might be in a completely different Teleport cluster
			// using the same keyring (I wouldn't advise this but someone might
			// do it).
			desc: "inactive key from other host",
			existingKeys: []keySpec{
				{keyring: localKeyring, id: "id_inactive_local", host: localHostID},
				{keyring: localKeyring, id: "id_inactive_remote", host: otherHostID},
			},
			expectDestroyed: []keySpec{
				{keyring: localKeyring, id: "id_inactive_local", host: localHostID},
			},
		},
		{
			// The presence of active keys created by a remote host in the local
			// keyring should not break the DeleteUnusedKeys operation.
			desc: "active key from other host",
			existingKeys: []keySpec{
				{keyring: localKeyring, id: "id_active_local", host: localHostID},
				{keyring: localKeyring, id: "id_inactive_local", host: localHostID},
				{keyring: localKeyring, id: "id_active_remote", host: otherHostID},
				{keyring: localKeyring, id: "id_inactive_remote", host: otherHostID},
			},
			activeKeys: []keySpec{
				{keyring: localKeyring, id: "id_active_local", host: localHostID},
				{keyring: localKeyring, id: "id_active_remote", host: otherHostID},
			},
			expectDestroyed: []keySpec{
				{keyring: localKeyring, id: "id_inactive_local", host: localHostID},
			},
		},
		{
			// Keys in other keyring should never be destroyed.
			desc: "keys in other keyring",
			existingKeys: []keySpec{
				{keyring: localKeyring, id: "id_inactive_local", host: localHostID},
				{keyring: otherKeyring, id: "id_inactive_other1", host: localHostID},
				{keyring: otherKeyring, id: "id_inactive_other2", host: otherHostID},
			},
			expectDestroyed: []keySpec{
				{keyring: localKeyring, id: "id_inactive_local", host: localHostID},
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			fakeKMSServer, dialer := newTestGCPKMSService(t)
			kmsClient := newTestGCPKMSClient(t, dialer)
			manager, err := NewManager(ctx, &servicecfg.KeystoreConfig{
				GCPKMS: servicecfg.GCPKMSConfig{
					ProtectionLevel: "HSM",
					KeyRing:         localKeyring,
				},
			}, &Options{
				ClusterName:          clusterName,
				HostUUID:             localHostID,
				AuthPreferenceGetter: &fakeAuthPreferenceGetter{types.SignatureAlgorithmSuite_SIGNATURE_ALGORITHM_SUITE_HSM_V1},
				kmsClient:            kmsClient,
			})
			require.NoError(t, err, "error while creating test keystore manager")

			// Pre-req: create existing keys in fake KMS backend
			for _, ks := range tc.existingKeys {
				_, err := fakeKMSServer.CreateCryptoKey(ctx, createKeyRequest(ks))
				require.NoError(t, err)
			}

			// Test: DeleteUnusedKeys with activeKeys from the testcase
			activeKeyIDs := make([][]byte, len(tc.activeKeys))
			for i, ks := range tc.activeKeys {
				activeKeyIDs[i] = ks.keyID()
			}
			err = manager.DeleteUnusedKeys(ctx, activeKeyIDs)
			assert.NoError(t, err)

			expectDestroyedSet := make(map[string]bool, len(tc.expectDestroyed))
			for _, ks := range tc.expectDestroyed {
				expectDestroyedSet[ks.keyVersionName()] = true
			}
			require.Len(t, fakeKMSServer.keyVersions, len(tc.existingKeys))
			for keyVersionName, state := range fakeKMSServer.keyVersions {
				if expectDestroyedSet[keyVersionName] {
					// Fake KMS server only sets state to DESTROY_SCHEDULED,
					// that's good enough for the test.
					require.Equal(t, kmspb.CryptoKeyVersion_DESTROY_SCHEDULED.String(), state.cryptoKeyVersion.State.String())
				} else {
					require.Equal(t, kmspb.CryptoKeyVersion_ENABLED.String(), state.cryptoKeyVersion.State.String())
				}
			}
		})
	}
}

type keySpec struct {
	keyring, id, host string
}

func (ks *keySpec) keyVersionName() string {
	return ks.keyring + "/cryptoKeys/" + ks.id + keyVersionSuffix
}

func (ks *keySpec) keyID() []byte {
	return gcpKMSKeyID{
		keyVersionName: ks.keyVersionName(),
	}.marshal()
}

func createKeyRequest(ks keySpec) *kmspb.CreateCryptoKeyRequest {
	return &kmspb.CreateCryptoKeyRequest{
		Parent:      ks.keyring,
		CryptoKeyId: ks.id,
		CryptoKey: &kmspb.CryptoKey{
			Purpose: kmspb.CryptoKey_ASYMMETRIC_SIGN,
			Labels: map[string]string{
				hostLabel: ks.host,
			},
			VersionTemplate: &kmspb.CryptoKeyVersionTemplate{
				ProtectionLevel: kmspb.ProtectionLevel_SOFTWARE,
				Algorithm:       kmspb.CryptoKeyVersion_RSA_SIGN_PKCS1_2048_SHA256,
			},
		},
	}
}
