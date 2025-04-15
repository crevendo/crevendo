// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/message/cart.proto

package message

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

type CartGetMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CartGetMessage) Reset() {
	*x = CartGetMessage{}
	mi := &file_proto_message_cart_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CartGetMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartGetMessage) ProtoMessage() {}

func (x *CartGetMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartGetMessage.ProtoReflect.Descriptor instead.
func (*CartGetMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_cart_proto_rawDescGZIP(), []int{0}
}

func (x *CartGetMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CartGetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cart          *data.Cart             `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CartGetResponse) Reset() {
	*x = CartGetResponse{}
	mi := &file_proto_message_cart_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CartGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartGetResponse) ProtoMessage() {}

func (x *CartGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartGetResponse.ProtoReflect.Descriptor instead.
func (*CartGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_cart_proto_rawDescGZIP(), []int{1}
}

func (x *CartGetResponse) GetCart() *data.Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

type CartCreateMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CartCreateMessage) Reset() {
	*x = CartCreateMessage{}
	mi := &file_proto_message_cart_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CartCreateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartCreateMessage) ProtoMessage() {}

func (x *CartCreateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartCreateMessage.ProtoReflect.Descriptor instead.
func (*CartCreateMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_cart_proto_rawDescGZIP(), []int{2}
}

type CartCreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cart          *data.Cart             `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CartCreateResponse) Reset() {
	*x = CartCreateResponse{}
	mi := &file_proto_message_cart_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CartCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartCreateResponse) ProtoMessage() {}

func (x *CartCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartCreateResponse.ProtoReflect.Descriptor instead.
func (*CartCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_cart_proto_rawDescGZIP(), []int{3}
}

func (x *CartCreateResponse) GetCart() *data.Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

type AddItemMessage struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	CartId          int32                  `protobuf:"varint,1,opt,name=cartId,proto3" json:"cartId,omitempty"`
	Quantity        int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	ProductId       *int32                 `protobuf:"varint,3,opt,name=productId,proto3,oneof" json:"productId,omitempty"`
	ProductCustomId *string                `protobuf:"bytes,4,opt,name=productCustomId,proto3,oneof" json:"productCustomId,omitempty"`
	CustomData      []*data.Data           `protobuf:"bytes,5,rep,name=customData,proto3" json:"customData,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *AddItemMessage) Reset() {
	*x = AddItemMessage{}
	mi := &file_proto_message_cart_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddItemMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemMessage) ProtoMessage() {}

func (x *AddItemMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemMessage.ProtoReflect.Descriptor instead.
func (*AddItemMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_cart_proto_rawDescGZIP(), []int{4}
}

func (x *AddItemMessage) GetCartId() int32 {
	if x != nil {
		return x.CartId
	}
	return 0
}

func (x *AddItemMessage) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *AddItemMessage) GetProductId() int32 {
	if x != nil && x.ProductId != nil {
		return *x.ProductId
	}
	return 0
}

func (x *AddItemMessage) GetProductCustomId() string {
	if x != nil && x.ProductCustomId != nil {
		return *x.ProductCustomId
	}
	return ""
}

func (x *AddItemMessage) GetCustomData() []*data.Data {
	if x != nil {
		return x.CustomData
	}
	return nil
}

type AddItemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cart          *data.Cart             `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddItemResponse) Reset() {
	*x = AddItemResponse{}
	mi := &file_proto_message_cart_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemResponse) ProtoMessage() {}

func (x *AddItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemResponse.ProtoReflect.Descriptor instead.
func (*AddItemResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_cart_proto_rawDescGZIP(), []int{5}
}

func (x *AddItemResponse) GetCart() *data.Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

type RemoveItemMessage struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	CartId          int32                  `protobuf:"varint,1,opt,name=cartId,proto3" json:"cartId,omitempty"`
	ProductId       *int32                 `protobuf:"varint,2,opt,name=productId,proto3,oneof" json:"productId,omitempty"`
	ProductCustomId *string                `protobuf:"bytes,3,opt,name=productCustomId,proto3,oneof" json:"productCustomId,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *RemoveItemMessage) Reset() {
	*x = RemoveItemMessage{}
	mi := &file_proto_message_cart_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveItemMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveItemMessage) ProtoMessage() {}

func (x *RemoveItemMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveItemMessage.ProtoReflect.Descriptor instead.
func (*RemoveItemMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_cart_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveItemMessage) GetCartId() int32 {
	if x != nil {
		return x.CartId
	}
	return 0
}

func (x *RemoveItemMessage) GetProductId() int32 {
	if x != nil && x.ProductId != nil {
		return *x.ProductId
	}
	return 0
}

func (x *RemoveItemMessage) GetProductCustomId() string {
	if x != nil && x.ProductCustomId != nil {
		return *x.ProductCustomId
	}
	return ""
}

type RemoveItemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cart          *data.Cart             `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveItemResponse) Reset() {
	*x = RemoveItemResponse{}
	mi := &file_proto_message_cart_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveItemResponse) ProtoMessage() {}

func (x *RemoveItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveItemResponse.ProtoReflect.Descriptor instead.
func (*RemoveItemResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_cart_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveItemResponse) GetCart() *data.Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

type UpdateItemQuantityMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CartId        int32                  `protobuf:"varint,1,opt,name=cartId,proto3" json:"cartId,omitempty"`
	ItemId        int32                  `protobuf:"varint,2,opt,name=itemId,proto3" json:"itemId,omitempty"`
	Quantity      int32                  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateItemQuantityMessage) Reset() {
	*x = UpdateItemQuantityMessage{}
	mi := &file_proto_message_cart_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateItemQuantityMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemQuantityMessage) ProtoMessage() {}

func (x *UpdateItemQuantityMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemQuantityMessage.ProtoReflect.Descriptor instead.
func (*UpdateItemQuantityMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_cart_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateItemQuantityMessage) GetCartId() int32 {
	if x != nil {
		return x.CartId
	}
	return 0
}

func (x *UpdateItemQuantityMessage) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *UpdateItemQuantityMessage) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type UpdateItemQuantityResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cart          *data.Cart             `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateItemQuantityResponse) Reset() {
	*x = UpdateItemQuantityResponse{}
	mi := &file_proto_message_cart_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateItemQuantityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemQuantityResponse) ProtoMessage() {}

func (x *UpdateItemQuantityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemQuantityResponse.ProtoReflect.Descriptor instead.
func (*UpdateItemQuantityResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_cart_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateItemQuantityResponse) GetCart() *data.Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

type GetOrderTotalResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         float32                `protobuf:"fixed32,1,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderTotalResponse) Reset() {
	*x = GetOrderTotalResponse{}
	mi := &file_proto_message_cart_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderTotalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderTotalResponse) ProtoMessage() {}

