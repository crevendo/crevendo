// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/message/payment.proto

package message

import (
	data "github.com/crevendo/crevendo/data"
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

type PaymentMethodDeleteMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentMethodId string `protobuf:"bytes,2,opt,name=paymentMethodId,proto3" json:"paymentMethodId,omitempty"`
}

func (x *PaymentMethodDeleteMessage) Reset() {
	*x = PaymentMethodDeleteMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentMethodDeleteMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethodDeleteMessage) ProtoMessage() {}

func (x *PaymentMethodDeleteMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethodDeleteMessage.ProtoReflect.Descriptor instead.
func (*PaymentMethodDeleteMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_payment_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentMethodDeleteMessage) GetPaymentMethodId() string {
	if x != nil {
		return x.PaymentMethodId
	}
	return ""
}

type PaymentMethodDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PaymentMethodDeleteResponse) Reset() {
	*x = PaymentMethodDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentMethodDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethodDeleteResponse) ProtoMessage() {}

func (x *PaymentMethodDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethodDeleteResponse.ProtoReflect.Descriptor instead.
func (*PaymentMethodDeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_payment_proto_rawDescGZIP(), []int{1}
}

type PaymentMethodListMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *PaymentMethodListMessage) Reset() {
	*x = PaymentMethodListMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentMethodListMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethodListMessage) ProtoMessage() {}

func (x *PaymentMethodListMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethodListMessage.ProtoReflect.Descriptor instead.
func (*PaymentMethodListMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_payment_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentMethodListMessage) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type PaymentMethodListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentMethods []*data.PaymentMethod `protobuf:"bytes,1,rep,name=PaymentMethods,proto3" json:"PaymentMethods,omitempty"`
}

func (x *PaymentMethodListResponse) Reset() {
	*x = PaymentMethodListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentMethodListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethodListResponse) ProtoMessage() {}

func (x *PaymentMethodListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethodListResponse.ProtoReflect.Descriptor instead.
func (*PaymentMethodListResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_payment_proto_rawDescGZIP(), []int{3}
}

func (x *PaymentMethodListResponse) GetPaymentMethods() []*data.PaymentMethod {
	if x != nil {
		return x.PaymentMethods
	}
	return nil
}

var File_proto_message_payment_proto protoreflect.FileDescriptor

var file_proto_message_payment_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x46, 0x0a, 0x1a, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x28, 0x0a, 0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x22, 0x1d, 0x0a, 0x1b, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x0a, 0x18, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x53, 0x0a,
	0x19, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x52, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_message_payment_proto_rawDescOnce sync.Once
	file_proto_message_payment_proto_rawDescData = file_proto_message_payment_proto_rawDesc
)

func file_proto_message_payment_proto_rawDescGZIP() []byte {
	file_proto_message_payment_proto_rawDescOnce.Do(func() {
		file_proto_message_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_payment_proto_rawDescData)
	})
	return file_proto_message_payment_proto_rawDescData
}

var file_proto_message_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_message_payment_proto_goTypes = []interface{}{
	(*PaymentMethodDeleteMessage)(nil),  // 0: PaymentMethodDeleteMessage
	(*PaymentMethodDeleteResponse)(nil), // 1: PaymentMethodDeleteResponse
	(*PaymentMethodListMessage)(nil),    // 2: PaymentMethodListMessage
	(*PaymentMethodListResponse)(nil),   // 3: PaymentMethodListResponse
	(*data.PaymentMethod)(nil),          // 4: PaymentMethod
}
var file_proto_message_payment_proto_depIdxs = []int32{
	4, // 0: PaymentMethodListResponse.PaymentMethods:type_name -> PaymentMethod
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_message_payment_proto_init() }
func file_proto_message_payment_proto_init() {
	if File_proto_message_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_message_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentMethodDeleteMessage); i {
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
		file_proto_message_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentMethodDeleteResponse); i {
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
		file_proto_message_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentMethodListMessage); i {
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
		file_proto_message_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentMethodListResponse); i {
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
			RawDescriptor: file_proto_message_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_message_payment_proto_goTypes,
		DependencyIndexes: file_proto_message_payment_proto_depIdxs,
		MessageInfos:      file_proto_message_payment_proto_msgTypes,
	}.Build()
	File_proto_message_payment_proto = out.File
	file_proto_message_payment_proto_rawDesc = nil
	file_proto_message_payment_proto_goTypes = nil
	file_proto_message_payment_proto_depIdxs = nil
}
