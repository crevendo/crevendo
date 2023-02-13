// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/payment.proto

package data

import (
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

type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreId  uint32  `protobuf:"varint,2,opt,name=storeId,proto3" json:"storeId,omitempty"`
	CartId   uint32  `protobuf:"varint,3,opt,name=cartId,proto3" json:"cartId,omitempty"`
	Cart     *Cart   `protobuf:"bytes,4,opt,name=cart,proto3" json:"cart,omitempty"`
	UserId   uint32  `protobuf:"varint,5,opt,name=userId,proto3" json:"userId,omitempty"`
	User     *User   `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	Gateway  string  `protobuf:"bytes,7,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Total    float64 `protobuf:"fixed64,8,opt,name=total,proto3" json:"total,omitempty"`
	Status   uint32  `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	MethodId string  `protobuf:"bytes,10,opt,name=methodId,proto3" json:"methodId,omitempty"`
	OrderId  uint32  `protobuf:"varint,11,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Data     []*Data `protobuf:"bytes,12,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{0}
}

func (x *Payment) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Payment) GetStoreId() uint32 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *Payment) GetCartId() uint32 {
	if x != nil {
		return x.CartId
	}
	return 0
}

func (x *Payment) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

func (x *Payment) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Payment) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Payment) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

func (x *Payment) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Payment) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Payment) GetMethodId() string {
	if x != nil {
		return x.MethodId
	}
	return ""
}

func (x *Payment) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Payment) GetData() []*Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type PaymentMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginId string `protobuf:"bytes,1,opt,name=originId,proto3" json:"originId,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Issuer   string `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Image    string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Type     string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	NeedCvv  bool   `protobuf:"varint,6,opt,name=need_cvv,json=needCvv,proto3" json:"need_cvv,omitempty"`
}

func (x *PaymentMethod) Reset() {
	*x = PaymentMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethod) ProtoMessage() {}

func (x *PaymentMethod) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethod.ProtoReflect.Descriptor instead.
func (*PaymentMethod) Descriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentMethod) GetOriginId() string {
	if x != nil {
		return x.OriginId
	}
	return ""
}

func (x *PaymentMethod) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PaymentMethod) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *PaymentMethod) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *PaymentMethod) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PaymentMethod) GetNeedCvv() bool {
	if x != nil {
		return x.NeedCvv
	}
	return false
}

var File_proto_payment_proto protoreflect.FileDescriptor

var file_proto_payment_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x02, 0x0a, 0x07,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x04, 0x63, 0x61, 0x72,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04,
	0x63, 0x61, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x9c, 0x01, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x65, 0x64, 0x5f, 0x63, 0x76, 0x76,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6e, 0x65, 0x65, 0x64, 0x43, 0x76, 0x76, 0x42,
	0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72,
	0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73,
	0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x72, 0x65,
	0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_payment_proto_rawDescOnce sync.Once
	file_proto_payment_proto_rawDescData = file_proto_payment_proto_rawDesc
)

func file_proto_payment_proto_rawDescGZIP() []byte {
	file_proto_payment_proto_rawDescOnce.Do(func() {
		file_proto_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_payment_proto_rawDescData)
	})
	return file_proto_payment_proto_rawDescData
}

var file_proto_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_payment_proto_goTypes = []interface{}{
	(*Payment)(nil),       // 0: Payment
	(*PaymentMethod)(nil), // 1: PaymentMethod
	(*Cart)(nil),          // 2: Cart
	(*User)(nil),          // 3: User
	(*Data)(nil),          // 4: Data
}
var file_proto_payment_proto_depIdxs = []int32{
	2, // 0: Payment.cart:type_name -> Cart
	3, // 1: Payment.user:type_name -> User
	4, // 2: Payment.data:type_name -> Data
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_payment_proto_init() }
func file_proto_payment_proto_init() {
	if File_proto_payment_proto != nil {
		return
	}
	file_proto_cart_proto_init()
	file_proto_user_proto_init()
	file_proto_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payment); i {
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
		file_proto_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentMethod); i {
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
			RawDescriptor: file_proto_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_payment_proto_goTypes,
		DependencyIndexes: file_proto_payment_proto_depIdxs,
		MessageInfos:      file_proto_payment_proto_msgTypes,
	}.Build()
	File_proto_payment_proto = out.File
	file_proto_payment_proto_rawDesc = nil
	file_proto_payment_proto_goTypes = nil
	file_proto_payment_proto_depIdxs = nil
}
