// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: role_messages.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Role      string                 `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_role_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_role_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Role) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Role) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Role) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Create Role
type CreateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *CreateRoleRequest) Reset() {
	*x = CreateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequest) ProtoMessage() {}

func (x *CreateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return file_role_messages_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoleRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type CreateRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role *Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *CreateRoleResponse) Reset() {
	*x = CreateRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleResponse) ProtoMessage() {}

func (x *CreateRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleResponse.ProtoReflect.Descriptor instead.
func (*CreateRoleResponse) Descriptor() ([]byte, []int) {
	return file_role_messages_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRoleResponse) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

// Delete Role By ID
type DeleteRoleByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRoleByIdRequest) Reset() {
	*x = DeleteRoleByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleByIdRequest) ProtoMessage() {}

func (x *DeleteRoleByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleByIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoleByIdRequest) Descriptor() ([]byte, []int) {
	return file_role_messages_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRoleByIdRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteRoleByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteRoleByIdResponse) Reset() {
	*x = DeleteRoleByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleByIdResponse) ProtoMessage() {}

func (x *DeleteRoleByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleByIdResponse.ProtoReflect.Descriptor instead.
func (*DeleteRoleByIdResponse) Descriptor() ([]byte, []int) {
	return file_role_messages_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRoleByIdResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Find Roles
type FindRolesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role  string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Page  uint32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Sort  string `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *FindRolesRequest) Reset() {
	*x = FindRolesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRolesRequest) ProtoMessage() {}

func (x *FindRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRolesRequest.ProtoReflect.Descriptor instead.
func (*FindRolesRequest) Descriptor() ([]byte, []int) {
	return file_role_messages_proto_rawDescGZIP(), []int{5}
}

func (x *FindRolesRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *FindRolesRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindRolesRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindRolesRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type FindRolesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount uint32  `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	TotalPages uint32  `protobuf:"varint,2,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	Page       uint32  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit      uint32  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	HasMore    bool    `protobuf:"varint,5,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	Roles      []*Role `protobuf:"bytes,6,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *FindRolesResponse) Reset() {
	*x = FindRolesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRolesResponse) ProtoMessage() {}

func (x *FindRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRolesResponse.ProtoReflect.Descriptor instead.
func (*FindRolesResponse) Descriptor() ([]byte, []int) {
	return file_role_messages_proto_rawDescGZIP(), []int{6}
}

func (x *FindRolesResponse) GetTotalCount() uint32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *FindRolesResponse) GetTotalPages() uint32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *FindRolesResponse) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindRolesResponse) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindRolesResponse) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (x *FindRolesResponse) GetRoles() []*Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

var File_role_messages_proto protoreflect.FileDescriptor

var file_role_messages_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x04, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x08, 0x92, 0x41, 0x05, 0x4a, 0x03, 0x22, 0x31, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0x92, 0x41, 0x08,
	0x4a, 0x06, 0x22, 0x75, 0x73, 0x65, 0x72, 0x22, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x34, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0x92, 0x41, 0x08, 0x4a, 0x06, 0x22, 0x75,
	0x73, 0x65, 0x72, 0x22, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x32, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x31,
	0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x08, 0x92, 0x41, 0x05, 0x4a, 0x03, 0x22, 0x31, 0x22, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x54, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x20, 0x92, 0x41,
	0x1d, 0x4a, 0x1b, 0x22, 0x52, 0x6f, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x22, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0x92, 0x41, 0x08, 0x4a,
	0x06, 0x22, 0x75, 0x73, 0x65, 0x72, 0x22, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1c, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x08, 0x92, 0x41, 0x05,
	0x4a, 0x03, 0x22, 0x31, 0x22, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0c, 0x92, 0x41, 0x09, 0x3a,
	0x01, 0x35, 0x4a, 0x04, 0x22, 0x31, 0x30, 0x22, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x3a, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0x92,
	0x41, 0x23, 0x3a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x3a, 0x64,
	0x65, 0x73, 0x63, 0x4a, 0x10, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x3a, 0x61, 0x73, 0x63, 0x22, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0xf0, 0x01, 0x0a, 0x11,
	0x46, 0x69, 0x6e, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0x92, 0x41, 0x06, 0x4a, 0x04, 0x22, 0x31, 0x30,
	0x22, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x08, 0x92, 0x41, 0x05, 0x4a, 0x03, 0x22, 0x32, 0x22, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x08, 0x92, 0x41, 0x05, 0x4a, 0x03, 0x22, 0x31, 0x22,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x08, 0x92, 0x41, 0x05, 0x4a, 0x03, 0x22, 0x35, 0x22, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f,
	0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0b, 0x92, 0x41, 0x08, 0x4a, 0x06, 0x22,
	0x74, 0x72, 0x75, 0x65, 0x22, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x12, 0x1e,
	0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_messages_proto_rawDescOnce sync.Once
	file_role_messages_proto_rawDescData = file_role_messages_proto_rawDesc
)

func file_role_messages_proto_rawDescGZIP() []byte {
	file_role_messages_proto_rawDescOnce.Do(func() {
		file_role_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_messages_proto_rawDescData)
	})
	return file_role_messages_proto_rawDescData
}

var file_role_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_role_messages_proto_goTypes = []interface{}{
	(*Role)(nil),                   // 0: pb.Role
	(*CreateRoleRequest)(nil),      // 1: pb.CreateRoleRequest
	(*CreateRoleResponse)(nil),     // 2: pb.CreateRoleResponse
	(*DeleteRoleByIdRequest)(nil),  // 3: pb.DeleteRoleByIdRequest
	(*DeleteRoleByIdResponse)(nil), // 4: pb.DeleteRoleByIdResponse
	(*FindRolesRequest)(nil),       // 5: pb.FindRolesRequest
	(*FindRolesResponse)(nil),      // 6: pb.FindRolesResponse
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
}
var file_role_messages_proto_depIdxs = []int32{
	7, // 0: pb.Role.created_at:type_name -> google.protobuf.Timestamp
	7, // 1: pb.Role.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: pb.CreateRoleResponse.role:type_name -> pb.Role
	0, // 3: pb.FindRolesResponse.roles:type_name -> pb.Role
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_role_messages_proto_init() }
func file_role_messages_proto_init() {
	if File_role_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_role_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_role_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleRequest); i {
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
		file_role_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleResponse); i {
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
		file_role_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleByIdRequest); i {
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
		file_role_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleByIdResponse); i {
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
		file_role_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRolesRequest); i {
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
		file_role_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRolesResponse); i {
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
			RawDescriptor: file_role_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_role_messages_proto_goTypes,
		DependencyIndexes: file_role_messages_proto_depIdxs,
		MessageInfos:      file_role_messages_proto_msgTypes,
	}.Build()
	File_role_messages_proto = out.File
	file_role_messages_proto_rawDesc = nil
	file_role_messages_proto_goTypes = nil
	file_role_messages_proto_depIdxs = nil
}