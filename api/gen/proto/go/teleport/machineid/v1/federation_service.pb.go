// Copyright 2024 Gravitational, Inc
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: teleport/machineid/v1/federation_service.proto

package machineidv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetSPIFFEFederationRequest is the request message for GetSPIFFEFederation.
type GetSPIFFEFederationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the SPIFFEFederation resource to fetch.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSPIFFEFederationRequest) Reset() {
	*x = GetSPIFFEFederationRequest{}
	mi := &file_teleport_machineid_v1_federation_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSPIFFEFederationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSPIFFEFederationRequest) ProtoMessage() {}

func (x *GetSPIFFEFederationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_federation_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSPIFFEFederationRequest.ProtoReflect.Descriptor instead.
func (*GetSPIFFEFederationRequest) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_federation_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetSPIFFEFederationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request for ListSPIFFEFederations.
//
// Follows the pagination semantics of
// https://cloud.google.com/apis/design/standard_methods#list
type ListSPIFFEFederationsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The maximum number of items to return.
	// The server may impose a different page size at its discretion.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The page_token value returned from a previous ListSPIFFEFederations
	// request, if any.
	PageToken     string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSPIFFEFederationsRequest) Reset() {
	*x = ListSPIFFEFederationsRequest{}
	mi := &file_teleport_machineid_v1_federation_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSPIFFEFederationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSPIFFEFederationsRequest) ProtoMessage() {}

func (x *ListSPIFFEFederationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_federation_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSPIFFEFederationsRequest.ProtoReflect.Descriptor instead.
func (*ListSPIFFEFederationsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_federation_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListSPIFFEFederationsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListSPIFFEFederationsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// ListSPIFFEFederationsResponse is the response message for ListSPIFFEFederations.
type ListSPIFFEFederationsResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	SpiffeFederations []*SPIFFEFederation    `protobuf:"bytes,1,rep,name=spiffe_federations,json=spiffeFederations,proto3" json:"spiffe_federations,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results exist.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSPIFFEFederationsResponse) Reset() {
	*x = ListSPIFFEFederationsResponse{}
	mi := &file_teleport_machineid_v1_federation_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSPIFFEFederationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSPIFFEFederationsResponse) ProtoMessage() {}

func (x *ListSPIFFEFederationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_federation_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSPIFFEFederationsResponse.ProtoReflect.Descriptor instead.
func (*ListSPIFFEFederationsResponse) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_federation_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListSPIFFEFederationsResponse) GetSpiffeFederations() []*SPIFFEFederation {
	if x != nil {
		return x.SpiffeFederations
	}
	return nil
}

func (x *ListSPIFFEFederationsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// DeleteSPIFFEFederationRequest is the request message for DeleteSPIFFEFederation.
type DeleteSPIFFEFederationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the SPIFFEFederation resource to delete.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteSPIFFEFederationRequest) Reset() {
	*x = DeleteSPIFFEFederationRequest{}
	mi := &file_teleport_machineid_v1_federation_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSPIFFEFederationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSPIFFEFederationRequest) ProtoMessage() {}

func (x *DeleteSPIFFEFederationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_federation_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSPIFFEFederationRequest.ProtoReflect.Descriptor instead.
func (*DeleteSPIFFEFederationRequest) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_federation_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteSPIFFEFederationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// CreateSPIFFEFederationRequest is the request message for CreateSPIFFEFederation.
type CreateSPIFFEFederationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The SPIFFEFederation resource to create.
	SpiffeFederation *SPIFFEFederation `protobuf:"bytes,1,opt,name=spiffe_federation,json=spiffeFederation,proto3" json:"spiffe_federation,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CreateSPIFFEFederationRequest) Reset() {
	*x = CreateSPIFFEFederationRequest{}
	mi := &file_teleport_machineid_v1_federation_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSPIFFEFederationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSPIFFEFederationRequest) ProtoMessage() {}

func (x *CreateSPIFFEFederationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_federation_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSPIFFEFederationRequest.ProtoReflect.Descriptor instead.
func (*CreateSPIFFEFederationRequest) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_federation_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateSPIFFEFederationRequest) GetSpiffeFederation() *SPIFFEFederation {
	if x != nil {
		return x.SpiffeFederation
	}
	return nil
}

var File_teleport_machineid_v1_federation_service_proto protoreflect.FileDescriptor

var file_teleport_machineid_v1_federation_service_proto_rawDesc = string([]byte{
	0x0a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5a,
	0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9f, 0x01, 0x0a, 0x1d, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x12,
	0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x5f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x11, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x33, 0x0a, 0x1d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x75, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x50, 0x49, 0x46, 0x46,
	0x45, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x54, 0x0a, 0x11, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x5f, 0x66, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xf2, 0x03, 0x0a, 0x17, 0x53, 0x50, 0x49,
	0x46, 0x46, 0x45, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x50, 0x49, 0x46, 0x46,
	0x45, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x46, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x46, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x82, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x33, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x50,
	0x49, 0x46, 0x46, 0x45, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x16,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x46, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x77, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x50,
	0x49, 0x46, 0x46, 0x45, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x50, 0x49,
	0x46, 0x46, 0x45, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x50, 0x49,
	0x46, 0x46, 0x45, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x56, 0x5a,
	0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x69, 0x64, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_teleport_machineid_v1_federation_service_proto_rawDescOnce sync.Once
	file_teleport_machineid_v1_federation_service_proto_rawDescData []byte
)

func file_teleport_machineid_v1_federation_service_proto_rawDescGZIP() []byte {
	file_teleport_machineid_v1_federation_service_proto_rawDescOnce.Do(func() {
		file_teleport_machineid_v1_federation_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_teleport_machineid_v1_federation_service_proto_rawDesc), len(file_teleport_machineid_v1_federation_service_proto_rawDesc)))
	})
	return file_teleport_machineid_v1_federation_service_proto_rawDescData
}

