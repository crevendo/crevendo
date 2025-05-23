// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/service/address.proto

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

var File_proto_service_address_proto protoreflect.FileDescriptor

const file_proto_service_address_proto_rawDesc = "" +
	"\n" +
	"\x1bproto/service/address.proto\x1a\x1bproto/message/address.proto2\xa8\x02\n" +
	"\x0eAddressService\x120\n" +
	"\x03Get\x12\x12.AddressGetMessage\x1a\x13.AddressGetResponse\"\x00\x123\n" +
	"\x04List\x12\x13.AddressListMessage\x1a\x14.AddressListResponse\"\x00\x129\n" +
	"\x06Create\x12\x15.AddressCreateMessage\x1a\x16.AddressCreateResponse\"\x00\x129\n" +
	"\x06Delete\x12\x15.AddressDeleteMessage\x1a\x16.AddressDeleteResponse\"\x00\x129\n" +
	"\x06Update\x12\x15.AddressUpdateMessage\x1a\x16.AddressUpdateResponse\"\x00B#Z!github.com/crevendo/crevendo/grpcb\x06proto3"

var file_proto_service_address_proto_goTypes = []any{
	(*message.AddressGetMessage)(nil),     // 0: AddressGetMessage
	(*message.AddressListMessage)(nil),    // 1: AddressListMessage
	(*message.AddressCreateMessage)(nil),  // 2: AddressCreateMessage
	(*message.AddressDeleteMessage)(nil),  // 3: AddressDeleteMessage
	(*message.AddressUpdateMessage)(nil),  // 4: AddressUpdateMessage
	(*message.AddressGetResponse)(nil),    // 5: AddressGetResponse
	(*message.AddressListResponse)(nil),   // 6: AddressListResponse
	(*message.AddressCreateResponse)(nil), // 7: AddressCreateResponse
	(*message.AddressDeleteResponse)(nil), // 8: AddressDeleteResponse
	(*message.AddressUpdateResponse)(nil), // 9: AddressUpdateResponse
}
var file_proto_service_address_proto_depIdxs = []int32{
	0, // 0: AddressService.Get:input_type -> AddressGetMessage
	1, // 1: AddressService.List:input_type -> AddressListMessage
	2, // 2: AddressService.Create:input_type -> AddressCreateMessage
	3, // 3: AddressService.Delete:input_type -> AddressDeleteMessage
	4, // 4: AddressService.Update:input_type -> AddressUpdateMessage
	5, // 5: AddressService.Get:output_type -> AddressGetResponse
	6, // 6: AddressService.List:output_type -> AddressListResponse
	7, // 7: AddressService.Create:output_type -> AddressCreateResponse
	8, // 8: AddressService.Delete:output_type -> AddressDeleteResponse
	9, // 9: AddressService.Update:output_type -> AddressUpdateResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_address_proto_init() }
func file_proto_service_address_proto_init() {
	if File_proto_service_address_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_service_address_proto_rawDesc), len(file_proto_service_address_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_address_proto_goTypes,
		DependencyIndexes: file_proto_service_address_proto_depIdxs,
	}.Build()
	File_proto_service_address_proto = out.File
	file_proto_service_address_proto_goTypes = nil
	file_proto_service_address_proto_depIdxs = nil
}
