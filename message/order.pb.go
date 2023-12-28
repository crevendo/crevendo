// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: proto/message/order.proto

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

type OrderCreateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     uint32  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	CartId     uint32  `protobuf:"varint,2,opt,name=cartId,proto3" json:"cartId,omitempty"`
	Gateway    string  `protobuf:"bytes,3,opt,name=Gateway,proto3" json:"Gateway,omitempty"`
	AddressId  uint32  `protobuf:"varint,4,opt,name=addressId,proto3" json:"addressId,omitempty"`
	ShipmentId uint32  `protobuf:"varint,5,opt,name=shipmentId,proto3" json:"shipmentId,omitempty"`
	MethodId   string  `protobuf:"bytes,6,opt,name=methodId,proto3" json:"methodId,omitempty"`
	Cvv        *string `protobuf:"bytes,7,opt,name=cvv,proto3,oneof" json:"cvv,omitempty"`
	Total      float64 `protobuf:"fixed64,8,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *OrderCreateMessage) Reset() {
	*x = OrderCreateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateMessage) ProtoMessage() {}

func (x *OrderCreateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateMessage.ProtoReflect.Descriptor instead.
func (*OrderCreateMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderCreateMessage) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderCreateMessage) GetCartId() uint32 {
	if x != nil {
		return x.CartId
	}
	return 0
}

func (x *OrderCreateMessage) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

func (x *OrderCreateMessage) GetAddressId() uint32 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *OrderCreateMessage) GetShipmentId() uint32 {
	if x != nil {
		return x.ShipmentId
	}
	return 0
}

func (x *OrderCreateMessage) GetMethodId() string {
	if x != nil {
		return x.MethodId
	}
	return ""
}

func (x *OrderCreateMessage) GetCvv() string {
	if x != nil && x.Cvv != nil {
		return *x.Cvv
	}
	return ""
}

func (x *OrderCreateMessage) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type OrderCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *data.Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *OrderCreateResponse) Reset() {
	*x = OrderCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateResponse) ProtoMessage() {}

func (x *OrderCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateResponse.ProtoReflect.Descriptor instead.
func (*OrderCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderCreateResponse) GetOrder() *data.Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type OrderGetMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrderGetMessage) Reset() {
	*x = OrderGetMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderGetMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderGetMessage) ProtoMessage() {}

func (x *OrderGetMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderGetMessage.ProtoReflect.Descriptor instead.
func (*OrderGetMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderGetMessage) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OrderGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *data.Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *OrderGetResponse) Reset() {
	*x = OrderGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderGetResponse) ProtoMessage() {}

func (x *OrderGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderGetResponse.ProtoReflect.Descriptor instead.
func (*OrderGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_order_proto_rawDescGZIP(), []int{3}
}

func (x *OrderGetResponse) GetOrder() *data.Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type OrderListMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentId *uint32 `protobuf:"varint,1,opt,name=paymentId,proto3,oneof" json:"paymentId,omitempty"`
	UserId    *uint32 `protobuf:"varint,2,opt,name=userId,proto3,oneof" json:"userId,omitempty"`
	Id        *uint32 `protobuf:"varint,3,opt,name=id,proto3,oneof" json:"id,omitempty"`
}

func (x *OrderListMessage) Reset() {
	*x = OrderListMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderListMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderListMessage) ProtoMessage() {}

func (x *OrderListMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderListMessage.ProtoReflect.Descriptor instead.
func (*OrderListMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_order_proto_rawDescGZIP(), []int{4}
}

func (x *OrderListMessage) GetPaymentId() uint32 {
	if x != nil && x.PaymentId != nil {
		return *x.PaymentId
	}
	return 0
}

func (x *OrderListMessage) GetUserId() uint32 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *OrderListMessage) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

type OrderListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*data.Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *OrderListResponse) Reset() {
	*x = OrderListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderListResponse) ProtoMessage() {}

func (x *OrderListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderListResponse.ProtoReflect.Descriptor instead.
func (*OrderListResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_order_proto_rawDescGZIP(), []int{5}
}

func (x *OrderListResponse) GetOrders() []*data.Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type OrderUpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status uint32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OrderUpdateMessage) Reset() {
	*x = OrderUpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderUpdateMessage) ProtoMessage() {}

func (x *OrderUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderUpdateMessage.ProtoReflect.Descriptor instead.
func (*OrderUpdateMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_order_proto_rawDescGZIP(), []int{6}
}

func (x *OrderUpdateMessage) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderUpdateMessage) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type OrderUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderUpdateResponse) Reset() {
	*x = OrderUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderUpdateResponse) ProtoMessage() {}

func (x *OrderUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderUpdateResponse.ProtoReflect.Descriptor instead.
func (*OrderUpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_order_proto_rawDescGZIP(), []int{7}
}

type OrderStatusListMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderStatusListMessage) Reset() {
	*x = OrderStatusListMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_order_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderStatusListMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStatusListMessage) ProtoMessage() {}

func (x *OrderStatusListMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_order_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderStatusListMessage.ProtoReflect.Descriptor instead.
func (*OrderStatusListMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_order_proto_rawDescGZIP(), []int{8}
}

type OrderStatusListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderStatuses []*data.OrderStatus `protobuf:"bytes,1,rep,name=orderStatuses,proto3" json:"orderStatuses,omitempty"`
}

func (x *OrderStatusListResponse) Reset() {
	*x = OrderStatusListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_order_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderStatusListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStatusListResponse) ProtoMessage() {}

func (x *OrderStatusListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_order_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderStatusListResponse.ProtoReflect.Descriptor instead.
func (*OrderStatusListResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_order_proto_rawDescGZIP(), []int{9}
}

func (x *OrderStatusListResponse) GetOrderStatuses() []*data.OrderStatus {
	if x != nil {
		return x.OrderStatuses
	}
	return nil
}

type OrderItemStatusListMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderItemStatusListMessage) Reset() {
	*x = OrderItemStatusListMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_order_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemStatusListMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemStatusListMessage) ProtoMessage() {}

func (x *OrderItemStatusListMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_order_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemStatusListMessage.ProtoReflect.Descriptor instead.
func (*OrderItemStatusListMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_order_proto_rawDescGZIP(), []int{10}
}

type OrderItemStatusListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderItemStatuses []*data.OrderItemStatus `protobuf:"bytes,1,rep,name=OrderItemStatuses,proto3" json:"OrderItemStatuses,omitempty"`
}

func (x *OrderItemStatusListResponse) Reset() {
	*x = OrderItemStatusListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_order_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemStatusListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemStatusListResponse) ProtoMessage() {}

func (x *OrderItemStatusListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_order_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemStatusListResponse.ProtoReflect.Descriptor instead.
func (*OrderItemStatusListResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_order_proto_rawDescGZIP(), []int{11}
}

func (x *OrderItemStatusListResponse) GetOrderItemStatuses() []*data.OrderItemStatus {
	if x != nil {
		return x.OrderItemStatuses
	}
	return nil
}

var File_proto_message_order_proto protoreflect.FileDescriptor

var file_proto_message_order_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed,
	0x01, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63,
	0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x03, 0x63, 0x76, 0x76,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x63, 0x76, 0x76, 0x88, 0x01, 0x01,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x63, 0x76, 0x76, 0x22, 0x33,
	0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x22, 0x21, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x87, 0x01, 0x0a, 0x10, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a,
	0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x00, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x02, 0x69, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x05, 0x0a, 0x03, 0x5f,
	0x69, 0x64, 0x22, 0x33, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x3c, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x0a, 0x16,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4d, 0x0a, 0x17, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x65, 0x73, 0x22, 0x1c, 0x0a, 0x1a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x5d, 0x0a, 0x1b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3e, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x65, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_message_order_proto_rawDescOnce sync.Once
	file_proto_message_order_proto_rawDescData = file_proto_message_order_proto_rawDesc
)

func file_proto_message_order_proto_rawDescGZIP() []byte {
	file_proto_message_order_proto_rawDescOnce.Do(func() {
		file_proto_message_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_order_proto_rawDescData)
	})
	return file_proto_message_order_proto_rawDescData
}

var file_proto_message_order_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_message_order_proto_goTypes = []interface{}{
	(*OrderCreateMessage)(nil),          // 0: OrderCreateMessage
	(*OrderCreateResponse)(nil),         // 1: OrderCreateResponse
	(*OrderGetMessage)(nil),             // 2: OrderGetMessage
	(*OrderGetResponse)(nil),            // 3: OrderGetResponse
	(*OrderListMessage)(nil),            // 4: OrderListMessage
	(*OrderListResponse)(nil),           // 5: OrderListResponse
	(*OrderUpdateMessage)(nil),          // 6: OrderUpdateMessage
	(*OrderUpdateResponse)(nil),         // 7: OrderUpdateResponse
	(*OrderStatusListMessage)(nil),      // 8: OrderStatusListMessage
	(*OrderStatusListResponse)(nil),     // 9: OrderStatusListResponse
	(*OrderItemStatusListMessage)(nil),  // 10: OrderItemStatusListMessage
	(*OrderItemStatusListResponse)(nil), // 11: OrderItemStatusListResponse
	(*data.Order)(nil),                  // 12: Order
	(*data.OrderStatus)(nil),            // 13: OrderStatus
	(*data.OrderItemStatus)(nil),        // 14: OrderItemStatus
}
var file_proto_message_order_proto_depIdxs = []int32{
	12, // 0: OrderCreateResponse.order:type_name -> Order
	12, // 1: OrderGetResponse.order:type_name -> Order
	12, // 2: OrderListResponse.orders:type_name -> Order
	13, // 3: OrderStatusListResponse.orderStatuses:type_name -> OrderStatus
	14, // 4: OrderItemStatusListResponse.OrderItemStatuses:type_name -> OrderItemStatus
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_message_order_proto_init() }
func file_proto_message_order_proto_init() {
	if File_proto_message_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_message_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateMessage); i {
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
		file_proto_message_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateResponse); i {
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
		file_proto_message_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderGetMessage); i {
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
		file_proto_message_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderGetResponse); i {
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
		file_proto_message_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderListMessage); i {
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
		file_proto_message_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderListResponse); i {
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
		file_proto_message_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderUpdateMessage); i {
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
		file_proto_message_order_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderUpdateResponse); i {
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
		file_proto_message_order_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderStatusListMessage); i {
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
		file_proto_message_order_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderStatusListResponse); i {
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
		file_proto_message_order_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemStatusListMessage); i {
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
		file_proto_message_order_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemStatusListResponse); i {
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
	file_proto_message_order_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_message_order_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_message_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_message_order_proto_goTypes,
		DependencyIndexes: file_proto_message_order_proto_depIdxs,
		MessageInfos:      file_proto_message_order_proto_msgTypes,
	}.Build()
	File_proto_message_order_proto = out.File
	file_proto_message_order_proto_rawDesc = nil
	file_proto_message_order_proto_goTypes = nil
	file_proto_message_order_proto_depIdxs = nil
}
