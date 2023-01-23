// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: workspace_messages.proto

package pb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Workspace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OfficeId uint32 `protobuf:"varint,1,opt,name=office_id,json=officeId,proto3" json:"office_id,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Workspace) Reset() {
	*x = Workspace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workspace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workspace) ProtoMessage() {}

func (x *Workspace) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workspace.ProtoReflect.Descriptor instead.
func (*Workspace) Descriptor() ([]byte, []int) {
	return file_workspace_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Workspace) GetOfficeId() uint32 {
	if x != nil {
		return x.OfficeId
	}
	return 0
}

func (x *Workspace) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Create Workspace
type CreateWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OfficeId uint32 `protobuf:"varint,1,opt,name=office_id,json=officeId,proto3" json:"office_id,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateWorkspaceRequest) Reset() {
	*x = CreateWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkspaceRequest) ProtoMessage() {}

func (x *CreateWorkspaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*CreateWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_workspace_messages_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWorkspaceRequest) GetOfficeId() uint32 {
	if x != nil {
		return x.OfficeId
	}
	return 0
}

func (x *CreateWorkspaceRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateWorkspaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspace *Workspace `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
}

func (x *CreateWorkspaceResponse) Reset() {
	*x = CreateWorkspaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWorkspaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkspaceResponse) ProtoMessage() {}

func (x *CreateWorkspaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkspaceResponse.ProtoReflect.Descriptor instead.
func (*CreateWorkspaceResponse) Descriptor() ([]byte, []int) {
	return file_workspace_messages_proto_rawDescGZIP(), []int{2}
}

func (x *CreateWorkspaceResponse) GetWorkspace() *Workspace {
	if x != nil {
		return x.Workspace
	}
	return nil
}

// Delete Workspace By ID
type DeleteWorkspaceByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteWorkspaceByIdRequest) Reset() {
	*x = DeleteWorkspaceByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWorkspaceByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkspaceByIdRequest) ProtoMessage() {}

func (x *DeleteWorkspaceByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkspaceByIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteWorkspaceByIdRequest) Descriptor() ([]byte, []int) {
	return file_workspace_messages_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteWorkspaceByIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DeleteWorkspaceByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteWorkspaceByIdResponse) Reset() {
	*x = DeleteWorkspaceByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWorkspaceByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkspaceByIdResponse) ProtoMessage() {}

func (x *DeleteWorkspaceByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkspaceByIdResponse.ProtoReflect.Descriptor instead.
func (*DeleteWorkspaceByIdResponse) Descriptor() ([]byte, []int) {
	return file_workspace_messages_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteWorkspaceByIdResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Find Workspaces
type FindWorkspacesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OfficeId uint32 `protobuf:"varint,1,opt,name=office_id,json=officeId,proto3" json:"office_id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Role     string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	Verified string `protobuf:"bytes,5,opt,name=verified,proto3" json:"verified,omitempty"`
	Page     uint32 `protobuf:"varint,6,opt,name=page,proto3" json:"page,omitempty"`
	Limit    uint32 `protobuf:"varint,7,opt,name=limit,proto3" json:"limit,omitempty"`
	Sort     string `protobuf:"bytes,8,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *FindWorkspacesRequest) Reset() {
	*x = FindWorkspacesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindWorkspacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindWorkspacesRequest) ProtoMessage() {}

func (x *FindWorkspacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindWorkspacesRequest.ProtoReflect.Descriptor instead.
func (*FindWorkspacesRequest) Descriptor() ([]byte, []int) {
	return file_workspace_messages_proto_rawDescGZIP(), []int{5}
}

func (x *FindWorkspacesRequest) GetOfficeId() uint32 {
	if x != nil {
		return x.OfficeId
	}
	return 0
}

func (x *FindWorkspacesRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *FindWorkspacesRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *FindWorkspacesRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *FindWorkspacesRequest) GetVerified() string {
	if x != nil {
		return x.Verified
	}
	return ""
}

func (x *FindWorkspacesRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindWorkspacesRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindWorkspacesRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type FindWorkspacesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount uint32  `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	TotalPages uint32  `protobuf:"varint,2,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	Page       uint32  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit      uint32  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	HasMore    bool    `protobuf:"varint,5,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	Users      []*User `protobuf:"bytes,6,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *FindWorkspacesResponse) Reset() {
	*x = FindWorkspacesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindWorkspacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindWorkspacesResponse) ProtoMessage() {}