var file_teleport_machineid_v1_federation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_teleport_machineid_v1_federation_service_proto_goTypes = []any{
	(*GetSPIFFEFederationRequest)(nil),    // 0: teleport.machineid.v1.GetSPIFFEFederationRequest
	(*ListSPIFFEFederationsRequest)(nil),  // 1: teleport.machineid.v1.ListSPIFFEFederationsRequest
	(*ListSPIFFEFederationsResponse)(nil), // 2: teleport.machineid.v1.ListSPIFFEFederationsResponse
	(*DeleteSPIFFEFederationRequest)(nil), // 3: teleport.machineid.v1.DeleteSPIFFEFederationRequest
	(*CreateSPIFFEFederationRequest)(nil), // 4: teleport.machineid.v1.CreateSPIFFEFederationRequest
	(*SPIFFEFederation)(nil),              // 5: teleport.machineid.v1.SPIFFEFederation
	(*emptypb.Empty)(nil),                 // 6: google.protobuf.Empty
}
var file_teleport_machineid_v1_federation_service_proto_depIdxs = []int32{
	5, // 0: teleport.machineid.v1.ListSPIFFEFederationsResponse.spiffe_federations:type_name -> teleport.machineid.v1.SPIFFEFederation
	5, // 1: teleport.machineid.v1.CreateSPIFFEFederationRequest.spiffe_federation:type_name -> teleport.machineid.v1.SPIFFEFederation
	0, // 2: teleport.machineid.v1.SPIFFEFederationService.GetSPIFFEFederation:input_type -> teleport.machineid.v1.GetSPIFFEFederationRequest
	1, // 3: teleport.machineid.v1.SPIFFEFederationService.ListSPIFFEFederations:input_type -> teleport.machineid.v1.ListSPIFFEFederationsRequest
	3, // 4: teleport.machineid.v1.SPIFFEFederationService.DeleteSPIFFEFederation:input_type -> teleport.machineid.v1.DeleteSPIFFEFederationRequest
	4, // 5: teleport.machineid.v1.SPIFFEFederationService.CreateSPIFFEFederation:input_type -> teleport.machineid.v1.CreateSPIFFEFederationRequest
	5, // 6: teleport.machineid.v1.SPIFFEFederationService.GetSPIFFEFederation:output_type -> teleport.machineid.v1.SPIFFEFederation
	2, // 7: teleport.machineid.v1.SPIFFEFederationService.ListSPIFFEFederations:output_type -> teleport.machineid.v1.ListSPIFFEFederationsResponse
	6, // 8: teleport.machineid.v1.SPIFFEFederationService.DeleteSPIFFEFederation:output_type -> google.protobuf.Empty
	5, // 9: teleport.machineid.v1.SPIFFEFederationService.CreateSPIFFEFederation:output_type -> teleport.machineid.v1.SPIFFEFederation
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_teleport_machineid_v1_federation_service_proto_init() }
func file_teleport_machineid_v1_federation_service_proto_init() {
	if File_teleport_machineid_v1_federation_service_proto != nil {
		return
	}
	file_teleport_machineid_v1_federation_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_teleport_machineid_v1_federation_service_proto_rawDesc), len(file_teleport_machineid_v1_federation_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_machineid_v1_federation_service_proto_goTypes,
		DependencyIndexes: file_teleport_machineid_v1_federation_service_proto_depIdxs,
		MessageInfos:      file_teleport_machineid_v1_federation_service_proto_msgTypes,
	}.Build()
	File_teleport_machineid_v1_federation_service_proto = out.File
	file_teleport_machineid_v1_federation_service_proto_goTypes = nil
	file_teleport_machineid_v1_federation_service_proto_depIdxs = nil
}
