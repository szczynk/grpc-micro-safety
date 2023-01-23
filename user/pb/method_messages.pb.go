// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: method_messages.proto

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

type Method struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Method    string                 `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Method) Reset() {
	*x = Method{}
	if protoimpl.UnsafeEnabled {
		mi := &file_method_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Method) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Method) ProtoMessage() {}

func (x *Method) ProtoReflect() protoreflect.Message {
	mi := &file_method_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Method.ProtoReflect.Descriptor instead.
func (*Method) Descriptor() ([]byte, []int) {
	return file_method_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Method) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Method) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Method) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Method) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Create Method
type CreateMethodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *CreateMethodRequest) Reset() {
	*x = CreateMethodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_method_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMethodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMethodRequest) ProtoMessage() {}

func (x *CreateMethodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_method_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMethodRequest.ProtoReflect.Descriptor instead.
func (*CreateMethodRequest) Descriptor() ([]byte, []int) {
	return file_method_messages_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMethodRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type CreateMethodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method *Method `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *CreateMethodResponse) Reset() {
	*x = CreateMethodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_method_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMethodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMethodResponse) ProtoMessage() {}

func (x *CreateMethodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_method_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMethodResponse.ProtoReflect.Descriptor instead.
func (*CreateMethodResponse) Descriptor() ([]byte, []int) {
	return file_method_messages_proto_rawDescGZIP(), []int{2}
}

func (x *CreateMethodResponse) GetMethod() *Method {
	if x != nil {
		return x.Method
	}
	return nil
}

// Delete Method By ID
type DeleteMethodByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMethodByIdRequest) Reset() {
	*x = DeleteMethodByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_method_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMethodByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMethodByIdRequest) ProtoMessage() {}

func (x *DeleteMethodByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_method_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMethodByIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteMethodByIdRequest) Descriptor() ([]byte, []int) {
	return file_method_messages_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteMethodByIdRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteMethodByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteMethodByIdResponse) Reset() {
	*x = DeleteMethodByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_method_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMethodByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMethodByIdResponse) ProtoMessage() {}

func (x *DeleteMethodByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_method_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMethodByIdResponse.ProtoReflect.Descriptor instead.
func (*DeleteMethodByIdResponse) Descriptor() ([]byte, []int) {
	return file_method_messages_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteMethodByIdResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Find Methods
type FindMethodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Page   uint32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit  uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Sort   string `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *FindMethodsRequest) Reset() {
	*x = FindMethodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_method_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMethodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMethodsRequest) ProtoMessage() {}

func (x *FindMethodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_method_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMethodsRequest.ProtoReflect.Descriptor instead.
func (*FindMethodsRequest) Descriptor() ([]byte, []int) {
	return file_method_messages_proto_rawDescGZIP(), []int{5}
}

func (x *FindMethodsRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *FindMethodsRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindMethodsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindMethodsRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type FindMethodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount uint32    `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	TotalPages uint32    `protobuf:"varint,2,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	Page       uint32    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit      uint32    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	HasMore    bool      `protobuf:"varint,5,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	Methods    []*Method `protobuf:"bytes,6,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (x *FindMethodsResponse) Reset() {
	*x = FindMethodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_method_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMethodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMethodsResponse) ProtoMessage() {}

func (x *FindMethodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_method_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMethodsResponse.ProtoReflect.Descriptor instead.
func (*FindMethodsResponse) Descriptor() ([]byte, []int) {
	return file_method_messages_proto_rawDescGZIP(), []int{6}
}

func (x *FindMethodsResponse) GetTotalCount() uint32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *FindMethodsResponse) GetTotalPages() uint32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *FindMethodsResponse) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindMethodsResponse) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindMethodsResponse) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (x *FindMethodsResponse) GetMethods() []*Method {
	if x != nil {
		return x.Methods
	}
	return nil
}

var File_method_messages_proto protoreflect.FileDescriptor

