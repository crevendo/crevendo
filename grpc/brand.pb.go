// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: proto/service/brand.proto

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

var File_proto_service_brand_proto protoreflect.FileDescriptor

var file_proto_service_brand_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa4, 0x01, 0x0a, 0x0c, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x11, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x12, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x10, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x11, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x13, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a,
	0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x65, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_service_brand_proto_goTypes = []interface{}{
	(*message.BrandListMessage)(nil),    // 0: BrandListMessage
	(*message.BrandGetMessage)(nil),     // 1: BrandGetMessage
	(*message.BrandUpdateMessage)(nil),  // 2: BrandUpdateMessage
	(*message.BrandListResponse)(nil),   // 3: BrandListResponse
	(*message.BrandGetResponse)(nil),    // 4: BrandGetResponse
	(*message.BrandUpdateResponse)(nil), // 5: BrandUpdateResponse
}
var file_proto_service_brand_proto_depIdxs = []int32{
	0, // 0: BrandService.List:input_type -> BrandListMessage
	1, // 1: BrandService.Get:input_type -> BrandGetMessage
	2, // 2: BrandService.Update:input_type -> BrandUpdateMessage
	3, // 3: BrandService.List:output_type -> BrandListResponse
	4, // 4: BrandService.Get:output_type -> BrandGetResponse
	5, // 5: BrandService.Update:output_type -> BrandUpdateResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_brand_proto_init() }
func file_proto_service_brand_proto_init() {
	if File_proto_service_brand_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_service_brand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_brand_proto_goTypes,
		DependencyIndexes: file_proto_service_brand_proto_depIdxs,
	}.Build()
	File_proto_service_brand_proto = out.File
	file_proto_service_brand_proto_rawDesc = nil
	file_proto_service_brand_proto_goTypes = nil
	file_proto_service_brand_proto_depIdxs = nil
}
