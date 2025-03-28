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
// source: teleport/decision/v1alpha1/ssh_access.proto

package decisionpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

// SSHPortForwardMode describes the mode of port forwarding permitted.
type SSHPortForwardMode int32

const (
	SSHPortForwardMode_SSH_PORT_FORWARD_MODE_UNSPECIFIED SSHPortForwardMode = 0
	SSHPortForwardMode_SSH_PORT_FORWARD_MODE_OFF         SSHPortForwardMode = 1
	SSHPortForwardMode_SSH_PORT_FORWARD_MODE_ON          SSHPortForwardMode = 2
	SSHPortForwardMode_SSH_PORT_FORWARD_MODE_LOCAL       SSHPortForwardMode = 3
	SSHPortForwardMode_SSH_PORT_FORWARD_MODE_REMOTE      SSHPortForwardMode = 4
)

// Enum value maps for SSHPortForwardMode.
var (
	SSHPortForwardMode_name = map[int32]string{
		0: "SSH_PORT_FORWARD_MODE_UNSPECIFIED",
		1: "SSH_PORT_FORWARD_MODE_OFF",
		2: "SSH_PORT_FORWARD_MODE_ON",
		3: "SSH_PORT_FORWARD_MODE_LOCAL",
		4: "SSH_PORT_FORWARD_MODE_REMOTE",
	}
	SSHPortForwardMode_value = map[string]int32{
		"SSH_PORT_FORWARD_MODE_UNSPECIFIED": 0,
		"SSH_PORT_FORWARD_MODE_OFF":         1,
		"SSH_PORT_FORWARD_MODE_ON":          2,
		"SSH_PORT_FORWARD_MODE_LOCAL":       3,
		"SSH_PORT_FORWARD_MODE_REMOTE":      4,
	}
)

func (x SSHPortForwardMode) Enum() *SSHPortForwardMode {
	p := new(SSHPortForwardMode)
	*p = x
	return p
}

func (x SSHPortForwardMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SSHPortForwardMode) Descriptor() protoreflect.EnumDescriptor {
	return file_teleport_decision_v1alpha1_ssh_access_proto_enumTypes[0].Descriptor()
}

func (SSHPortForwardMode) Type() protoreflect.EnumType {
	return &file_teleport_decision_v1alpha1_ssh_access_proto_enumTypes[0]
}

func (x SSHPortForwardMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SSHPortForwardMode.Descriptor instead.
func (SSHPortForwardMode) EnumDescriptor() ([]byte, []int) {
	return file_teleport_decision_v1alpha1_ssh_access_proto_rawDescGZIP(), []int{0}
}

// EvaluateSSHAccessRequest describes a request to evaluate whether or not a
// given ssh access attempt should be permitted.
type EvaluateSSHAccessRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Metadata holds common authorization decision request fields.
	Metadata *RequestMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// SshAuthority identifies the authority that issued the below identity.
	SshAuthority *SSHAuthority `protobuf:"bytes,2,opt,name=ssh_authority,json=sshAuthority,proto3" json:"ssh_authority,omitempty"`
	// SshIdentity describes the teleport user requesting access.
	SshIdentity *SSHIdentity `protobuf:"bytes,3,opt,name=ssh_identity,json=sshIdentity,proto3" json:"ssh_identity,omitempty"`
	// Node references the target node the user is attempting to access.
	Node *Resource `protobuf:"bytes,4,opt,name=node,proto3" json:"node,omitempty"`
	// OSUser is the user on the target node the user is attempting to access.
	OsUser        string `protobuf:"bytes,5,opt,name=os_user,json=osUser,proto3" json:"os_user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EvaluateSSHAccessRequest) Reset() {
	*x = EvaluateSSHAccessRequest{}
	mi := &file_teleport_decision_v1alpha1_ssh_access_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EvaluateSSHAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluateSSHAccessRequest) ProtoMessage() {}

func (x *EvaluateSSHAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_decision_v1alpha1_ssh_access_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluateSSHAccessRequest.ProtoReflect.Descriptor instead.
func (*EvaluateSSHAccessRequest) Descriptor() ([]byte, []int) {
	return file_teleport_decision_v1alpha1_ssh_access_proto_rawDescGZIP(), []int{0}
}

func (x *EvaluateSSHAccessRequest) GetMetadata() *RequestMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *EvaluateSSHAccessRequest) GetSshAuthority() *SSHAuthority {
	if x != nil {
		return x.SshAuthority
	}
	return nil
}

func (x *EvaluateSSHAccessRequest) GetSshIdentity() *SSHIdentity {
	if x != nil {
		return x.SshIdentity
	}
	return nil
}

func (x *EvaluateSSHAccessRequest) GetNode() *Resource {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *EvaluateSSHAccessRequest) GetOsUser() string {
	if x != nil {
		return x.OsUser
	}
	return ""
}

// EvaluateSSHAccessResponse describes the result of an SSH access evaluation.
type EvaluateSSHAccessResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Decision:
	//
	//	*EvaluateSSHAccessResponse_Permit
	//	*EvaluateSSHAccessResponse_Denial
	Decision      isEvaluateSSHAccessResponse_Decision `protobuf_oneof:"decision"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EvaluateSSHAccessResponse) Reset() {
	*x = EvaluateSSHAccessResponse{}
	mi := &file_teleport_decision_v1alpha1_ssh_access_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EvaluateSSHAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluateSSHAccessResponse) ProtoMessage() {}

