// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: check.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_check_proto protoreflect.FileDescriptor

var file_check_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf9, 0x04, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa4, 0x01, 0x0a, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x70, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1e, 0x32, 0x19, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x2f, 0x7b,
	0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x01,
	0x2a, 0x92, 0x41, 0x49, 0x12, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x20, 0x69, 0x6e, 0x20, 0x62, 0x79, 0x20, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61,
	0x6e, 0x63, 0x65, 0x20, 0x69, 0x64, 0x1a, 0x25, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x20, 0x69, 0x6e, 0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x12, 0xa7, 0x01,
	0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x70, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x32, 0x1a, 0x2f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x2d, 0x6f, 0x75, 0x74, 0x2f, 0x7b, 0x61, 0x74, 0x74, 0x65, 0x6e,
	0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x92, 0x41, 0x4b, 0x12, 0x21, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x62,
	0x79, 0x20, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x20, 0x69, 0x64, 0x1a,
	0x26, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f,
	0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x6f, 0x75,
	0x74, 0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x12, 0x79, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x73, 0x92, 0x41, 0x2a, 0x12, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x20, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x73, 0x1a, 0x1b, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20,
	0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x66, 0x69, 0x6e, 0x64, 0x20, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x73, 0x12, 0x9c, 0x01, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x12, 0x17, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2f, 0x7b, 0x61, 0x74, 0x74, 0x65,
	0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x92, 0x41, 0x34, 0x12, 0x10, 0x46,
	0x69, 0x6e, 0x64, 0x20, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x42, 0x79, 0x20, 0x49, 0x64, 0x1a,
	0x20, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f,
	0x20, 0x66, 0x69, 0x6e, 0x64, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x62, 0x79, 0x20, 0x69,
	0x64, 0x42, 0x87, 0x01, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x92, 0x41, 0x7d, 0x12, 0x12,
	0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x41, 0x50, 0x49, 0x32, 0x05, 0x30, 0x2e, 0x30,
	0x2e, 0x31, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x4d,
	0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a,
	0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_check_proto_goTypes = []interface{}{
	(*CheckInRequest)(nil),        // 0: pb.CheckInRequest
	(*CheckOutRequest)(nil),       // 1: pb.CheckOutRequest
	(*FindChecksRequest)(nil),     // 2: pb.FindChecksRequest
	(*FindCheckByIdRequest)(nil),  // 3: pb.FindCheckByIdRequest
	(*CheckInResponse)(nil),       // 4: pb.CheckInResponse
	(*CheckOutResponse)(nil),      // 5: pb.CheckOutResponse
	(*FindChecksResponse)(nil),    // 6: pb.FindChecksResponse
	(*FindCheckByIdResponse)(nil), // 7: pb.FindCheckByIdResponse
}
var file_check_proto_depIdxs = []int32{
	0, // 0: pb.CheckService.CheckIn:input_type -> pb.CheckInRequest
	1, // 1: pb.CheckService.CheckOut:input_type -> pb.CheckOutRequest
	2, // 2: pb.CheckService.FindChecks:input_type -> pb.FindChecksRequest
	3, // 3: pb.CheckService.FindCheckById:input_type -> pb.FindCheckByIdRequest
	4, // 4: pb.CheckService.CheckIn:output_type -> pb.CheckInResponse
	5, // 5: pb.CheckService.CheckOut:output_type -> pb.CheckOutResponse
	6, // 6: pb.CheckService.FindChecks:output_type -> pb.FindChecksResponse
	7, // 7: pb.CheckService.FindCheckById:output_type -> pb.FindCheckByIdResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_check_proto_init() }
func file_check_proto_init() {
	if File_check_proto != nil {
		return
	}
	file_check_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_check_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_check_proto_goTypes,
		DependencyIndexes: file_check_proto_depIdxs,
	}.Build()
	File_check_proto = out.File
	file_check_proto_rawDesc = nil
	file_check_proto_goTypes = nil
	file_check_proto_depIdxs = nil
}