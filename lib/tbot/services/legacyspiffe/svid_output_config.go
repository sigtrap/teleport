/*
 * Teleport
 * Copyright (C) 2025  Gravitational, Inc.
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

package legacyspiffe

import (
	"context"
	"log/slog"
	"net"
	"strings"

	"github.com/gravitational/trace"
	"gopkg.in/yaml.v3"

	"github.com/gravitational/teleport/lib/tbot/bot"
	"github.com/gravitational/teleport/lib/tbot/bot/destination"
	"github.com/gravitational/teleport/lib/tbot/internal"
	"github.com/gravitational/teleport/lib/tbot/internal/encoding"
)

const SVIDOutputServiceType = "spiffe-svid"

// SVIDRequestSANs is the configuration for the SANs of a single SVID request.
type SVIDRequestSANs struct {
	// DNS is the list of DNS names that are requested to be included in the SVID.
	DNS []string `yaml:"dns,omitempty"`
	// IP is the list of IP addresses that are requested to be included in the SVID.
	// These can be IPv4 or IPv6 addresses.
	IP []string `yaml:"ip,omitempty"`
}

// SVIDRequest is the configuration for a single SVID request.
type SVIDRequest struct {
	// Path is the SPIFFE ID path of the SVID. It should be prefixed with "/".
	Path string `yaml:"path,omitempty"`
	// Hint is the hint for the SVID that will be provided to consumers of the
	// SVID to help them identify it.
	Hint string `yaml:"hint,omitempty"`
	// SANS is the Subject Alternative Names that are requested to be included
	// in the SVID.
	SANS SVIDRequestSANs `yaml:"sans,omitempty"`
}

func (o SVIDRequest) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("path", o.Path),
		slog.String("hint", o.Hint),
		slog.Any("dns_sans", o.SANS.DNS),
		slog.Any("ip_sans", o.SANS.IP),
	)
}

// CheckAndSetDefaults checks the SVIDRequest values and sets any defaults.
func (o *SVIDRequest) CheckAndSetDefaults() error {
	switch {
	case o.Path == "":
		return trace.BadParameter("svid.path: should not be empty")
	case !strings.HasPrefix(o.Path, "/"):
		return trace.BadParameter("svid.path: should be prefixed with /")
	}
	for i, stringIP := range o.SANS.IP {
		ip := net.ParseIP(stringIP)
		if ip == nil {
			return trace.BadParameter(
				"ip_sans[%d]: invalid IP address %q", i, stringIP,
			)
		}
	}
	return nil
}

// JWTSVID the configuration for a single JWT SVID request as part of the SPIFFE
// SVID output.
type JWTSVID struct {
	// FileName is the name of the artifact/file the JWT should be written to.
	FileName string `yaml:"file_name"`
	// Audience is the audience of the JWT.
	Audience string `yaml:"audience"`
}

func (o JWTSVID) CheckAndSetDefaults() error {
	switch {
	case o.Audience == "":
		return trace.BadParameter("audience: should not be empty")
	case o.FileName == "":
		return trace.BadParameter("name: should not be empty")
	}
	return nil
}

// SVIDOutputConfig is the configuration for the SPIFFE SVID output.
// Emulates the output of https://github.com/spiffe/spiffe-helper
type SVIDOutputConfig struct {
	// Name of the service for logs and the /readyz endpoint.
	Name string `yaml:"name,omitempty"`
	// Destination is where the credentials should be written to.
	Destination                  destination.Destination `yaml:"destination"`
	SVID                         SVIDRequest             `yaml:"svid"`
	IncludeFederatedTrustBundles bool                    `yaml:"include_federated_trust_bundles,omitempty"`
	// JWTs is an optional list of audiences and file names to write JWT SVIDs
	// to.
	JWTs []JWTSVID `yaml:"jwts,omitempty"`

	// CredentialLifetime contains configuration for how long credentials will
	// last and the frequency at which they'll be renewed.
	CredentialLifetime bot.CredentialLifetime `yaml:",inline"`
}

// GetName returns the user-given name of the service, used for validation purposes.
func (o *SVIDOutputConfig) GetName() string {
	return o.Name
}

// Init initializes the destination.
func (o *SVIDOutputConfig) Init(ctx context.Context) error {
	return trace.Wrap(o.Destination.Init(ctx, []string{}))
}

// GetDestination returns the destination.
func (o *SVIDOutputConfig) GetDestination() destination.Destination {
	return o.Destination
}

// CheckAndSetDefaults checks the SPIFFESVIDOutput values and sets any defaults.
func (o *SVIDOutputConfig) CheckAndSetDefaults() error {
	if err := o.SVID.CheckAndSetDefaults(); err != nil {
		return trace.Wrap(err, "validating svid")
	}
	if o.Destination == nil {
		return trace.BadParameter("no destination configured for output")
	}
	if err := o.Destination.CheckAndSetDefaults(); err != nil {
		return trace.Wrap(err, "validating destination")
	}
	for i, jwt := range o.JWTs {
		if err := jwt.CheckAndSetDefaults(); err != nil {
			return trace.Wrap(err, "validating jwts[%d]", i)
		}
	}
	return nil
}

// Describe returns the file descriptions for the SPIFFE SVID output.
func (o *SVIDOutputConfig) Describe() []bot.FileDescription {
	fds := []bot.FileDescription{
		{
			Name: internal.SVIDPEMPath,
		},
		{
			Name: internal.SVIDKeyPEMPath,
		},
		{
			Name: internal.SVIDTrustBundlePEMPath,
		},
	}
	for _, jwt := range o.JWTs {
		fds = append(fds, bot.FileDescription{Name: jwt.FileName})
	}
	return nil
}

func (o *SVIDOutputConfig) Type() string {
	return SVIDOutputServiceType
}

// MarshalYAML marshals the SPIFFESVIDOutput into YAML.
func (o *SVIDOutputConfig) MarshalYAML() (any, error) {
	type raw SVIDOutputConfig
	return encoding.WithTypeHeader((*raw)(o), SVIDOutputServiceType)
}

func (o *SVIDOutputConfig) UnmarshalYAML(*yaml.Node) error {
	return trace.NotImplemented("unmarshaling %T with UnmarshalYAML is not supported, use UnmarshalConfig instead", o)
}

func (o *SVIDOutputConfig) UnmarshalConfig(ctx bot.UnmarshalConfigContext, node *yaml.Node) error {
	dest, err := internal.ExtractOutputDestination(ctx, node)
	if err != nil {
		return trace.Wrap(err)
	}
	// Alias type to remove UnmarshalYAML to avoid getting our "not implemented" error
	type raw SVIDOutputConfig
	if err := node.Decode((*raw)(o)); err != nil {
		return trace.Wrap(err)
	}
	o.Destination = dest
	return nil
}

func (o *SVIDOutputConfig) GetCredentialLifetime() bot.CredentialLifetime {
	return o.CredentialLifetime
}