func (x *EvaluateSSHAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_decision_v1alpha1_ssh_access_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluateSSHAccessResponse.ProtoReflect.Descriptor instead.
func (*EvaluateSSHAccessResponse) Descriptor() ([]byte, []int) {
	return file_teleport_decision_v1alpha1_ssh_access_proto_rawDescGZIP(), []int{1}
}

func (x *EvaluateSSHAccessResponse) GetDecision() isEvaluateSSHAccessResponse_Decision {
	if x != nil {
		return x.Decision
	}
	return nil
}

func (x *EvaluateSSHAccessResponse) GetPermit() *SSHAccessPermit {
	if x != nil {
		if x, ok := x.Decision.(*EvaluateSSHAccessResponse_Permit); ok {
			return x.Permit
		}
	}
	return nil
}

func (x *EvaluateSSHAccessResponse) GetDenial() *SSHAccessDenial {
	if x != nil {
		if x, ok := x.Decision.(*EvaluateSSHAccessResponse_Denial); ok {
			return x.Denial
		}
	}
	return nil
}

type isEvaluateSSHAccessResponse_Decision interface {
	isEvaluateSSHAccessResponse_Decision()
}

type EvaluateSSHAccessResponse_Permit struct {
	Permit *SSHAccessPermit `protobuf:"bytes,1,opt,name=permit,proto3,oneof"`
}

type EvaluateSSHAccessResponse_Denial struct {
	Denial *SSHAccessDenial `protobuf:"bytes,2,opt,name=denial,proto3,oneof"`
}

func (*EvaluateSSHAccessResponse_Permit) isEvaluateSSHAccessResponse_Decision() {}

func (*EvaluateSSHAccessResponse_Denial) isEvaluateSSHAccessResponse_Decision() {}

// SSHAccessPermit describes the parameters/constraints of a permissible SSH
// access attempt.
type SSHAccessPermit struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	Metadata              *PermitMetadata        `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Logins                []string               `protobuf:"bytes,2,rep,name=logins,proto3" json:"logins,omitempty"`
	ForwardAgent          bool                   `protobuf:"varint,3,opt,name=forward_agent,json=forwardAgent,proto3" json:"forward_agent,omitempty"`
	MaxSessionTtl         *durationpb.Duration   `protobuf:"bytes,4,opt,name=max_session_ttl,json=maxSessionTtl,proto3" json:"max_session_ttl,omitempty"`
	PortForwardMode       SSHPortForwardMode     `protobuf:"varint,5,opt,name=port_forward_mode,json=portForwardMode,proto3,enum=teleport.decision.v1alpha1.SSHPortForwardMode" json:"port_forward_mode,omitempty"`
	ClientIdleTimeout     *durationpb.Duration   `protobuf:"bytes,6,opt,name=client_idle_timeout,json=clientIdleTimeout,proto3" json:"client_idle_timeout,omitempty"`
	DisconnectExpiredCert bool                   `protobuf:"varint,7,opt,name=disconnect_expired_cert,json=disconnectExpiredCert,proto3" json:"disconnect_expired_cert,omitempty"`
	Bpf                   []string               `protobuf:"bytes,8,rep,name=bpf,proto3" json:"bpf,omitempty"`
	X11Forwarding         bool                   `protobuf:"varint,9,opt,name=x11_forwarding,json=x11Forwarding,proto3" json:"x11_forwarding,omitempty"`
	MaxConnections        int64                  `protobuf:"varint,10,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	MaxSessions           int64                  `protobuf:"varint,11,opt,name=max_sessions,json=maxSessions,proto3" json:"max_sessions,omitempty"`
	Lock                  string                 `protobuf:"bytes,12,opt,name=lock,proto3" json:"lock,omitempty"`
	CreateHostUser        bool                   `protobuf:"varint,13,opt,name=create_host_user,json=createHostUser,proto3" json:"create_host_user,omitempty"`
	SshFileCopy           bool                   `protobuf:"varint,14,opt,name=ssh_file_copy,json=sshFileCopy,proto3" json:"ssh_file_copy,omitempty"`
	CreateHostUserMode    string                 `protobuf:"bytes,15,opt,name=create_host_user_mode,json=createHostUserMode,proto3" json:"create_host_user_mode,omitempty"`
	CreateHostUserShell   string                 `protobuf:"bytes,16,opt,name=create_host_user_shell,json=createHostUserShell,proto3" json:"create_host_user_shell,omitempty"`
	HostGroups            []string               `protobuf:"bytes,17,rep,name=host_groups,json=hostGroups,proto3" json:"host_groups,omitempty"`
	HostSudoers           []string               `protobuf:"bytes,18,rep,name=host_sudoers,json=hostSudoers,proto3" json:"host_sudoers,omitempty"`
	SessionRecordingMode  string                 `protobuf:"bytes,19,opt,name=session_recording_mode,json=sessionRecordingMode,proto3" json:"session_recording_mode,omitempty"`
	LockingMode           string                 `protobuf:"bytes,20,opt,name=locking_mode,json=lockingMode,proto3" json:"locking_mode,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *SSHAccessPermit) Reset() {
	*x = SSHAccessPermit{}
	mi := &file_teleport_decision_v1alpha1_ssh_access_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SSHAccessPermit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSHAccessPermit) ProtoMessage() {}

func (x *SSHAccessPermit) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_decision_v1alpha1_ssh_access_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSHAccessPermit.ProtoReflect.Descriptor instead.
func (*SSHAccessPermit) Descriptor() ([]byte, []int) {
	return file_teleport_decision_v1alpha1_ssh_access_proto_rawDescGZIP(), []int{2}
}

func (x *SSHAccessPermit) GetMetadata() *PermitMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *SSHAccessPermit) GetLogins() []string {
	if x != nil {
		return x.Logins
	}
	return nil
}

func (x *SSHAccessPermit) GetForwardAgent() bool {
	if x != nil {
		return x.ForwardAgent
	}
	return false
}

func (x *SSHAccessPermit) GetMaxSessionTtl() *durationpb.Duration {
	if x != nil {
		return x.MaxSessionTtl
	}
	return nil
}

func (x *SSHAccessPermit) GetPortForwardMode() SSHPortForwardMode {
	if x != nil {
		return x.PortForwardMode
	}
	return SSHPortForwardMode_SSH_PORT_FORWARD_MODE_UNSPECIFIED
}

func (x *SSHAccessPermit) GetClientIdleTimeout() *durationpb.Duration {
	if x != nil {
		return x.ClientIdleTimeout
	}
	return nil
}

func (x *SSHAccessPermit) GetDisconnectExpiredCert() bool {
	if x != nil {
		return x.DisconnectExpiredCert
	}
	return false
}

func (x *SSHAccessPermit) GetBpf() []string {
	if x != nil {
		return x.Bpf
	}
	return nil
}

func (x *SSHAccessPermit) GetX11Forwarding() bool {
	if x != nil {
		return x.X11Forwarding
	}
	return false
}

func (x *SSHAccessPermit) GetMaxConnections() int64 {
	if x != nil {
		return x.MaxConnections
	}
	return 0
}

func (x *SSHAccessPermit) GetMaxSessions() int64 {
	if x != nil {
		return x.MaxSessions
	}
	return 0
}

func (x *SSHAccessPermit) GetLock() string {
	if x != nil {
		return x.Lock
	}
	return ""
}

func (x *SSHAccessPermit) GetCreateHostUser() bool {
	if x != nil {
		return x.CreateHostUser
	}
	return false
}

func (x *SSHAccessPermit) GetSshFileCopy() bool {
	if x != nil {
		return x.SshFileCopy
	}
	return false
}

func (x *SSHAccessPermit) GetCreateHostUserMode() string {
	if x != nil {
		return x.CreateHostUserMode
	}
	return ""
}

func (x *SSHAccessPermit) GetCreateHostUserShell() string {
	if x != nil {
		return x.CreateHostUserShell
	}
	return ""
}

func (x *SSHAccessPermit) GetHostGroups() []string {
	if x != nil {
		return x.HostGroups
	}
	return nil
}

func (x *SSHAccessPermit) GetHostSudoers() []string {
	if x != nil {
		return x.HostSudoers
	}
	return nil
}

func (x *SSHAccessPermit) GetSessionRecordingMode() string {
	if x != nil {
		return x.SessionRecordingMode
	}
	return ""
}

func (x *SSHAccessPermit) GetLockingMode() string {
	if x != nil {
		return x.LockingMode
	}
	return ""
}

// SSHAccessDenial describes an SSH access denial.
type SSHAccessDenial struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *DenialMetadata        `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SSHAccessDenial) Reset() {
	*x = SSHAccessDenial{}
	mi := &file_teleport_decision_v1alpha1_ssh_access_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SSHAccessDenial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSHAccessDenial) ProtoMessage() {}