func (x *GetOrderTotalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderTotalResponse.ProtoReflect.Descriptor instead.
func (*GetOrderTotalResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_cart_proto_rawDescGZIP(), []int{10}
}

func (x *GetOrderTotalResponse) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_proto_message_cart_proto protoreflect.FileDescriptor

const file_proto_message_cart_proto_rawDesc = "" +
	"\n" +
	"\x18proto/message/cart.proto\x1a\x10proto/cart.proto\x1a\x10proto/data.proto\" \n" +
	"\x0eCartGetMessage\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\",\n" +
	"\x0fCartGetResponse\x12\x19\n" +
	"\x04cart\x18\x01 \x01(\v2\x05.CartR\x04cart\"\x13\n" +
	"\x11CartCreateMessage\"/\n" +
	"\x12CartCreateResponse\x12\x19\n" +
	"\x04cart\x18\x01 \x01(\v2\x05.CartR\x04cart\"\xdf\x01\n" +
	"\x0eAddItemMessage\x12\x16\n" +
	"\x06cartId\x18\x01 \x01(\x05R\x06cartId\x12\x1a\n" +
	"\bquantity\x18\x02 \x01(\x05R\bquantity\x12!\n" +
	"\tproductId\x18\x03 \x01(\x05H\x00R\tproductId\x88\x01\x01\x12-\n" +
	"\x0fproductCustomId\x18\x04 \x01(\tH\x01R\x0fproductCustomId\x88\x01\x01\x12%\n" +
	"\n" +
	"customData\x18\x05 \x03(\v2\x05.DataR\n" +
	"customDataB\f\n" +
	"\n" +
	"_productIdB\x12\n" +
	"\x10_productCustomId\",\n" +
	"\x0fAddItemResponse\x12\x19\n" +
	"\x04cart\x18\x01 \x01(\v2\x05.CartR\x04cart\"\x9f\x01\n" +
	"\x11RemoveItemMessage\x12\x16\n" +
	"\x06cartId\x18\x01 \x01(\x05R\x06cartId\x12!\n" +
	"\tproductId\x18\x02 \x01(\x05H\x00R\tproductId\x88\x01\x01\x12-\n" +
	"\x0fproductCustomId\x18\x03 \x01(\tH\x01R\x0fproductCustomId\x88\x01\x01B\f\n" +
	"\n" +
	"_productIdB\x12\n" +
	"\x10_productCustomId\"/\n" +
	"\x12RemoveItemResponse\x12\x19\n" +
	"\x04cart\x18\x01 \x01(\v2\x05.CartR\x04cart\"g\n" +
	"\x19UpdateItemQuantityMessage\x12\x16\n" +
	"\x06cartId\x18\x01 \x01(\x05R\x06cartId\x12\x16\n" +
	"\x06itemId\x18\x02 \x01(\x05R\x06itemId\x12\x1a\n" +
	"\bquantity\x18\x03 \x01(\x05R\bquantity\"7\n" +
	"\x1aUpdateItemQuantityResponse\x12\x19\n" +
	"\x04cart\x18\x01 \x01(\v2\x05.CartR\x04cart\"-\n" +
	"\x15GetOrderTotalResponse\x12\x14\n" +
	"\x05total\x18\x01 \x01(\x02R\x05totalB&Z$github.com/crevendo/crevendo/messageb\x06proto3"

var (
	file_proto_message_cart_proto_rawDescOnce sync.Once
	file_proto_message_cart_proto_rawDescData []byte
)

func file_proto_message_cart_proto_rawDescGZIP() []byte {
	file_proto_message_cart_proto_rawDescOnce.Do(func() {
		file_proto_message_cart_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_message_cart_proto_rawDesc), len(file_proto_message_cart_proto_rawDesc)))
	})
	return file_proto_message_cart_proto_rawDescData
}