func (x *FindWorkspacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindWorkspacesResponse.ProtoReflect.Descriptor instead.
func (*FindWorkspacesResponse) Descriptor() ([]byte, []int) {
	return file_workspace_messages_proto_rawDescGZIP(), []int{6}
}

func (x *FindWorkspacesResponse) GetTotalCount() uint32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *FindWorkspacesResponse) GetTotalPages() uint32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *FindWorkspacesResponse) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindWorkspacesResponse) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindWorkspacesResponse) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (x *FindWorkspacesResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_workspace_messages_proto protoreflect.FileDescriptor

var file_workspace_messages_proto_rawDesc = []byte{
	0x0a, 0x18, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x09,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x09, 0x6f, 0x66, 0x66,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x08, 0x92, 0x41,
	0x05, 0x4a, 0x03, 0x22, 0x31, 0x22, 0x52, 0x08, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x4b, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x32, 0x92, 0x41, 0x2f, 0x4a, 0x26, 0x22, 0x32, 0x34, 0x33, 0x38, 0x61, 0x63, 0x33,
	0x63, 0x2d, 0x33, 0x37, 0x65, 0x62, 0x2d, 0x34, 0x39, 0x30, 0x32, 0x2d, 0x61, 0x64, 0x65, 0x66,
	0x2d, 0x65, 0x64, 0x31, 0x36, 0x62, 0x34, 0x34, 0x33, 0x31, 0x30, 0x33, 0x30, 0x22, 0xa2, 0x02,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x9b, 0x01,
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x09, 0x6f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0f, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x20, 0x00, 0x92, 0x41, 0x05, 0x4a, 0x03, 0x22, 0x31, 0x22, 0x52, 0x08, 0x6f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x53, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01,
	0x01, 0x92, 0x41, 0x2f, 0x4a, 0x26, 0x22, 0x32, 0x34, 0x33, 0x38, 0x61, 0x63, 0x33, 0x63, 0x2d,
	0x33, 0x37, 0x65, 0x62, 0x2d, 0x34, 0x39, 0x30, 0x32, 0x2d, 0x61, 0x64, 0x65, 0x66, 0x2d, 0x65,
	0x64, 0x31, 0x36, 0x62, 0x34, 0x34, 0x33, 0x31, 0x30, 0x33, 0x30, 0x22, 0xa2, 0x02, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x22, 0x71, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x53, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x3a, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x92, 0x41, 0x2f, 0x4a,
	0x26, 0x22, 0x32, 0x34, 0x33, 0x38, 0x61, 0x63, 0x33, 0x63, 0x2d, 0x33, 0x37, 0x65, 0x62, 0x2d,
	0x34, 0x39, 0x30, 0x32, 0x2d, 0x61, 0x64, 0x65, 0x66, 0x2d, 0x65, 0x64, 0x31, 0x36, 0x62, 0x34,
	0x34, 0x33, 0x31, 0x30, 0x33, 0x30, 0x22, 0xa2, 0x02, 0x04, 0x75, 0x75, 0x69, 0x64, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5e, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0x92, 0x41, 0x22, 0x4a, 0x20, 0x22, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x20,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x22, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xe2, 0x02, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x09, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x08, 0x92, 0x41, 0x05, 0x4a, 0x03, 0x22, 0x31, 0x22, 0x52, 0x08, 0x6f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0x92, 0x41, 0x0b, 0x4a, 0x09,
	0x22, 0x6a, 0x6f, 0x68, 0x6e, 0x64, 0x65, 0x65, 0x22, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x18, 0x92, 0x41, 0x15, 0x4a, 0x13, 0x22, 0x6a, 0x6f, 0x68, 0x6e, 0x64, 0x65,
	0x65, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x22, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0b, 0x92, 0x41, 0x08, 0x4a, 0x06, 0x22, 0x75, 0x73, 0x65, 0x72, 0x22, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0x92, 0x41, 0x08, 0x4a, 0x06, 0x22, 0x74, 0x72,
	0x75, 0x65, 0x22, 0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1c, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x08, 0x92, 0x41, 0x05,
	0x4a, 0x03, 0x22, 0x31, 0x22, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0c, 0x92, 0x41, 0x09, 0x3a,
	0x01, 0x35, 0x4a, 0x04, 0x22, 0x31, 0x30, 0x22, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x3a, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0x92,
	0x41, 0x23, 0x3a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x3a, 0x64,
	0x65, 0x73, 0x63, 0x4a, 0x10, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x3a, 0x61, 0x73, 0x63, 0x22, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0xf5, 0x01, 0x0a, 0x16,
	0x46, 0x69, 0x6e, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0x92, 0x41, 0x06,
	0x4a, 0x04, 0x22, 0x31, 0x30, 0x22, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x29, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x08, 0x92, 0x41, 0x05, 0x4a, 0x03, 0x22, 0x32,
	0x22, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1c, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x08, 0x92, 0x41, 0x05,
	0x4a, 0x03, 0x22, 0x31, 0x22, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x08, 0x92, 0x41, 0x05, 0x4a,
	0x03, 0x22, 0x35, 0x22, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x08, 0x68,
	0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0b, 0x92,
	0x41, 0x08, 0x4a, 0x06, 0x22, 0x74, 0x72, 0x75, 0x65, 0x22, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d,
	0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workspace_messages_proto_rawDescOnce sync.Once
	file_workspace_messages_proto_rawDescData = file_workspace_messages_proto_rawDesc
)

func file_workspace_messages_proto_rawDescGZIP() []byte {
	file_workspace_messages_proto_rawDescOnce.Do(func() {
		file_workspace_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_workspace_messages_proto_rawDescData)
	})
	return file_workspace_messages_proto_rawDescData
}

