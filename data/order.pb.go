// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/order.proto

package data

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type OrderStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Key           string                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Label         string                 `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderStatus) Reset() {
	*x = OrderStatus{}
	mi := &file_proto_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStatus) ProtoMessage() {}

func (x *OrderStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderStatus.ProtoReflect.Descriptor instead.
func (*OrderStatus) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderStatus) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderStatus) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *OrderStatus) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type OrderItemStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Key           string                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Label         string                 `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderItemStatus) Reset() {
	*x = OrderItemStatus{}
	mi := &file_proto_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItemStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemStatus) ProtoMessage() {}

func (x *OrderItemStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemStatus.ProtoReflect.Descriptor instead.
func (*OrderItemStatus) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderItemStatus) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderItemStatus) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *OrderItemStatus) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid          string                 `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserId        int32                  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	AddressId     int32                  `protobuf:"varint,4,opt,name=addressId,proto3" json:"addressId,omitempty"`
	Items         []*OrderItem           `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
	Total         float64                `protobuf:"fixed64,6,opt,name=total,proto3" json:"total,omitempty"`
	PaymentId     int32                  `protobuf:"varint,7,opt,name=paymentId,proto3" json:"paymentId,omitempty"`
	Status        *OrderStatus           `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Fees          []*Fee                 `protobuf:"bytes,9,rep,name=fees,proto3" json:"fees,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Data          []*Data                `protobuf:"bytes,14,rep,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_proto_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{2}
}

func (x *Order) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Order) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Order) GetAddressId() int32 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *Order) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Order) GetPaymentId() int32 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *Order) GetStatus() *OrderStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Order) GetFees() []*Fee {
	if x != nil {
		return x.Fees
	}
	return nil
}

func (x *Order) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Order) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Order) GetData() []*Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrderItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId       int32                  `protobuf:"varint,3,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Name          string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Image         string                 `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Quantity      int32                  `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price         float64                `protobuf:"fixed64,7,opt,name=price,proto3" json:"price,omitempty"`
	ProductId     int32                  `protobuf:"varint,8,opt,name=productId,proto3" json:"productId,omitempty"`
	Data          []*Data                `protobuf:"bytes,9,rep,name=data,proto3" json:"data,omitempty"`
	Status        *OrderItemStatus       `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	mi := &file_proto_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{3}
}

func (x *OrderItem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderItem) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderItem) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *OrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderItem) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderItem) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *OrderItem) GetData() []*Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *OrderItem) GetStatus() *OrderItemStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_proto_order_proto protoreflect.FileDescriptor

const file_proto_order_proto_rawDesc = "" +
	"\n" +
	"\x11proto/order.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x10proto/data.proto\x1a\x0fproto/fee.proto\"E\n" +
	"\vOrderStatus\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x10\n" +
	"\x03key\x18\x02 \x01(\tR\x03key\x12\x14\n" +
	"\x05label\x18\x03 \x01(\tR\x05label\"I\n" +
	"\x0fOrderItemStatus\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x10\n" +
	"\x03key\x18\x02 \x01(\tR\x03key\x12\x14\n" +
	"\x05label\x18\x03 \x01(\tR\x05label\"\x86\x03\n" +
	"\x05Order\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04uuid\x18\x02 \x01(\tR\x04uuid\x12\x16\n" +
	"\x06userId\x18\x03 \x01(\x05R\x06userId\x12\x1c\n" +
	"\taddressId\x18\x04 \x01(\x05R\taddressId\x12 \n" +
	"\x05items\x18\x05 \x03(\v2\n" +
	".OrderItemR\x05items\x12\x14\n" +
	"\x05total\x18\x06 \x01(\x01R\x05total\x12\x1c\n" +
	"\tpaymentId\x18\a \x01(\x05R\tpaymentId\x12$\n" +
	"\x06status\x18\b \x01(\v2\f.OrderStatusR\x06status\x12\x18\n" +
	"\x04fees\x18\t \x03(\v2\x04.FeeR\x04fees\x128\n" +
	"\tcreatedAt\x18\f \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x128\n" +
	"\tupdatedAt\x18\r \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x19\n" +
	"\x04data\x18\x0e \x03(\v2\x05.DataR\x04data\"\xf4\x01\n" +
	"\tOrderItem\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x18\n" +
	"\aorderId\x18\x03 \x01(\x05R\aorderId\x12\x12\n" +
	"\x04name\x18\x04 \x01(\tR\x04name\x12\x14\n" +
	"\x05image\x18\x05 \x01(\tR\x05image\x12\x1a\n" +
	"\bquantity\x18\x06 \x01(\x05R\bquantity\x12\x14\n" +
	"\x05price\x18\a \x01(\x01R\x05price\x12\x1c\n" +
	"\tproductId\x18\b \x01(\x05R\tproductId\x12\x19\n" +
	"\x04data\x18\t \x03(\v2\x05.DataR\x04data\x12(\n" +
	"\x06status\x18\n" +
	" \x01(\v2\x10.OrderItemStatusR\x06statusB#Z!github.com/crevendo/crevendo/datab\x06proto3"

var (
	file_proto_order_proto_rawDescOnce sync.Once
	file_proto_order_proto_rawDescData []byte
)

func file_proto_order_proto_rawDescGZIP() []byte {
	file_proto_order_proto_rawDescOnce.Do(func() {
		file_proto_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_order_proto_rawDesc), len(file_proto_order_proto_rawDesc)))
	})
	return file_proto_order_proto_rawDescData
}

var file_proto_order_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_order_proto_goTypes = []any{
	(*OrderStatus)(nil),           // 0: OrderStatus
	(*OrderItemStatus)(nil),       // 1: OrderItemStatus
	(*Order)(nil),                 // 2: Order
	(*OrderItem)(nil),             // 3: OrderItem
	(*Fee)(nil),                   // 4: Fee
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*Data)(nil),                  // 6: Data
}
var file_proto_order_proto_depIdxs = []int32{
	3, // 0: Order.items:type_name -> OrderItem
	0, // 1: Order.status:type_name -> OrderStatus
	4, // 2: Order.fees:type_name -> Fee
	5, // 3: Order.createdAt:type_name -> google.protobuf.Timestamp
	5, // 4: Order.updatedAt:type_name -> google.protobuf.Timestamp
	6, // 5: Order.data:type_name -> Data
	6, // 6: OrderItem.data:type_name -> Data
	1, // 7: OrderItem.status:type_name -> OrderItemStatus
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_proto_order_proto_init() }
func file_proto_order_proto_init() {
	if File_proto_order_proto != nil {
		return
	}
	file_proto_data_proto_init()
	file_proto_fee_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_order_proto_rawDesc), len(file_proto_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_order_proto_goTypes,
		DependencyIndexes: file_proto_order_proto_depIdxs,
		MessageInfos:      file_proto_order_proto_msgTypes,
	}.Build()
	File_proto_order_proto = out.File
	file_proto_order_proto_goTypes = nil
	file_proto_order_proto_depIdxs = nil
}
