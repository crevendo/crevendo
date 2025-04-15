// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/service/order.proto

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

var File_proto_service_order_proto protoreflect.FileDescriptor

const file_proto_service_order_proto_rawDesc = "" +
	"\n" +
	"\x19proto/service/order.proto\x1a\x19proto/message/order.proto2\xf2\x02\n" +
	"\fOrderService\x12,\n" +
	"\x03Get\x12\x10.OrderGetMessage\x1a\x11.OrderGetResponse\"\x00\x125\n" +
	"\x06Create\x12\x13.OrderCreateMessage\x1a\x14.OrderCreateResponse\"\x00\x12/\n" +
	"\x04List\x12\x11.OrderListMessage\x1a\x12.OrderListResponse\"\x00\x125\n" +
	"\x06Update\x12\x13.OrderUpdateMessage\x1a\x14.OrderUpdateResponse\"\x00\x12A\n" +
	"\n" +
	"StatusList\x12\x17.OrderStatusListMessage\x1a\x18.OrderStatusListResponse\"\x00\x12R\n" +
	"\x13OrderItemStatusList\x12\x1b.OrderItemStatusListMessage\x1a\x1c.OrderItemStatusListResponse\"\x00B#Z!github.com/crevendo/crevendo/grpcb\x06proto3"

var file_proto_service_order_proto_goTypes = []any{
	(*message.OrderGetMessage)(nil),             // 0: OrderGetMessage
	(*message.OrderCreateMessage)(nil),          // 1: OrderCreateMessage
	(*message.OrderListMessage)(nil),            // 2: OrderListMessage
	(*message.OrderUpdateMessage)(nil),          // 3: OrderUpdateMessage
	(*message.OrderStatusListMessage)(nil),      // 4: OrderStatusListMessage
	(*message.OrderItemStatusListMessage)(nil),  // 5: OrderItemStatusListMessage
	(*message.OrderGetResponse)(nil),            // 6: OrderGetResponse
	(*message.OrderCreateResponse)(nil),         // 7: OrderCreateResponse
	(*message.OrderListResponse)(nil),           // 8: OrderListResponse
	(*message.OrderUpdateResponse)(nil),         // 9: OrderUpdateResponse
	(*message.OrderStatusListResponse)(nil),     // 10: OrderStatusListResponse
	(*message.OrderItemStatusListResponse)(nil), // 11: OrderItemStatusListResponse
}
var file_proto_service_order_proto_depIdxs = []int32{
	0,  // 0: OrderService.Get:input_type -> OrderGetMessage
	1,  // 1: OrderService.Create:input_type -> OrderCreateMessage
	2,  // 2: OrderService.List:input_type -> OrderListMessage
	3,  // 3: OrderService.Update:input_type -> OrderUpdateMessage
	4,  // 4: OrderService.StatusList:input_type -> OrderStatusListMessage
	5,  // 5: OrderService.OrderItemStatusList:input_type -> OrderItemStatusListMessage
	6,  // 6: OrderService.Get:output_type -> OrderGetResponse
	7,  // 7: OrderService.Create:output_type -> OrderCreateResponse
	8,  // 8: OrderService.List:output_type -> OrderListResponse
	9,  // 9: OrderService.Update:output_type -> OrderUpdateResponse
	10, // 10: OrderService.StatusList:output_type -> OrderStatusListResponse
	11, // 11: OrderService.OrderItemStatusList:output_type -> OrderItemStatusListResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_order_proto_init() }
func file_proto_service_order_proto_init() {
	if File_proto_service_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_service_order_proto_rawDesc), len(file_proto_service_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_order_proto_goTypes,
		DependencyIndexes: file_proto_service_order_proto_depIdxs,
	}.Build()
	File_proto_service_order_proto = out.File
	file_proto_service_order_proto_goTypes = nil
	file_proto_service_order_proto_depIdxs = nil
}
