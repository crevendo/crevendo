// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/service/field.proto

package grpc

import (
	message "github.com/crevendo/crevendo/message"
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

var File_proto_service_field_proto protoreflect.FileDescriptor

var file_proto_service_field_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x85, 0x01, 0x0a, 0x0c, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x23,
	0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x65,
	0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_service_field_proto_goTypes = []interface{}{
	(*message.FieldListMessage)(nil),  // 0: FieldListMessage
	(*message.FieldListResponse)(nil), // 1: FieldListResponse
}
var file_proto_service_field_proto_depIdxs = []int32{
	0, // 0: FieldService.UserFieldList:input_type -> FieldListMessage
	0, // 1: FieldService.AddressFieldList:input_type -> FieldListMessage
	1, // 2: FieldService.UserFieldList:output_type -> FieldListResponse
	1, // 3: FieldService.AddressFieldList:output_type -> FieldListResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_field_proto_init() }
func file_proto_service_field_proto_init() {
	if File_proto_service_field_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_service_field_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_field_proto_goTypes,
		DependencyIndexes: file_proto_service_field_proto_depIdxs,
	}.Build()
	File_proto_service_field_proto = out.File
	file_proto_service_field_proto_rawDesc = nil
	file_proto_service_field_proto_goTypes = nil
	file_proto_service_field_proto_depIdxs = nil
}