func (x *SSHAccessDenial) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_decision_v1alpha1_ssh_access_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSHAccessDenial.ProtoReflect.Descriptor instead.
func (*SSHAccessDenial) Descriptor() ([]byte, []int) {
	return file_teleport_decision_v1alpha1_ssh_access_proto_rawDescGZIP(), []int{3}
}

func (x *SSHAccessDenial) GetMetadata() *DenialMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_teleport_decision_v1alpha1_ssh_access_proto protoreflect.FileDescriptor

var file_teleport_decision_v1alpha1_ssh_access_proto_rawDesc = string([]byte{
	0x0a, 0x2b, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x73, 0x68,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x65, 0x6e, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x29, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x73, 0x68, 0x5f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x02, 0x0a, 0x18, 0x45,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x53, 0x53, 0x48, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x4d, 0x0a, 0x0d, 0x73, 0x73, 0x68, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x53, 0x48, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x52, 0x0c, 0x73, 0x73, 0x68, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x4a, 0x0a, 0x0c, 0x73, 0x73, 0x68, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x53, 0x53, 0x48, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0b,
	0x73, 0x73, 0x68, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x73, 0x55, 0x73, 0x65, 0x72, 0x22, 0xb5,
	0x01, 0x0a, 0x19, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x53, 0x53, 0x48, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x53, 0x48, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x48, 0x00, 0x52, 0x06, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x74, 0x12, 0x45, 0x0a, 0x06, 0x64, 0x65, 0x6e, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64,
	0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x53, 0x53, 0x48, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x6e, 0x69, 0x61, 0x6c,
	0x48, 0x00, 0x52, 0x06, 0x64, 0x65, 0x6e, 0x69, 0x61, 0x6c, 0x42, 0x0a, 0x0a, 0x08, 0x64, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xa4, 0x07, 0x0a, 0x0f, 0x53, 0x53, 0x48, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12,
	0x41, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x74, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x74, 0x6c, 0x12, 0x5a, 0x0a, 0x11, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x53, 0x48, 0x50, 0x6f,
	0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0f, 0x70,
	0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x49,
	0x0a, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f,
	0x63, 0x65, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x43, 0x65, 0x72,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x70, 0x66, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03,
	0x62, 0x70, 0x66, 0x12, 0x25, 0x0a, 0x0e, 0x78, 0x31, 0x31, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x78, 0x31, 0x31,
	0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61,
	0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x73, 0x68, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x63, 0x6f, 0x70, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x73, 0x68,
	0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x70, 0x79, 0x12, 0x31, 0x0a, 0x15, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x68, 0x65, 0x6c, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x65, 0x6c, 0x6c,
	0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18,
	0x11, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x75, 0x64, 0x6f, 0x65, 0x72,
	0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x73, 0x74, 0x53, 0x75, 0x64,
	0x6f, 0x65, 0x72, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x59, 0x0a,
	0x0f, 0x53, 0x53, 0x48, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x6e, 0x69, 0x61, 0x6c,
	0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x44, 0x65, 0x6e, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2a, 0xbb, 0x01, 0x0a, 0x12, 0x53, 0x53, 0x48,
	0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x25, 0x0a, 0x21, 0x53, 0x53, 0x48, 0x5f, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x57,
	0x41, 0x52, 0x44, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x53, 0x48, 0x5f, 0x50, 0x4f,
	0x52, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f,
	0x4f, 0x46, 0x46, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x53, 0x48, 0x5f, 0x50, 0x4f, 0x52,
	0x54, 0x5f, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x4f,
	0x4e, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x53, 0x48, 0x5f, 0x50, 0x4f, 0x52, 0x54, 0x5f,
	0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x4c, 0x4f, 0x43,
	0x41, 0x4c, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x53, 0x48, 0x5f, 0x50, 0x4f, 0x52, 0x54,
	0x5f, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45,
	0x4d, 0x4f, 0x54, 0x45, 0x10, 0x04, 0x42, 0x5a, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_teleport_decision_v1alpha1_ssh_access_proto_rawDescOnce sync.Once
	file_teleport_decision_v1alpha1_ssh_access_proto_rawDescData []byte
)

func file_teleport_decision_v1alpha1_ssh_access_proto_rawDescGZIP() []byte {
	file_teleport_decision_v1alpha1_ssh_access_proto_rawDescOnce.Do(func() {
		file_teleport_decision_v1alpha1_ssh_access_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_teleport_decision_v1alpha1_ssh_access_proto_rawDesc), len(file_teleport_decision_v1alpha1_ssh_access_proto_rawDesc)))
	})
	return file_teleport_decision_v1alpha1_ssh_access_proto_rawDescData
}

