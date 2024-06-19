// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: proto/message/category.proto

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

type CategoryListMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IgnoreCache *bool `protobuf:"varint,1,opt,name=ignoreCache,proto3,oneof" json:"ignoreCache,omitempty"`
}

func (x *CategoryListMessage) Reset() {
	*x = CategoryListMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryListMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryListMessage) ProtoMessage() {}

func (x *CategoryListMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryListMessage.ProtoReflect.Descriptor instead.
func (*CategoryListMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_category_proto_rawDescGZIP(), []int{0}
}

func (x *CategoryListMessage) GetIgnoreCache() bool {
	if x != nil && x.IgnoreCache != nil {
		return *x.IgnoreCache
	}
	return false
}

type CategoryListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories []*data.Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *CategoryListResponse) Reset() {
	*x = CategoryListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryListResponse) ProtoMessage() {}

func (x *CategoryListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryListResponse.ProtoReflect.Descriptor instead.
func (*CategoryListResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_category_proto_rawDescGZIP(), []int{1}
}

func (x *CategoryListResponse) GetCategories() []*data.Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

type CategoryGetMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IgnoreCache bool   `protobuf:"varint,2,opt,name=ignoreCache,proto3" json:"ignoreCache,omitempty"`
}

func (x *CategoryGetMessage) Reset() {
	*x = CategoryGetMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryGetMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryGetMessage) ProtoMessage() {}

func (x *CategoryGetMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryGetMessage.ProtoReflect.Descriptor instead.
func (*CategoryGetMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_category_proto_rawDescGZIP(), []int{2}
}

func (x *CategoryGetMessage) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CategoryGetMessage) GetIgnoreCache() bool {
	if x != nil {
		return x.IgnoreCache
	}
	return false
}

type CategoryGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *data.Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *CategoryGetResponse) Reset() {
	*x = CategoryGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryGetResponse) ProtoMessage() {}

func (x *CategoryGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryGetResponse.ProtoReflect.Descriptor instead.
func (*CategoryGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_category_proto_rawDescGZIP(), []int{3}
}

func (x *CategoryGetResponse) GetCategory() *data.Category {
	if x != nil {
		return x.Category
	}
	return nil
}

type CategoryCreateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ParentId *uint32 `protobuf:"varint,2,opt,name=parentId,proto3,oneof" json:"parentId,omitempty"`
}

func (x *CategoryCreateMessage) Reset() {
	*x = CategoryCreateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryCreateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryCreateMessage) ProtoMessage() {}

func (x *CategoryCreateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryCreateMessage.ProtoReflect.Descriptor instead.
func (*CategoryCreateMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_category_proto_rawDescGZIP(), []int{4}
}

func (x *CategoryCreateMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CategoryCreateMessage) GetParentId() uint32 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

type CategoryCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *data.Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *CategoryCreateResponse) Reset() {
	*x = CategoryCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_category_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryCreateResponse) ProtoMessage() {}

func (x *CategoryCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_category_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryCreateResponse.ProtoReflect.Descriptor instead.
func (*CategoryCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_category_proto_rawDescGZIP(), []int{5}
}

func (x *CategoryCreateResponse) GetCategory() *data.Category {
	if x != nil {
		return x.Category
	}
	return nil
}

type CategoryUpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ParentId *uint32 `protobuf:"varint,3,opt,name=parentId,proto3,oneof" json:"parentId,omitempty"`
}

func (x *CategoryUpdateMessage) Reset() {
	*x = CategoryUpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_category_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryUpdateMessage) ProtoMessage() {}

func (x *CategoryUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_category_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryUpdateMessage.ProtoReflect.Descriptor instead.
func (*CategoryUpdateMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_category_proto_rawDescGZIP(), []int{6}
}

func (x *CategoryUpdateMessage) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CategoryUpdateMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CategoryUpdateMessage) GetParentId() uint32 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

type CategoryUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *data.Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *CategoryUpdateResponse) Reset() {
	*x = CategoryUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_category_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryUpdateResponse) ProtoMessage() {}

