// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: proto/message/product.proto

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

type ProductGetMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           *uint32 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	CustomId     *string `protobuf:"bytes,2,opt,name=customId,proto3,oneof" json:"customId,omitempty"`
	WithVariants *bool   `protobuf:"varint,3,opt,name=withVariants,proto3,oneof" json:"withVariants,omitempty"`
}

func (x *ProductGetMessage) Reset() {
	*x = ProductGetMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductGetMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductGetMessage) ProtoMessage() {}

func (x *ProductGetMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductGetMessage.ProtoReflect.Descriptor instead.
func (*ProductGetMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductGetMessage) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *ProductGetMessage) GetCustomId() string {
	if x != nil && x.CustomId != nil {
		return *x.CustomId
	}
	return ""
}

func (x *ProductGetMessage) GetWithVariants() bool {
	if x != nil && x.WithVariants != nil {
		return *x.WithVariants
	}
	return false
}

type ProductGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *data.Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *ProductGetResponse) Reset() {
	*x = ProductGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductGetResponse) ProtoMessage() {}

func (x *ProductGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductGetResponse.ProtoReflect.Descriptor instead.
func (*ProductGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductGetResponse) GetProduct() *data.Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type ProductListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products  []*data.Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	Page      *uint32         `protobuf:"varint,2,opt,name=page,proto3,oneof" json:"page,omitempty"`
	PageTotal *uint32         `protobuf:"varint,3,opt,name=pageTotal,proto3,oneof" json:"pageTotal,omitempty"`
}

func (x *ProductListResponse) Reset() {
	*x = ProductListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductListResponse) ProtoMessage() {}

func (x *ProductListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductListResponse.ProtoReflect.Descriptor instead.
func (*ProductListResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_product_proto_rawDescGZIP(), []int{2}
}

func (x *ProductListResponse) GetProducts() []*data.Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *ProductListResponse) GetPage() uint32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *ProductListResponse) GetPageTotal() uint32 {
	if x != nil && x.PageTotal != nil {
		return *x.PageTotal
	}
	return 0
}

type ProductListMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids         []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	CustomIds   []string `protobuf:"bytes,2,rep,name=customIds,proto3" json:"customIds,omitempty"`
	CategoryId  *uint32  `protobuf:"varint,3,opt,name=categoryId,proto3,oneof" json:"categoryId,omitempty"`
	Type        *string  `protobuf:"bytes,4,opt,name=type,proto3,oneof" json:"type,omitempty"`
	SearchTerm  *string  `protobuf:"bytes,6,opt,name=searchTerm,proto3,oneof" json:"searchTerm,omitempty"`
	Page        *uint32  `protobuf:"varint,7,opt,name=page,proto3,oneof" json:"page,omitempty"`
	Limit       *uint32  `protobuf:"varint,8,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	OrderBy     *uint32  `protobuf:"varint,9,opt,name=orderBy,proto3,oneof" json:"orderBy,omitempty"`
	IgnoreCache *bool    `protobuf:"varint,10,opt,name=ignoreCache,proto3,oneof" json:"ignoreCache,omitempty"`
	SessionUUID *string  `protobuf:"bytes,11,opt,name=sessionUUID,proto3,oneof" json:"sessionUUID,omitempty"`
	TimeLimit   *uint32  `protobuf:"varint,12,opt,name=timeLimit,proto3,oneof" json:"timeLimit,omitempty"`
}

func (x *ProductListMessage) Reset() {
	*x = ProductListMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductListMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductListMessage) ProtoMessage() {}

func (x *ProductListMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductListMessage.ProtoReflect.Descriptor instead.
func (*ProductListMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_product_proto_rawDescGZIP(), []int{3}
}

func (x *ProductListMessage) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ProductListMessage) GetCustomIds() []string {
	if x != nil {
		return x.CustomIds
	}
	return nil
}

func (x *ProductListMessage) GetCategoryId() uint32 {
	if x != nil && x.CategoryId != nil {
		return *x.CategoryId
	}
	return 0
}

func (x *ProductListMessage) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *ProductListMessage) GetSearchTerm() string {
	if x != nil && x.SearchTerm != nil {
		return *x.SearchTerm
	}
	return ""
}

func (x *ProductListMessage) GetPage() uint32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *ProductListMessage) GetLimit() uint32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *ProductListMessage) GetOrderBy() uint32 {
	if x != nil && x.OrderBy != nil {
		return *x.OrderBy
	}
	return 0
}

func (x *ProductListMessage) GetIgnoreCache() bool {
	if x != nil && x.IgnoreCache != nil {
		return *x.IgnoreCache
	}
	return false
}

func (x *ProductListMessage) GetSessionUUID() string {
	if x != nil && x.SessionUUID != nil {
		return *x.SessionUUID
	}
	return ""
}

func (x *ProductListMessage) GetTimeLimit() uint32 {
	if x != nil && x.TimeLimit != nil {
		return *x.TimeLimit
	}
	return 0
}

type ProductUpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *data.Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *ProductUpdateMessage) Reset() {
	*x = ProductUpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductUpdateMessage) ProtoMessage() {}

func (x *ProductUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductUpdateMessage.ProtoReflect.Descriptor instead.
func (*ProductUpdateMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_product_proto_rawDescGZIP(), []int{4}
}

func (x *ProductUpdateMessage) GetProduct() *data.Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type ProductUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *data.Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *ProductUpdateResponse) Reset() {
	*x = ProductUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductUpdateResponse) ProtoMessage() {}

func (x *ProductUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductUpdateResponse.ProtoReflect.Descriptor instead.
func (*ProductUpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_product_proto_rawDescGZIP(), []int{5}
}

func (x *ProductUpdateResponse) GetProduct() *data.Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type ProductCreateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *data.Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *ProductCreateMessage) Reset() {
	*x = ProductCreateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCreateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCreateMessage) ProtoMessage() {}

func (x *ProductCreateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCreateMessage.ProtoReflect.Descriptor instead.
func (*ProductCreateMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_product_proto_rawDescGZIP(), []int{6}
}

func (x *ProductCreateMessage) GetProduct() *data.Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type ProductCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *data.Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *ProductCreateResponse) Reset() {
	*x = ProductCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_product_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCreateResponse) ProtoMessage() {}

func (x *ProductCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_product_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCreateResponse.ProtoReflect.Descriptor instead.
func (*ProductCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_product_proto_rawDescGZIP(), []int{7}
}

func (x *ProductCreateResponse) GetProduct() *data.Product {
	if x != nil {
		return x.Product
	}
	return nil
}

var File_proto_message_product_proto protoreflect.FileDescriptor

var file_proto_message_product_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a,
	0x02, 0x20, 0x00, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a,
	0x0c, 0x77, 0x69, 0x74, 0x68, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x0c, 0x77, 0x69, 0x74, 0x68, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x73, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x77,
	0x69, 0x74, 0x68, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x38, 0x0a, 0x12, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xdf, 0x03, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x64, 0x73, 0x12, 0x23, 0x0a,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x03,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x48, 0x06, 0x52, 0x0b, 0x69, 0x67, 0x6e, 0x6f,
	0x72, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x07, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x88, 0x01,
	0x01, 0x12, 0x21, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0d, 0x48, 0x08, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x69,
	0x67, 0x6e, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x3a, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x22, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x22, 0x3b, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x22, 0x3a, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x3b, 0x0a,
	0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_message_product_proto_rawDescOnce sync.Once
	file_proto_message_product_proto_rawDescData = file_proto_message_product_proto_rawDesc
)

func file_proto_message_product_proto_rawDescGZIP() []byte {
	file_proto_message_product_proto_rawDescOnce.Do(func() {
		file_proto_message_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_product_proto_rawDescData)
	})
	return file_proto_message_product_proto_rawDescData
}

var file_proto_message_product_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_message_product_proto_goTypes = []interface{}{
	(*ProductGetMessage)(nil),     // 0: ProductGetMessage
	(*ProductGetResponse)(nil),    // 1: ProductGetResponse
	(*ProductListResponse)(nil),   // 2: ProductListResponse
	(*ProductListMessage)(nil),    // 3: ProductListMessage
	(*ProductUpdateMessage)(nil),  // 4: ProductUpdateMessage
	(*ProductUpdateResponse)(nil), // 5: ProductUpdateResponse
	(*ProductCreateMessage)(nil),  // 6: ProductCreateMessage
	(*ProductCreateResponse)(nil), // 7: ProductCreateResponse
	(*data.Product)(nil),          // 8: Product
}
var file_proto_message_product_proto_depIdxs = []int32{
	8, // 0: ProductGetResponse.product:type_name -> Product
	8, // 1: ProductListResponse.products:type_name -> Product
	8, // 2: ProductUpdateMessage.product:type_name -> Product
	8, // 3: ProductUpdateResponse.product:type_name -> Product
	8, // 4: ProductCreateMessage.product:type_name -> Product
	8, // 5: ProductCreateResponse.product:type_name -> Product
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_message_product_proto_init() }
func file_proto_message_product_proto_init() {
	if File_proto_message_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_message_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductGetMessage); i {
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
		file_proto_message_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductGetResponse); i {
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
		file_proto_message_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductListResponse); i {
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
		file_proto_message_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductListMessage); i {
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
		file_proto_message_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductUpdateMessage); i {
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
		file_proto_message_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductUpdateResponse); i {
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
		file_proto_message_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCreateMessage); i {
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
		file_proto_message_product_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCreateResponse); i {
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
	file_proto_message_product_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_message_product_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_proto_message_product_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_message_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_message_product_proto_goTypes,
		DependencyIndexes: file_proto_message_product_proto_depIdxs,
		MessageInfos:      file_proto_message_product_proto_msgTypes,
	}.Build()
	File_proto_message_product_proto = out.File
	file_proto_message_product_proto_rawDesc = nil
	file_proto_message_product_proto_goTypes = nil
	file_proto_message_product_proto_depIdxs = nil
}
