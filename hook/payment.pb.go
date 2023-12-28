// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: proto/hook/payment.proto

package hook

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

type PaymentProcessHookParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payment *data.Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
	Fees    []*data.Fee   `protobuf:"bytes,2,rep,name=fees,proto3" json:"fees,omitempty"`
	Cvv     string        `protobuf:"bytes,3,opt,name=cvv,proto3" json:"cvv,omitempty"`
}

func (x *PaymentProcessHookParams) Reset() {
	*x = PaymentProcessHookParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hook_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentProcessHookParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentProcessHookParams) ProtoMessage() {}

func (x *PaymentProcessHookParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hook_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentProcessHookParams.ProtoReflect.Descriptor instead.
func (*PaymentProcessHookParams) Descriptor() ([]byte, []int) {
	return file_proto_hook_payment_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentProcessHookParams) GetPayment() *data.Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

func (x *PaymentProcessHookParams) GetFees() []*data.Fee {
	if x != nil {
		return x.Fees
	}
	return nil
}

func (x *PaymentProcessHookParams) GetCvv() string {
	if x != nil {
		return x.Cvv
	}
	return ""
}

type PaymentMethodListHookParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUUID string                `protobuf:"bytes,1,opt,name=userUUID,proto3" json:"userUUID,omitempty"`
	Methods  []*data.PaymentMethod `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (x *PaymentMethodListHookParams) Reset() {
	*x = PaymentMethodListHookParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hook_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentMethodListHookParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethodListHookParams) ProtoMessage() {}

func (x *PaymentMethodListHookParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hook_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethodListHookParams.ProtoReflect.Descriptor instead.
func (*PaymentMethodListHookParams) Descriptor() ([]byte, []int) {
	return file_proto_hook_payment_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentMethodListHookParams) GetUserUUID() string {
	if x != nil {
		return x.UserUUID
	}
	return ""
}

func (x *PaymentMethodListHookParams) GetMethods() []*data.PaymentMethod {
	if x != nil {
		return x.Methods
	}
	return nil
}

type PaymentMethodDeleteHookParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUUID string `protobuf:"bytes,1,opt,name=userUUID,proto3" json:"userUUID,omitempty"`
	MethodId string `protobuf:"bytes,2,opt,name=methodId,proto3" json:"methodId,omitempty"`
}

func (x *PaymentMethodDeleteHookParams) Reset() {
	*x = PaymentMethodDeleteHookParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hook_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentMethodDeleteHookParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethodDeleteHookParams) ProtoMessage() {}

func (x *PaymentMethodDeleteHookParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hook_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethodDeleteHookParams.ProtoReflect.Descriptor instead.
func (*PaymentMethodDeleteHookParams) Descriptor() ([]byte, []int) {
	return file_proto_hook_payment_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentMethodDeleteHookParams) GetUserUUID() string {
	if x != nil {
		return x.UserUUID
	}
	return ""
}

func (x *PaymentMethodDeleteHookParams) GetMethodId() string {
	if x != nil {
		return x.MethodId
	}
	return ""
}

var File_proto_hook_payment_proto protoreflect.FileDescriptor

var file_proto_hook_payment_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x6f, 0x6f, 0x6b, 0x2f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6a, 0x0a, 0x18, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x48, 0x6f, 0x6f, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x22, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x04,
	0x2e, 0x46, 0x65, 0x65, 0x52, 0x04, 0x66, 0x65, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x76,
	0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x76, 0x76, 0x22, 0x63, 0x0a, 0x1b,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x48, 0x6f, 0x6f, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x22, 0x57, 0x0a, 0x1d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x32, 0xf6, 0x01, 0x0a, 0x0c, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x4a, 0x0a, 0x0a, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x6f,
	0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1c, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x6f, 0x6b, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6f,
	0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1e, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6f,
	0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0e, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x6f, 0x6b,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x19, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x6f, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x63, 0x72, 0x65, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x2f, 0x68, 0x6f, 0x6f, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hook_payment_proto_rawDescOnce sync.Once
	file_proto_hook_payment_proto_rawDescData = file_proto_hook_payment_proto_rawDesc
)

func file_proto_hook_payment_proto_rawDescGZIP() []byte {
	file_proto_hook_payment_proto_rawDescOnce.Do(func() {
		file_proto_hook_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hook_payment_proto_rawDescData)
	})
	return file_proto_hook_payment_proto_rawDescData
}

var file_proto_hook_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_hook_payment_proto_goTypes = []interface{}{
	(*PaymentProcessHookParams)(nil),      // 0: PaymentProcessHookParams
	(*PaymentMethodListHookParams)(nil),   // 1: PaymentMethodListHookParams
	(*PaymentMethodDeleteHookParams)(nil), // 2: PaymentMethodDeleteHookParams
	(*data.Payment)(nil),                  // 3: Payment
	(*data.Fee)(nil),                      // 4: Fee
	(*data.PaymentMethod)(nil),            // 5: PaymentMethod
}
var file_proto_hook_payment_proto_depIdxs = []int32{
	3, // 0: PaymentProcessHookParams.payment:type_name -> Payment
	4, // 1: PaymentProcessHookParams.fees:type_name -> Fee
	5, // 2: PaymentMethodListHookParams.methods:type_name -> PaymentMethod
	1, // 3: PaymentHooks.MethodList:input_type -> PaymentMethodListHookParams
	2, // 4: PaymentHooks.MethodDelete:input_type -> PaymentMethodDeleteHookParams
	0, // 5: PaymentHooks.ProcessPayment:input_type -> PaymentProcessHookParams
	1, // 6: PaymentHooks.MethodList:output_type -> PaymentMethodListHookParams
	2, // 7: PaymentHooks.MethodDelete:output_type -> PaymentMethodDeleteHookParams
	0, // 8: PaymentHooks.ProcessPayment:output_type -> PaymentProcessHookParams
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_hook_payment_proto_init() }
func file_proto_hook_payment_proto_init() {
	if File_proto_hook_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hook_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentProcessHookParams); i {
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
		file_proto_hook_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentMethodListHookParams); i {
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
		file_proto_hook_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentMethodDeleteHookParams); i {
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
			RawDescriptor: file_proto_hook_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hook_payment_proto_goTypes,
		DependencyIndexes: file_proto_hook_payment_proto_depIdxs,
		MessageInfos:      file_proto_hook_payment_proto_msgTypes,
	}.Build()
	File_proto_hook_payment_proto = out.File
	file_proto_hook_payment_proto_rawDesc = nil
	file_proto_hook_payment_proto_goTypes = nil
	file_proto_hook_payment_proto_depIdxs = nil
}
