// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/service/integration.proto

package grpc

import (
	message "github.com/crevendo/crevendo/message"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_service_integration_proto protoreflect.FileDescriptor

const file_proto_service_integration_proto_rawDesc = "" +
	"\n" +
	"\x1fproto/service/integration.proto\x1a\x1fproto/message/integration.proto2\x94\x02\n" +
	"\x12IntegrationService\x12;\n" +
	"\x04List\x12\x17.IntegrationListMessage\x1a\x18.IntegrationListResponse\"\x00\x128\n" +
	"\x03Get\x12\x16.IntegrationGetMessage\x1a\x17.IntegrationGetResponse\"\x00\x12A\n" +
	"\x06Enable\x12\x19.IntegrationEnableMessage\x1a\x1a.IntegrationEnableResponse\"\x00\x12D\n" +
	"\aDisable\x12\x1a.IntegrationDisableMessage\x1a\x1b.IntegrationDisableResponse\"\x00B#Z!github.com/crevendo/crevendo/grpcb\x06proto3"

var file_proto_service_integration_proto_goTypes = []any{
	(*message.IntegrationListMessage)(nil),     // 0: IntegrationListMessage
	(*message.IntegrationGetMessage)(nil),      // 1: IntegrationGetMessage
	(*message.IntegrationEnableMessage)(nil),   // 2: IntegrationEnableMessage
	(*message.IntegrationDisableMessage)(nil),  // 3: IntegrationDisableMessage
	(*message.IntegrationListResponse)(nil),    // 4: IntegrationListResponse
	(*message.IntegrationGetResponse)(nil),     // 5: IntegrationGetResponse
	(*message.IntegrationEnableResponse)(nil),  // 6: IntegrationEnableResponse
	(*message.IntegrationDisableResponse)(nil), // 7: IntegrationDisableResponse
}
var file_proto_service_integration_proto_depIdxs = []int32{
	0, // 0: IntegrationService.List:input_type -> IntegrationListMessage
	1, // 1: IntegrationService.Get:input_type -> IntegrationGetMessage
	2, // 2: IntegrationService.Enable:input_type -> IntegrationEnableMessage
	3, // 3: IntegrationService.Disable:input_type -> IntegrationDisableMessage
	4, // 4: IntegrationService.List:output_type -> IntegrationListResponse
	5, // 5: IntegrationService.Get:output_type -> IntegrationGetResponse
	6, // 6: IntegrationService.Enable:output_type -> IntegrationEnableResponse
	7, // 7: IntegrationService.Disable:output_type -> IntegrationDisableResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_integration_proto_init() }
func file_proto_service_integration_proto_init() {
	if File_proto_service_integration_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_service_integration_proto_rawDesc), len(file_proto_service_integration_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_integration_proto_goTypes,
		DependencyIndexes: file_proto_service_integration_proto_depIdxs,
	}.Build()
	File_proto_service_integration_proto = out.File
	file_proto_service_integration_proto_goTypes = nil
	file_proto_service_integration_proto_depIdxs = nil
}