func (x *CategoryUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_category_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryUpdateResponse.ProtoReflect.Descriptor instead.
func (*CategoryUpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_category_proto_rawDescGZIP(), []int{7}
}

func (x *CategoryUpdateResponse) GetCategory() *data.Category {
	if x != nil {
		return x.Category
	}
	return nil
}

type CategoryDeleteMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CategoryDeleteMessage) Reset() {
	*x = CategoryDeleteMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_category_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryDeleteMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryDeleteMessage) ProtoMessage() {}

func (x *CategoryDeleteMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_category_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryDeleteMessage.ProtoReflect.Descriptor instead.
func (*CategoryDeleteMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_category_proto_rawDescGZIP(), []int{8}
}

func (x *CategoryDeleteMessage) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CategoryDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CategoryDeleteResponse) Reset() {
	*x = CategoryDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_category_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryDeleteResponse) ProtoMessage() {}

func (x *CategoryDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_category_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryDeleteResponse.ProtoReflect.Descriptor instead.
func (*CategoryDeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_category_proto_rawDescGZIP(), []int{9}
}

func (x *CategoryDeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_message_category_proto protoreflect.FileDescriptor

var file_proto_message_category_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x13, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x69,
	0x67, 0x6e, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x0b, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x22, 0x41, 0x0a, 0x14, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x46, 0x0a, 0x12, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x69,
	0x67, 0x6e, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x22, 0x3c, 0x0a,
	0x13, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x59, 0x0a, 0x15, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x16, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x69, 0x0a, 0x15, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x3f, 0x0a, 0x16, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x22, 0x27, 0x0a, 0x15, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x16,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x72, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_message_category_proto_rawDescOnce sync.Once
	file_proto_message_category_proto_rawDescData = file_proto_message_category_proto_rawDesc
)

func file_proto_message_category_proto_rawDescGZIP() []byte {
	file_proto_message_category_proto_rawDescOnce.Do(func() {
		file_proto_message_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_category_proto_rawDescData)
	})
	return file_proto_message_category_proto_rawDescData
}

var file_proto_message_category_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_message_category_proto_goTypes = []interface{}{
	(*CategoryListMessage)(nil),    // 0: CategoryListMessage
	(*CategoryListResponse)(nil),   // 1: CategoryListResponse
	(*CategoryGetMessage)(nil),     // 2: CategoryGetMessage
	(*CategoryGetResponse)(nil),    // 3: CategoryGetResponse
	(*CategoryCreateMessage)(nil),  // 4: CategoryCreateMessage
	(*CategoryCreateResponse)(nil), // 5: CategoryCreateResponse
	(*CategoryUpdateMessage)(nil),  // 6: CategoryUpdateMessage
	(*CategoryUpdateResponse)(nil), // 7: CategoryUpdateResponse
	(*CategoryDeleteMessage)(nil),  // 8: CategoryDeleteMessage
	(*CategoryDeleteResponse)(nil), // 9: CategoryDeleteResponse
	(*data.Category)(nil),          // 10: Category
}
var file_proto_message_category_proto_depIdxs = []int32{
	10, // 0: CategoryListResponse.categories:type_name -> Category
	10, // 1: CategoryGetResponse.category:type_name -> Category
	10, // 2: CategoryCreateResponse.category:type_name -> Category
	10, // 3: CategoryUpdateResponse.category:type_name -> Category
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_message_category_proto_init() }
func file_proto_message_category_proto_init() {
	if File_proto_message_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_message_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryListMessage); i {
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
		file_proto_message_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryListResponse); i {
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
		file_proto_message_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryGetMessage); i {
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
		file_proto_message_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryGetResponse); i {
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
		file_proto_message_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryCreateMessage); i {
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
		file_proto_message_category_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryCreateResponse); i {
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
		file_proto_message_category_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryUpdateMessage); i {
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
		file_proto_message_category_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryUpdateResponse); i {
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
		file_proto_message_category_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryDeleteMessage); i {
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
		file_proto_message_category_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryDeleteResponse); i {
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
	file_proto_message_category_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_message_category_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_proto_message_category_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_message_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_message_category_proto_goTypes,
		DependencyIndexes: file_proto_message_category_proto_depIdxs,
		MessageInfos:      file_proto_message_category_proto_msgTypes,
	}.Build()
	File_proto_message_category_proto = out.File
	file_proto_message_category_proto_rawDesc = nil
	file_proto_message_category_proto_goTypes = nil
	file_proto_message_category_proto_depIdxs = nil
}
