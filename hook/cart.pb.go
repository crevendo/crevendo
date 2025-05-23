// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/hook/cart.proto

package hook

import (
	data "github.com/crevendo/crevendo/data"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FeeListHookParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cart          *data.Cart             `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	User          *data.User             `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Fees          []*data.Fee            `protobuf:"bytes,2,rep,name=fees,proto3" json:"fees,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FeeListHookParams) Reset() {
	*x = FeeListHookParams{}
	mi := &file_proto_hook_cart_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeeListHookParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeeListHookParams) ProtoMessage() {}

func (x *FeeListHookParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hook_cart_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeeListHookParams.ProtoReflect.Descriptor instead.
func (*FeeListHookParams) Descriptor() ([]byte, []int) {
	return file_proto_hook_cart_proto_rawDescGZIP(), []int{0}
}

func (x *FeeListHookParams) GetCart() *data.Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

func (x *FeeListHookParams) GetUser() *data.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *FeeListHookParams) GetFees() []*data.Fee {
	if x != nil {
		return x.Fees
	}
	return nil
}

var File_proto_hook_cart_proto protoreflect.FileDescriptor

const file_proto_hook_cart_proto_rawDesc = "" +
	"\n" +
	"\x15proto/hook/cart.proto\x1a\x0fproto/fee.proto\x1a\x10proto/user.proto\x1a\x10proto/cart.proto\"c\n" +
	"\x11FeeListHookParams\x12\x19\n" +
	"\x04cart\x18\x01 \x01(\v2\x05.CartR\x04cart\x12\x19\n" +
	"\x04user\x18\x03 \x01(\v2\x05.UserR\x04user\x12\x18\n" +
	"\x04fees\x18\x02 \x03(\v2\x04.FeeR\x04fees2@\n" +
	"\tCartHooks\x123\n" +
	"\aFeeList\x12\x12.FeeListHookParams\x1a\x12.FeeListHookParams\"\x00B#Z!github.com/crevendo/crevendo/hookb\x06proto3"

var (
	file_proto_hook_cart_proto_rawDescOnce sync.Once
	file_proto_hook_cart_proto_rawDescData []byte
)

func file_proto_hook_cart_proto_rawDescGZIP() []byte {
	file_proto_hook_cart_proto_rawDescOnce.Do(func() {
		file_proto_hook_cart_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_hook_cart_proto_rawDesc), len(file_proto_hook_cart_proto_rawDesc)))
	})
	return file_proto_hook_cart_proto_rawDescData
}

var file_proto_hook_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_hook_cart_proto_goTypes = []any{
	(*FeeListHookParams)(nil), // 0: FeeListHookParams
	(*data.Cart)(nil),         // 1: Cart
	(*data.User)(nil),         // 2: User
	(*data.Fee)(nil),          // 3: Fee
}
var file_proto_hook_cart_proto_depIdxs = []int32{
	1, // 0: FeeListHookParams.cart:type_name -> Cart
	2, // 1: FeeListHookParams.user:type_name -> User
	3, // 2: FeeListHookParams.fees:type_name -> Fee
	0, // 3: CartHooks.FeeList:input_type -> FeeListHookParams
	0, // 4: CartHooks.FeeList:output_type -> FeeListHookParams
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_hook_cart_proto_init() }
func file_proto_hook_cart_proto_init() {
	if File_proto_hook_cart_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_hook_cart_proto_rawDesc), len(file_proto_hook_cart_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hook_cart_proto_goTypes,
		DependencyIndexes: file_proto_hook_cart_proto_depIdxs,
		MessageInfos:      file_proto_hook_cart_proto_msgTypes,
	}.Build()
	File_proto_hook_cart_proto = out.File
	file_proto_hook_cart_proto_goTypes = nil
	file_proto_hook_cart_proto_depIdxs = nil
}
