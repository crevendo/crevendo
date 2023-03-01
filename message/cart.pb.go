// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/message/cart.proto

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

type CartGetMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CartGetMessage) Reset() {
	*x = CartGetMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartGetMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartGetMessage) ProtoMessage() {}

func (x *CartGetMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (x *CartGetMessage) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CartGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cart *data.Cart `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
}

func (x *CartGetResponse) Reset() {
	*x = CartGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartGetResponse) ProtoMessage() {}

func (x *CartGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId uint32 `protobuf:"varint,1,opt,name=storeId,proto3" json:"storeId,omitempty"`
}

func (x *CartCreateMessage) Reset() {
	*x = CartCreateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartCreateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartCreateMessage) ProtoMessage() {}

func (x *CartCreateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (x *CartCreateMessage) GetStoreId() uint32 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

type CartCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cart *data.Cart `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
}

func (x *CartCreateResponse) Reset() {
	*x = CartCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartCreateResponse) ProtoMessage() {}

func (x *CartCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartId          uint32  `protobuf:"varint,1,opt,name=cartId,proto3" json:"cartId,omitempty"`
	Quantity        uint32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	ProductId       *uint32 `protobuf:"varint,3,opt,name=productId,proto3,oneof" json:"productId,omitempty"`
	ProductCustomId *string `protobuf:"bytes,4,opt,name=productCustomId,proto3,oneof" json:"productCustomId,omitempty"`
}

func (x *AddItemMessage) Reset() {
	*x = AddItemMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemMessage) ProtoMessage() {}

func (x *AddItemMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (x *AddItemMessage) GetCartId() uint32 {
	if x != nil {
		return x.CartId
	}
	return 0
}

func (x *AddItemMessage) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *AddItemMessage) GetProductId() uint32 {
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

type AddItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cart *data.Cart `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
}

func (x *AddItemResponse) Reset() {
	*x = AddItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemResponse) ProtoMessage() {}

func (x *AddItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartId          uint32  `protobuf:"varint,1,opt,name=cartId,proto3" json:"cartId,omitempty"`
	ProductId       *uint32 `protobuf:"varint,2,opt,name=productId,proto3,oneof" json:"productId,omitempty"`
	ProductCustomId *string `protobuf:"bytes,3,opt,name=productCustomId,proto3,oneof" json:"productCustomId,omitempty"`
}

func (x *RemoveItemMessage) Reset() {
	*x = RemoveItemMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveItemMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveItemMessage) ProtoMessage() {}

func (x *RemoveItemMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (x *RemoveItemMessage) GetCartId() uint32 {
	if x != nil {
		return x.CartId
	}
	return 0
}

func (x *RemoveItemMessage) GetProductId() uint32 {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cart *data.Cart `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
}

func (x *RemoveItemResponse) Reset() {
	*x = RemoveItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_cart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveItemResponse) ProtoMessage() {}

func (x *RemoveItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartId   uint32 `protobuf:"varint,1,opt,name=cartId,proto3" json:"cartId,omitempty"`
	ItemId   uint32 `protobuf:"varint,2,opt,name=itemId,proto3" json:"itemId,omitempty"`
	Quantity uint32 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *UpdateItemQuantityMessage) Reset() {
	*x = UpdateItemQuantityMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_cart_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemQuantityMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemQuantityMessage) ProtoMessage() {}

func (x *UpdateItemQuantityMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (x *UpdateItemQuantityMessage) GetCartId() uint32 {
	if x != nil {
		return x.CartId
	}
	return 0
}

func (x *UpdateItemQuantityMessage) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *UpdateItemQuantityMessage) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type UpdateItemQuantityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cart *data.Cart `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
}

func (x *UpdateItemQuantityResponse) Reset() {
	*x = UpdateItemQuantityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_cart_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemQuantityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemQuantityResponse) ProtoMessage() {}

func (x *UpdateItemQuantityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cart_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
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

var File_proto_message_cart_proto protoreflect.FileDescriptor

var file_proto_message_cart_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e,
	0x43, 0x61, 0x72, 0x74, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c,
	0x0a, 0x0f, 0x43, 0x61, 0x72, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x19, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x2d, 0x0a, 0x11,
	0x43, 0x61, 0x72, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x12, 0x43,
	0x61, 0x72, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x19, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0xb8, 0x01, 0x0a,
	0x0e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x63, 0x61,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x9f, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x61, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x61, 0x72,
	0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x67, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x69,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x22, 0x37, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_message_cart_proto_rawDescOnce sync.Once
	file_proto_message_cart_proto_rawDescData = file_proto_message_cart_proto_rawDesc
)

func file_proto_message_cart_proto_rawDescGZIP() []byte {
	file_proto_message_cart_proto_rawDescOnce.Do(func() {
		file_proto_message_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_cart_proto_rawDescData)
	})
	return file_proto_message_cart_proto_rawDescData
}

var file_proto_message_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_message_cart_proto_goTypes = []interface{}{
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
	(*data.Cart)(nil),                  // 10: Cart
}
var file_proto_message_cart_proto_depIdxs = []int32{
	10, // 0: CartGetResponse.cart:type_name -> Cart
	10, // 1: CartCreateResponse.cart:type_name -> Cart
	10, // 2: AddItemResponse.cart:type_name -> Cart
	10, // 3: RemoveItemResponse.cart:type_name -> Cart
	10, // 4: UpdateItemQuantityResponse.cart:type_name -> Cart
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_message_cart_proto_init() }
func file_proto_message_cart_proto_init() {
	if File_proto_message_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_message_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartGetMessage); i {
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
		file_proto_message_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartGetResponse); i {
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
		file_proto_message_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartCreateMessage); i {
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
		file_proto_message_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartCreateResponse); i {
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
		file_proto_message_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemMessage); i {
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
		file_proto_message_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemResponse); i {
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
		file_proto_message_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveItemMessage); i {
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
		file_proto_message_cart_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveItemResponse); i {
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
		file_proto_message_cart_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemQuantityMessage); i {
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
		file_proto_message_cart_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemQuantityResponse); i {
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
	file_proto_message_cart_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_proto_message_cart_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_message_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_message_cart_proto_goTypes,
		DependencyIndexes: file_proto_message_cart_proto_depIdxs,
		MessageInfos:      file_proto_message_cart_proto_msgTypes,
	}.Build()
	File_proto_message_cart_proto = out.File
	file_proto_message_cart_proto_rawDesc = nil
	file_proto_message_cart_proto_goTypes = nil
	file_proto_message_cart_proto_depIdxs = nil
}
