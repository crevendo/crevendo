// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/message/brand.proto

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

type BrandListMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BrandListMessage) Reset() {
	*x = BrandListMessage{}
	mi := &file_proto_message_brand_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BrandListMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandListMessage) ProtoMessage() {}

func (x *BrandListMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_brand_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrandListMessage.ProtoReflect.Descriptor instead.
func (*BrandListMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_brand_proto_rawDescGZIP(), []int{0}
}

type BrandListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Brands        []*data.Brand          `protobuf:"bytes,1,rep,name=brands,proto3" json:"brands,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BrandListResponse) Reset() {
	*x = BrandListResponse{}
	mi := &file_proto_message_brand_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BrandListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandListResponse) ProtoMessage() {}

func (x *BrandListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_brand_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrandListResponse.ProtoReflect.Descriptor instead.
func (*BrandListResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_brand_proto_rawDescGZIP(), []int{1}
}

func (x *BrandListResponse) GetBrands() []*data.Brand {
	if x != nil {
		return x.Brands
	}
	return nil
}

type BrandGetMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BrandGetMessage) Reset() {
	*x = BrandGetMessage{}
	mi := &file_proto_message_brand_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BrandGetMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandGetMessage) ProtoMessage() {}

func (x *BrandGetMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_brand_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrandGetMessage.ProtoReflect.Descriptor instead.
func (*BrandGetMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_brand_proto_rawDescGZIP(), []int{2}
}

func (x *BrandGetMessage) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BrandGetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Brand         *data.Brand            `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BrandGetResponse) Reset() {
	*x = BrandGetResponse{}
	mi := &file_proto_message_brand_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BrandGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandGetResponse) ProtoMessage() {}

func (x *BrandGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_brand_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrandGetResponse.ProtoReflect.Descriptor instead.
func (*BrandGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_brand_proto_rawDescGZIP(), []int{3}
}

func (x *BrandGetResponse) GetBrand() *data.Brand {
	if x != nil {
		return x.Brand
	}
	return nil
}

type BrandUpdateMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Url           string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BrandUpdateMessage) Reset() {
	*x = BrandUpdateMessage{}
	mi := &file_proto_message_brand_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BrandUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandUpdateMessage) ProtoMessage() {}

func (x *BrandUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_brand_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrandUpdateMessage.ProtoReflect.Descriptor instead.
func (*BrandUpdateMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_brand_proto_rawDescGZIP(), []int{4}
}

func (x *BrandUpdateMessage) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BrandUpdateMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BrandUpdateMessage) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type BrandUpdateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BrandUpdateResponse) Reset() {
	*x = BrandUpdateResponse{}
	mi := &file_proto_message_brand_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BrandUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandUpdateResponse) ProtoMessage() {}

func (x *BrandUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_brand_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrandUpdateResponse.ProtoReflect.Descriptor instead.
func (*BrandUpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_brand_proto_rawDescGZIP(), []int{5}
}

var File_proto_message_brand_proto protoreflect.FileDescriptor

const file_proto_message_brand_proto_rawDesc = "" +
	"\n" +
	"\x19proto/message/brand.proto\x1a\x11proto/brand.proto\"\x12\n" +
	"\x10BrandListMessage\"3\n" +
	"\x11BrandListResponse\x12\x1e\n" +
	"\x06brands\x18\x01 \x03(\v2\x06.BrandR\x06brands\"!\n" +
	"\x0fBrandGetMessage\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\"0\n" +
	"\x10BrandGetResponse\x12\x1c\n" +
	"\x05brand\x18\x01 \x01(\v2\x06.BrandR\x05brand\"J\n" +
	"\x12BrandUpdateMessage\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x10\n" +
	"\x03url\x18\x03 \x01(\tR\x03url\"\x15\n" +
	"\x13BrandUpdateResponseB&Z$github.com/crevendo/crevendo/messageb\x06proto3"

var (
	file_proto_message_brand_proto_rawDescOnce sync.Once
	file_proto_message_brand_proto_rawDescData []byte
)

func file_proto_message_brand_proto_rawDescGZIP() []byte {
	file_proto_message_brand_proto_rawDescOnce.Do(func() {
		file_proto_message_brand_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_message_brand_proto_rawDesc), len(file_proto_message_brand_proto_rawDesc)))
	})
	return file_proto_message_brand_proto_rawDescData
}

var file_proto_message_brand_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_message_brand_proto_goTypes = []any{
	(*BrandListMessage)(nil),    // 0: BrandListMessage
	(*BrandListResponse)(nil),   // 1: BrandListResponse
	(*BrandGetMessage)(nil),     // 2: BrandGetMessage
	(*BrandGetResponse)(nil),    // 3: BrandGetResponse
	(*BrandUpdateMessage)(nil),  // 4: BrandUpdateMessage
	(*BrandUpdateResponse)(nil), // 5: BrandUpdateResponse
	(*data.Brand)(nil),          // 6: Brand
}
var file_proto_message_brand_proto_depIdxs = []int32{
	6, // 0: BrandListResponse.brands:type_name -> Brand
	6, // 1: BrandGetResponse.brand:type_name -> Brand
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_message_brand_proto_init() }
func file_proto_message_brand_proto_init() {
	if File_proto_message_brand_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_message_brand_proto_rawDesc), len(file_proto_message_brand_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_message_brand_proto_goTypes,
		DependencyIndexes: file_proto_message_brand_proto_depIdxs,
		MessageInfos:      file_proto_message_brand_proto_msgTypes,
	}.Build()
	File_proto_message_brand_proto = out.File
	file_proto_message_brand_proto_goTypes = nil
	file_proto_message_brand_proto_depIdxs = nil
}
