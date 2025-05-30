// Copyright 2023 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: teleport/integration/v1/integration_service.proto

package integrationv1

import (
	context "context"
	types "github.com/gravitational/teleport/api/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	IntegrationService_ListIntegrations_FullMethodName                 = "/teleport.integration.v1.IntegrationService/ListIntegrations"
	IntegrationService_GetIntegration_FullMethodName                   = "/teleport.integration.v1.IntegrationService/GetIntegration"
	IntegrationService_CreateIntegration_FullMethodName                = "/teleport.integration.v1.IntegrationService/CreateIntegration"
	IntegrationService_UpdateIntegration_FullMethodName                = "/teleport.integration.v1.IntegrationService/UpdateIntegration"
	IntegrationService_DeleteIntegration_FullMethodName                = "/teleport.integration.v1.IntegrationService/DeleteIntegration"
	IntegrationService_DeleteAllIntegrations_FullMethodName            = "/teleport.integration.v1.IntegrationService/DeleteAllIntegrations"
	IntegrationService_GenerateAWSOIDCToken_FullMethodName             = "/teleport.integration.v1.IntegrationService/GenerateAWSOIDCToken"
	IntegrationService_GenerateAzureOIDCToken_FullMethodName           = "/teleport.integration.v1.IntegrationService/GenerateAzureOIDCToken"
	IntegrationService_GenerateGitHubUserCert_FullMethodName           = "/teleport.integration.v1.IntegrationService/GenerateGitHubUserCert"
	IntegrationService_ExportIntegrationCertAuthorities_FullMethodName = "/teleport.integration.v1.IntegrationService/ExportIntegrationCertAuthorities"
	IntegrationService_GenerateAWSRACredentials_FullMethodName         = "/teleport.integration.v1.IntegrationService/GenerateAWSRACredentials"
)