var file_workspace_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_workspace_messages_proto_goTypes = []interface{}{
	(*Workspace)(nil),                   // 0: pb.Workspace
	(*CreateWorkspaceRequest)(nil),      // 1: pb.CreateWorkspaceRequest
	(*CreateWorkspaceResponse)(nil),     // 2: pb.CreateWorkspaceResponse
	(*DeleteWorkspaceByIdRequest)(nil),  // 3: pb.DeleteWorkspaceByIdRequest
	(*DeleteWorkspaceByIdResponse)(nil), // 4: pb.DeleteWorkspaceByIdResponse
	(*FindWorkspacesRequest)(nil),       // 5: pb.FindWorkspacesRequest
	(*FindWorkspacesResponse)(nil),      // 6: pb.FindWorkspacesResponse
	(*User)(nil),                        // 7: pb.User
}
var file_workspace_messages_proto_depIdxs = []int32{
	0, // 0: pb.CreateWorkspaceResponse.workspace:type_name -> pb.Workspace
	7, // 1: pb.FindWorkspacesResponse.users:type_name -> pb.User
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_workspace_messages_proto_init() }
func file_workspace_messages_proto_init() {
	if File_workspace_messages_proto != nil {
		return
	}
	file_user_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_workspace_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workspace); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workspace_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWorkspaceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workspace_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWorkspaceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workspace_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWorkspaceByIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workspace_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWorkspaceByIdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workspace_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindWorkspacesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workspace_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindWorkspacesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_workspace_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_workspace_messages_proto_goTypes,
		DependencyIndexes: file_workspace_messages_proto_depIdxs,
		MessageInfos:      file_workspace_messages_proto_msgTypes,
	}.Build()
	File_workspace_messages_proto = out.File
	file_workspace_messages_proto_rawDesc = nil
	file_workspace_messages_proto_goTypes = nil
	file_workspace_messages_proto_depIdxs = nil
}