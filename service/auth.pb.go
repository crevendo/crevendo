// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/service/auth.proto

package service

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

var File_proto_service_auth_proto protoreflect.FileDescriptor

var file_proto_service_auth_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc8, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0e, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a,
	0x51, 0x75, 0x69, 0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x51, 0x75, 0x69,
	0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x13,
	0x2e, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x0e, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0f, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x11, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72,
	0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_service_auth_proto_goTypes = []interface{}{
	(*message.LoginMessage)(nil),       // 0: LoginMessage
	(*message.QuickLoginMessage)(nil),  // 1: QuickLoginMessage
	(*message.VerifyMessage)(nil),      // 2: VerifyMessage
	(*message.RegisterMessage)(nil),    // 3: RegisterMessage
	(*message.LoginResponse)(nil),      // 4: LoginResponse
	(*message.QuickLoginResponse)(nil), // 5: QuickLoginResponse
	(*message.VerifyResponse)(nil),     // 6: VerifyResponse
	(*message.RegisterResponse)(nil),   // 7: RegisterResponse
}
var file_proto_service_auth_proto_depIdxs = []int32{
	0, // 0: AuthService.Login:input_type -> LoginMessage
	1, // 1: AuthService.QuickLogin:input_type -> QuickLoginMessage
	2, // 2: AuthService.Verify:input_type -> VerifyMessage
	3, // 3: AuthService.Register:input_type -> RegisterMessage
	4, // 4: AuthService.Login:output_type -> LoginResponse
	5, // 5: AuthService.QuickLogin:output_type -> QuickLoginResponse
	6, // 6: AuthService.Verify:output_type -> VerifyResponse
	7, // 7: AuthService.Register:output_type -> RegisterResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_auth_proto_init() }
func file_proto_service_auth_proto_init() {
	if File_proto_service_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_service_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_auth_proto_goTypes,
		DependencyIndexes: file_proto_service_auth_proto_depIdxs,
	}.Build()
	File_proto_service_auth_proto = out.File
	file_proto_service_auth_proto_rawDesc = nil
	file_proto_service_auth_proto_goTypes = nil
	file_proto_service_auth_proto_depIdxs = nil
}