// IntegrationServiceClient is the client API for IntegrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// IntegrationService provides methods to manage Integrations with 3rd party APIs.
type IntegrationServiceClient interface {
	// ListIntegrations returns a paginated list of Integration resources.
	ListIntegrations(ctx context.Context, in *ListIntegrationsRequest, opts ...grpc.CallOption) (*ListIntegrationsResponse, error)
	// GetIntegration returns the specified Integration resource.
	GetIntegration(ctx context.Context, in *GetIntegrationRequest, opts ...grpc.CallOption) (*types.IntegrationV1, error)
	// CreateIntegration creates a new Integration resource.
	CreateIntegration(ctx context.Context, in *CreateIntegrationRequest, opts ...grpc.CallOption) (*types.IntegrationV1, error)
	// UpdateIntegration updates an existing Integration resource.
	UpdateIntegration(ctx context.Context, in *UpdateIntegrationRequest, opts ...grpc.CallOption) (*types.IntegrationV1, error)
	// DeleteIntegration removes the specified Integration resource.
	DeleteIntegration(ctx context.Context, in *DeleteIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DeleteAllIntegrations removes all Integrations.
	// DEPRECATED: Can't delete all integrations over gRPC.
	DeleteAllIntegrations(ctx context.Context, in *DeleteAllIntegrationsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GenerateAWSOIDCToken generates a token to be used when executing an AWS OIDC Integration action.
	GenerateAWSOIDCToken(ctx context.Context, in *GenerateAWSOIDCTokenRequest, opts ...grpc.CallOption) (*GenerateAWSOIDCTokenResponse, error)
	// GenerateAzureOIDCToken generates a token to be used when executing an Azure OIDC Integration action.
	GenerateAzureOIDCToken(ctx context.Context, in *GenerateAzureOIDCTokenRequest, opts ...grpc.CallOption) (*GenerateAzureOIDCTokenResponse, error)
	// GenerateGitHubUserCert signs a SSH certificate for GitHub integration.
	GenerateGitHubUserCert(ctx context.Context, in *GenerateGitHubUserCertRequest, opts ...grpc.CallOption) (*GenerateGitHubUserCertResponse, error)
	// ExportIntegrationCertAuthorities exports cert authorities for an integration.
	ExportIntegrationCertAuthorities(ctx context.Context, in *ExportIntegrationCertAuthoritiesRequest, opts ...grpc.CallOption) (*ExportIntegrationCertAuthoritiesResponse, error)
	// GenerateAWSRACredentials generates a set of AWS Credentials using the AWS IAM Roles Anywhere integration.
	GenerateAWSRACredentials(ctx context.Context, in *GenerateAWSRACredentialsRequest, opts ...grpc.CallOption) (*GenerateAWSRACredentialsResponse, error)
}

type integrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIntegrationServiceClient(cc grpc.ClientConnInterface) IntegrationServiceClient {
	return &integrationServiceClient{cc}
}

func (c *integrationServiceClient) ListIntegrations(ctx context.Context, in *ListIntegrationsRequest, opts ...grpc.CallOption) (*ListIntegrationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListIntegrationsResponse)
	err := c.cc.Invoke(ctx, IntegrationService_ListIntegrations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationServiceClient) GetIntegration(ctx context.Context, in *GetIntegrationRequest, opts ...grpc.CallOption) (*types.IntegrationV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(types.IntegrationV1)
	err := c.cc.Invoke(ctx, IntegrationService_GetIntegration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationServiceClient) CreateIntegration(ctx context.Context, in *CreateIntegrationRequest, opts ...grpc.CallOption) (*types.IntegrationV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(types.IntegrationV1)
	err := c.cc.Invoke(ctx, IntegrationService_CreateIntegration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationServiceClient) UpdateIntegration(ctx context.Context, in *UpdateIntegrationRequest, opts ...grpc.CallOption) (*types.IntegrationV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(types.IntegrationV1)
	err := c.cc.Invoke(ctx, IntegrationService_UpdateIntegration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationServiceClient) DeleteIntegration(ctx context.Context, in *DeleteIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, IntegrationService_DeleteIntegration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationServiceClient) DeleteAllIntegrations(ctx context.Context, in *DeleteAllIntegrationsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, IntegrationService_DeleteAllIntegrations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationServiceClient) GenerateAWSOIDCToken(ctx context.Context, in *GenerateAWSOIDCTokenRequest, opts ...grpc.CallOption) (*GenerateAWSOIDCTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateAWSOIDCTokenResponse)
	err := c.cc.Invoke(ctx, IntegrationService_GenerateAWSOIDCToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationServiceClient) GenerateAzureOIDCToken(ctx context.Context, in *GenerateAzureOIDCTokenRequest, opts ...grpc.CallOption) (*GenerateAzureOIDCTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateAzureOIDCTokenResponse)
	err := c.cc.Invoke(ctx, IntegrationService_GenerateAzureOIDCToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationServiceClient) GenerateGitHubUserCert(ctx context.Context, in *GenerateGitHubUserCertRequest, opts ...grpc.CallOption) (*GenerateGitHubUserCertResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateGitHubUserCertResponse)
	err := c.cc.Invoke(ctx, IntegrationService_GenerateGitHubUserCert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationServiceClient) ExportIntegrationCertAuthorities(ctx context.Context, in *ExportIntegrationCertAuthoritiesRequest, opts ...grpc.CallOption) (*ExportIntegrationCertAuthoritiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExportIntegrationCertAuthoritiesResponse)
	err := c.cc.Invoke(ctx, IntegrationService_ExportIntegrationCertAuthorities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationServiceClient) GenerateAWSRACredentials(ctx context.Context, in *GenerateAWSRACredentialsRequest, opts ...grpc.CallOption) (*GenerateAWSRACredentialsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateAWSRACredentialsResponse)
	err := c.cc.Invoke(ctx, IntegrationService_GenerateAWSRACredentials_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IntegrationServiceServer is the server API for IntegrationService service.
// All implementations must embed UnimplementedIntegrationServiceServer
// for forward compatibility.
//
// IntegrationService provides methods to manage Integrations with 3rd party APIs.
type IntegrationServiceServer interface {
	// ListIntegrations returns a paginated list of Integration resources.
	ListIntegrations(context.Context, *ListIntegrationsRequest) (*ListIntegrationsResponse, error)
	// GetIntegration returns the specified Integration resource.
	GetIntegration(context.Context, *GetIntegrationRequest) (*types.IntegrationV1, error)
	// CreateIntegration creates a new Integration resource.
	CreateIntegration(context.Context, *CreateIntegrationRequest) (*types.IntegrationV1, error)
	// UpdateIntegration updates an existing Integration resource.
	UpdateIntegration(context.Context, *UpdateIntegrationRequest) (*types.IntegrationV1, error)
	// DeleteIntegration removes the specified Integration resource.
	DeleteIntegration(context.Context, *DeleteIntegrationRequest) (*emptypb.Empty, error)
	// DeleteAllIntegrations removes all Integrations.
	// DEPRECATED: Can't delete all integrations over gRPC.
	DeleteAllIntegrations(context.Context, *DeleteAllIntegrationsRequest) (*emptypb.Empty, error)
	// GenerateAWSOIDCToken generates a token to be used when executing an AWS OIDC Integration action.
	GenerateAWSOIDCToken(context.Context, *GenerateAWSOIDCTokenRequest) (*GenerateAWSOIDCTokenResponse, error)
	// GenerateAzureOIDCToken generates a token to be used when executing an Azure OIDC Integration action.
	GenerateAzureOIDCToken(context.Context, *GenerateAzureOIDCTokenRequest) (*GenerateAzureOIDCTokenResponse, error)
	// GenerateGitHubUserCert signs a SSH certificate for GitHub integration.
	GenerateGitHubUserCert(context.Context, *GenerateGitHubUserCertRequest) (*GenerateGitHubUserCertResponse, error)
	// ExportIntegrationCertAuthorities exports cert authorities for an integration.
	ExportIntegrationCertAuthorities(context.Context, *ExportIntegrationCertAuthoritiesRequest) (*ExportIntegrationCertAuthoritiesResponse, error)
	// GenerateAWSRACredentials generates a set of AWS Credentials using the AWS IAM Roles Anywhere integration.
	GenerateAWSRACredentials(context.Context, *GenerateAWSRACredentialsRequest) (*GenerateAWSRACredentialsResponse, error)
	mustEmbedUnimplementedIntegrationServiceServer()
}

// UnimplementedIntegrationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedIntegrationServiceServer struct{}

func (UnimplementedIntegrationServiceServer) ListIntegrations(context.Context, *ListIntegrationsRequest) (*ListIntegrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIntegrations not implemented")
}
func (UnimplementedIntegrationServiceServer) GetIntegration(context.Context, *GetIntegrationRequest) (*types.IntegrationV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIntegration not implemented")
}
func (UnimplementedIntegrationServiceServer) CreateIntegration(context.Context, *CreateIntegrationRequest) (*types.IntegrationV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIntegration not implemented")
}
func (UnimplementedIntegrationServiceServer) UpdateIntegration(context.Context, *UpdateIntegrationRequest) (*types.IntegrationV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIntegration not implemented")
}
func (UnimplementedIntegrationServiceServer) DeleteIntegration(context.Context, *DeleteIntegrationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIntegration not implemented")
}
func (UnimplementedIntegrationServiceServer) DeleteAllIntegrations(context.Context, *DeleteAllIntegrationsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllIntegrations not implemented")
}
func (UnimplementedIntegrationServiceServer) GenerateAWSOIDCToken(context.Context, *GenerateAWSOIDCTokenRequest) (*GenerateAWSOIDCTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAWSOIDCToken not implemented")
}
func (UnimplementedIntegrationServiceServer) GenerateAzureOIDCToken(context.Context, *GenerateAzureOIDCTokenRequest) (*GenerateAzureOIDCTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAzureOIDCToken not implemented")
}
func (UnimplementedIntegrationServiceServer) GenerateGitHubUserCert(context.Context, *GenerateGitHubUserCertRequest) (*GenerateGitHubUserCertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateGitHubUserCert not implemented")
}
func (UnimplementedIntegrationServiceServer) ExportIntegrationCertAuthorities(context.Context, *ExportIntegrationCertAuthoritiesRequest) (*ExportIntegrationCertAuthoritiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportIntegrationCertAuthorities not implemented")
}
func (UnimplementedIntegrationServiceServer) GenerateAWSRACredentials(context.Context, *GenerateAWSRACredentialsRequest) (*GenerateAWSRACredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAWSRACredentials not implemented")
}
func (UnimplementedIntegrationServiceServer) mustEmbedUnimplementedIntegrationServiceServer() {}
func (UnimplementedIntegrationServiceServer) testEmbeddedByValue()                            {}

// UnsafeIntegrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IntegrationServiceServer will
// result in compilation errors.
type UnsafeIntegrationServiceServer interface {
	mustEmbedUnimplementedIntegrationServiceServer()
}

func RegisterIntegrationServiceServer(s grpc.ServiceRegistrar, srv IntegrationServiceServer) {
	// If the following call pancis, it indicates UnimplementedIntegrationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&IntegrationService_ServiceDesc, srv)
}

func _IntegrationService_ListIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIntegrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).ListIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationService_ListIntegrations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).ListIntegrations(ctx, req.(*ListIntegrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationService_GetIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).GetIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationService_GetIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).GetIntegration(ctx, req.(*GetIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationService_CreateIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).CreateIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationService_CreateIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).CreateIntegration(ctx, req.(*CreateIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationService_UpdateIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).UpdateIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationService_UpdateIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).UpdateIntegration(ctx, req.(*UpdateIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationService_DeleteIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).DeleteIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationService_DeleteIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).DeleteIntegration(ctx, req.(*DeleteIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationService_DeleteAllIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllIntegrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).DeleteAllIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationService_DeleteAllIntegrations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).DeleteAllIntegrations(ctx, req.(*DeleteAllIntegrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationService_GenerateAWSOIDCToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAWSOIDCTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).GenerateAWSOIDCToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationService_GenerateAWSOIDCToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).GenerateAWSOIDCToken(ctx, req.(*GenerateAWSOIDCTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationService_GenerateAzureOIDCToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAzureOIDCTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).GenerateAzureOIDCToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationService_GenerateAzureOIDCToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).GenerateAzureOIDCToken(ctx, req.(*GenerateAzureOIDCTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationService_GenerateGitHubUserCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateGitHubUserCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).GenerateGitHubUserCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationService_GenerateGitHubUserCert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).GenerateGitHubUserCert(ctx, req.(*GenerateGitHubUserCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationService_ExportIntegrationCertAuthorities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportIntegrationCertAuthoritiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).ExportIntegrationCertAuthorities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationService_ExportIntegrationCertAuthorities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).ExportIntegrationCertAuthorities(ctx, req.(*ExportIntegrationCertAuthoritiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationService_GenerateAWSRACredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAWSRACredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).GenerateAWSRACredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationService_GenerateAWSRACredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).GenerateAWSRACredentials(ctx, req.(*GenerateAWSRACredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IntegrationService_ServiceDesc is the grpc.ServiceDesc for IntegrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IntegrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.integration.v1.IntegrationService",
	HandlerType: (*IntegrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListIntegrations",
			Handler:    _IntegrationService_ListIntegrations_Handler,
		},
		{
			MethodName: "GetIntegration",
			Handler:    _IntegrationService_GetIntegration_Handler,
		},
		{
			MethodName: "CreateIntegration",
			Handler:    _IntegrationService_CreateIntegration_Handler,
		},
		{
			MethodName: "UpdateIntegration",
			Handler:    _IntegrationService_UpdateIntegration_Handler,
		},
		{
			MethodName: "DeleteIntegration",
			Handler:    _IntegrationService_DeleteIntegration_Handler,
		},
		{
			MethodName: "DeleteAllIntegrations",
			Handler:    _IntegrationService_DeleteAllIntegrations_Handler,
		},
		{
			MethodName: "GenerateAWSOIDCToken",
			Handler:    _IntegrationService_GenerateAWSOIDCToken_Handler,
		},
		{
			MethodName: "GenerateAzureOIDCToken",
			Handler:    _IntegrationService_GenerateAzureOIDCToken_Handler,
		},
		{
			MethodName: "GenerateGitHubUserCert",
			Handler:    _IntegrationService_GenerateGitHubUserCert_Handler,
		},
		{
			MethodName: "ExportIntegrationCertAuthorities",
			Handler:    _IntegrationService_ExportIntegrationCertAuthorities_Handler,
		},
		{
			MethodName: "GenerateAWSRACredentials",
			Handler:    _IntegrationService_GenerateAWSRACredentials_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/integration/v1/integration_service.proto",
}