var file_teleport_decision_v1alpha1_ssh_access_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_teleport_decision_v1alpha1_ssh_access_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_teleport_decision_v1alpha1_ssh_access_proto_goTypes = []any{
	(SSHPortForwardMode)(0),           // 0: teleport.decision.v1alpha1.SSHPortForwardMode
	(*EvaluateSSHAccessRequest)(nil),  // 1: teleport.decision.v1alpha1.EvaluateSSHAccessRequest
	(*EvaluateSSHAccessResponse)(nil), // 2: teleport.decision.v1alpha1.EvaluateSSHAccessResponse
	(*SSHAccessPermit)(nil),           // 3: teleport.decision.v1alpha1.SSHAccessPermit
	(*SSHAccessDenial)(nil),           // 4: teleport.decision.v1alpha1.SSHAccessDenial
	(*RequestMetadata)(nil),           // 5: teleport.decision.v1alpha1.RequestMetadata
	(*SSHAuthority)(nil),              // 6: teleport.decision.v1alpha1.SSHAuthority
	(*SSHIdentity)(nil),               // 7: teleport.decision.v1alpha1.SSHIdentity
	(*Resource)(nil),                  // 8: teleport.decision.v1alpha1.Resource
	(*PermitMetadata)(nil),            // 9: teleport.decision.v1alpha1.PermitMetadata
	(*durationpb.Duration)(nil),       // 10: google.protobuf.Duration
	(*DenialMetadata)(nil),            // 11: teleport.decision.v1alpha1.DenialMetadata
}
var file_teleport_decision_v1alpha1_ssh_access_proto_depIdxs = []int32{
	5,  // 0: teleport.decision.v1alpha1.EvaluateSSHAccessRequest.metadata:type_name -> teleport.decision.v1alpha1.RequestMetadata
	6,  // 1: teleport.decision.v1alpha1.EvaluateSSHAccessRequest.ssh_authority:type_name -> teleport.decision.v1alpha1.SSHAuthority
	7,  // 2: teleport.decision.v1alpha1.EvaluateSSHAccessRequest.ssh_identity:type_name -> teleport.decision.v1alpha1.SSHIdentity
	8,  // 3: teleport.decision.v1alpha1.EvaluateSSHAccessRequest.node:type_name -> teleport.decision.v1alpha1.Resource
	3,  // 4: teleport.decision.v1alpha1.EvaluateSSHAccessResponse.permit:type_name -> teleport.decision.v1alpha1.SSHAccessPermit
	4,  // 5: teleport.decision.v1alpha1.EvaluateSSHAccessResponse.denial:type_name -> teleport.decision.v1alpha1.SSHAccessDenial
	9,  // 6: teleport.decision.v1alpha1.SSHAccessPermit.metadata:type_name -> teleport.decision.v1alpha1.PermitMetadata
	10, // 7: teleport.decision.v1alpha1.SSHAccessPermit.max_session_ttl:type_name -> google.protobuf.Duration
	0,  // 8: teleport.decision.v1alpha1.SSHAccessPermit.port_forward_mode:type_name -> teleport.decision.v1alpha1.SSHPortForwardMode
	10, // 9: teleport.decision.v1alpha1.SSHAccessPermit.client_idle_timeout:type_name -> google.protobuf.Duration
	11, // 10: teleport.decision.v1alpha1.SSHAccessDenial.metadata:type_name -> teleport.decision.v1alpha1.DenialMetadata
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_teleport_decision_v1alpha1_ssh_access_proto_init() }
func file_teleport_decision_v1alpha1_ssh_access_proto_init() {
	if File_teleport_decision_v1alpha1_ssh_access_proto != nil {
		return
	}
	file_teleport_decision_v1alpha1_denial_metadata_proto_init()
	file_teleport_decision_v1alpha1_permit_metadata_proto_init()
	file_teleport_decision_v1alpha1_request_metadata_proto_init()
	file_teleport_decision_v1alpha1_resource_proto_init()
	file_teleport_decision_v1alpha1_ssh_identity_proto_init()
	file_teleport_decision_v1alpha1_ssh_access_proto_msgTypes[1].OneofWrappers = []any{
		(*EvaluateSSHAccessResponse_Permit)(nil),
		(*EvaluateSSHAccessResponse_Denial)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_teleport_decision_v1alpha1_ssh_access_proto_rawDesc), len(file_teleport_decision_v1alpha1_ssh_access_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_decision_v1alpha1_ssh_access_proto_goTypes,
		DependencyIndexes: file_teleport_decision_v1alpha1_ssh_access_proto_depIdxs,
		EnumInfos:         file_teleport_decision_v1alpha1_ssh_access_proto_enumTypes,
		MessageInfos:      file_teleport_decision_v1alpha1_ssh_access_proto_msgTypes,
	}.Build()
	File_teleport_decision_v1alpha1_ssh_access_proto = out.File
	file_teleport_decision_v1alpha1_ssh_access_proto_goTypes = nil
	file_teleport_decision_v1alpha1_ssh_access_proto_depIdxs = nil
}