var file_method_messages_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a,
	0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x08, 0x92, 0x41, 0x05, 0x4a, 0x03, 0x22, 0x31, 0x22, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2b, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x13, 0x92, 0x41, 0x10, 0x4a, 0x0e, 0x22, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x2d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x22, 0x3a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22,
	0x33, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x08, 0x92, 0x41, 0x05, 0x4a, 0x03, 0x22, 0x31, 0x22,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x58, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3c, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x22, 0x92, 0x41, 0x1f, 0x4a, 0x1d, 0x22, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x20,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66,
	0x75, 0x6c, 0x6c, 0x79, 0x22, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xbf,
	0x01, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0x92, 0x41, 0x10, 0x4a, 0x0e, 0x22, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x08, 0x92, 0x41, 0x05, 0x4a, 0x03, 0x22, 0x31, 0x22, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x22, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x0c, 0x92, 0x41, 0x09, 0x3a, 0x01, 0x35, 0x4a, 0x04, 0x22, 0x31, 0x30, 0x22, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x3a, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x26, 0x92, 0x41, 0x23, 0x3a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x3a, 0x64, 0x65, 0x73, 0x63, 0x4a, 0x10, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x3a, 0x61, 0x73, 0x63, 0x22, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x22, 0xf8, 0x01, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0x92,
	0x41, 0x06, 0x4a, 0x04, 0x22, 0x31, 0x30, 0x22, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x08, 0x92, 0x41, 0x05, 0x4a, 0x03,
	0x22, 0x32, 0x22, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x08, 0x92,
	0x41, 0x05, 0x4a, 0x03, 0x22, 0x31, 0x22, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x08, 0x92, 0x41,
	0x05, 0x4a, 0x03, 0x22, 0x35, 0x22, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x26, 0x0a,
	0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x0b, 0x92, 0x41, 0x08, 0x4a, 0x06, 0x22, 0x74, 0x72, 0x75, 0x65, 0x22, 0x52, 0x07, 0x68, 0x61,
	0x73, 0x4d, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x42, 0x07, 0x5a, 0x05, 0x2e,
	0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_method_messages_proto_rawDescOnce sync.Once
	file_method_messages_proto_rawDescData = file_method_messages_proto_rawDesc
)

func file_method_messages_proto_rawDescGZIP() []byte {
	file_method_messages_proto_rawDescOnce.Do(func() {
		file_method_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_method_messages_proto_rawDescData)
	})
	return file_method_messages_proto_rawDescData
}

var file_method_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_method_messages_proto_goTypes = []interface{}{
	(*Method)(nil),                   // 0: pb.Method
	(*CreateMethodRequest)(nil),      // 1: pb.CreateMethodRequest
	(*CreateMethodResponse)(nil),     // 2: pb.CreateMethodResponse
	(*DeleteMethodByIdRequest)(nil),  // 3: pb.DeleteMethodByIdRequest
	(*DeleteMethodByIdResponse)(nil), // 4: pb.DeleteMethodByIdResponse
	(*FindMethodsRequest)(nil),       // 5: pb.FindMethodsRequest
	(*FindMethodsResponse)(nil),      // 6: pb.FindMethodsResponse
	(*timestamppb.Timestamp)(nil),    // 7: google.protobuf.Timestamp
}
var file_method_messages_proto_depIdxs = []int32{
	7, // 0: pb.Method.created_at:type_name -> google.protobuf.Timestamp
	7, // 1: pb.Method.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: pb.CreateMethodResponse.method:type_name -> pb.Method
	0, // 3: pb.FindMethodsResponse.methods:type_name -> pb.Method
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_method_messages_proto_init() }
func file_method_messages_proto_init() {
	if File_method_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_method_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Method); i {
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
		file_method_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMethodRequest); i {
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
		file_method_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMethodResponse); i {
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
		file_method_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMethodByIdRequest); i {
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
		file_method_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMethodByIdResponse); i {
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
		file_method_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMethodsRequest); i {
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
		file_method_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMethodsResponse); i {
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
			RawDescriptor: file_method_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_method_messages_proto_goTypes,
		DependencyIndexes: file_method_messages_proto_depIdxs,
		MessageInfos:      file_method_messages_proto_msgTypes,
	}.Build()
	File_method_messages_proto = out.File
	file_method_messages_proto_rawDesc = nil
	file_method_messages_proto_goTypes = nil
	file_method_messages_proto_depIdxs = nil
}