var file_proto_message_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_message_cart_proto_goTypes = []any{
	(*CartGetMessage)(nil),             // 0: CartGetMessage
	(*CartGetResponse)(nil),            // 1: CartGetResponse
	(*CartCreateMessage)(nil),          // 2: CartCreateMessage
	(*CartCreateResponse)(nil),         // 3: CartCreateResponse
	(*AddItemMessage)(nil),             // 4: AddItemMessage
	(*AddItemResponse)(nil),            // 5: AddItemResponse
	(*RemoveItemMessage)(nil),          // 6: RemoveItemMessage
	(*RemoveItemResponse)(nil),         // 7: RemoveItemResponse
	(*UpdateItemQuantityMessage)(nil),  // 8: UpdateItemQuantityMessage
	(*UpdateItemQuantityResponse)(nil), // 9: UpdateItemQuantityResponse
	(*GetOrderTotalResponse)(nil),      // 10: GetOrderTotalResponse
	(*data.Cart)(nil),                  // 11: Cart
	(*data.Data)(nil),                  // 12: Data
}
var file_proto_message_cart_proto_depIdxs = []int32{
	11, // 0: CartGetResponse.cart:type_name -> Cart
	11, // 1: CartCreateResponse.cart:type_name -> Cart
	12, // 2: AddItemMessage.customData:type_name -> Data
	11, // 3: AddItemResponse.cart:type_name -> Cart
	11, // 4: RemoveItemResponse.cart:type_name -> Cart
	11, // 5: UpdateItemQuantityResponse.cart:type_name -> Cart
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_message_cart_proto_init() }
func file_proto_message_cart_proto_init() {
	if File_proto_message_cart_proto != nil {
		return
	}
	file_proto_message_cart_proto_msgTypes[4].OneofWrappers = []any{}
	file_proto_message_cart_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_message_cart_proto_rawDesc), len(file_proto_message_cart_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_message_cart_proto_goTypes,
		DependencyIndexes: file_proto_message_cart_proto_depIdxs,
		MessageInfos:      file_proto_message_cart_proto_msgTypes,
	}.Build()
	File_proto_message_cart_proto = out.File
	file_proto_message_cart_proto_goTypes = nil
	file_proto_message_cart_proto_depIdxs = nil
}
