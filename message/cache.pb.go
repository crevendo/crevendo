// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/message/cache.proto

package message

import (
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

type CacheCleanRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CacheCleanRequest) Reset() {
	*x = CacheCleanRequest{}
	mi := &file_proto_message_cache_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CacheCleanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheCleanRequest) ProtoMessage() {}

func (x *CacheCleanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cache_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheCleanRequest.ProtoReflect.Descriptor instead.
func (*CacheCleanRequest) Descriptor() ([]byte, []int) {
	return file_proto_message_cache_proto_rawDescGZIP(), []int{0}
}

type CacheCleanResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CacheCleanResponse) Reset() {
	*x = CacheCleanResponse{}
	mi := &file_proto_message_cache_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CacheCleanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheCleanResponse) ProtoMessage() {}

func (x *CacheCleanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cache_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheCleanResponse.ProtoReflect.Descriptor instead.
func (*CacheCleanResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_cache_proto_rawDescGZIP(), []int{1}
}

type CacheProductDeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *uint32                `protobuf:"varint,1,opt,name=Id,proto3,oneof" json:"Id,omitempty"`
	CustomId      *string                `protobuf:"bytes,2,opt,name=customId,proto3,oneof" json:"customId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CacheProductDeleteRequest) Reset() {
	*x = CacheProductDeleteRequest{}
	mi := &file_proto_message_cache_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CacheProductDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheProductDeleteRequest) ProtoMessage() {}

func (x *CacheProductDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cache_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheProductDeleteRequest.ProtoReflect.Descriptor instead.
func (*CacheProductDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_message_cache_proto_rawDescGZIP(), []int{2}
}

func (x *CacheProductDeleteRequest) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *CacheProductDeleteRequest) GetCustomId() string {
	if x != nil && x.CustomId != nil {
		return *x.CustomId
	}
	return ""
}

type CacheProductDeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CacheProductDeleteResponse) Reset() {
	*x = CacheProductDeleteResponse{}
	mi := &file_proto_message_cache_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CacheProductDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheProductDeleteResponse) ProtoMessage() {}

func (x *CacheProductDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cache_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheProductDeleteResponse.ProtoReflect.Descriptor instead.
func (*CacheProductDeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_cache_proto_rawDescGZIP(), []int{3}
}

type CacheSearchDeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SearchTerm    string                 `protobuf:"bytes,1,opt,name=searchTerm,proto3" json:"searchTerm,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CacheSearchDeleteRequest) Reset() {
	*x = CacheSearchDeleteRequest{}
	mi := &file_proto_message_cache_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CacheSearchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheSearchDeleteRequest) ProtoMessage() {}

func (x *CacheSearchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cache_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheSearchDeleteRequest.ProtoReflect.Descriptor instead.
func (*CacheSearchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_message_cache_proto_rawDescGZIP(), []int{4}
}

func (x *CacheSearchDeleteRequest) GetSearchTerm() string {
	if x != nil {
		return x.SearchTerm
	}
	return ""
}

type CacheSearchDeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CacheSearchDeleteResponse) Reset() {
	*x = CacheSearchDeleteResponse{}
	mi := &file_proto_message_cache_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CacheSearchDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheSearchDeleteResponse) ProtoMessage() {}

func (x *CacheSearchDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_cache_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheSearchDeleteResponse.ProtoReflect.Descriptor instead.
func (*CacheSearchDeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_cache_proto_rawDescGZIP(), []int{5}
}

var File_proto_message_cache_proto protoreflect.FileDescriptor

const file_proto_message_cache_proto_rawDesc = "" +
	"\n" +
	"\x19proto/message/cache.proto\"\x13\n" +
	"\x11CacheCleanRequest\"\x14\n" +
	"\x12CacheCleanResponse\"e\n" +
	"\x19CacheProductDeleteRequest\x12\x13\n" +
	"\x02Id\x18\x01 \x01(\rH\x00R\x02Id\x88\x01\x01\x12\x1f\n" +
	"\bcustomId\x18\x02 \x01(\tH\x01R\bcustomId\x88\x01\x01B\x05\n" +
	"\x03_IdB\v\n" +
	"\t_customId\"\x1c\n" +
	"\x1aCacheProductDeleteResponse\":\n" +
	"\x18CacheSearchDeleteRequest\x12\x1e\n" +
	"\n" +
	"searchTerm\x18\x01 \x01(\tR\n" +
	"searchTerm\"\x1b\n" +
	"\x19CacheSearchDeleteResponseB&Z$github.com/crevendo/crevendo/messageb\x06proto3"

var (
	file_proto_message_cache_proto_rawDescOnce sync.Once
	file_proto_message_cache_proto_rawDescData []byte
)

func file_proto_message_cache_proto_rawDescGZIP() []byte {
	file_proto_message_cache_proto_rawDescOnce.Do(func() {
		file_proto_message_cache_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_message_cache_proto_rawDesc), len(file_proto_message_cache_proto_rawDesc)))
	})
	return file_proto_message_cache_proto_rawDescData
}

var file_proto_message_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_message_cache_proto_goTypes = []any{
	(*CacheCleanRequest)(nil),          // 0: CacheCleanRequest
	(*CacheCleanResponse)(nil),         // 1: CacheCleanResponse
	(*CacheProductDeleteRequest)(nil),  // 2: CacheProductDeleteRequest
	(*CacheProductDeleteResponse)(nil), // 3: CacheProductDeleteResponse
	(*CacheSearchDeleteRequest)(nil),   // 4: CacheSearchDeleteRequest
	(*CacheSearchDeleteResponse)(nil),  // 5: CacheSearchDeleteResponse
}
var file_proto_message_cache_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_message_cache_proto_init() }
func file_proto_message_cache_proto_init() {
	if File_proto_message_cache_proto != nil {
		return
	}
	file_proto_message_cache_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_message_cache_proto_rawDesc), len(file_proto_message_cache_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_message_cache_proto_goTypes,
		DependencyIndexes: file_proto_message_cache_proto_depIdxs,
		MessageInfos:      file_proto_message_cache_proto_msgTypes,
	}.Build()
	File_proto_message_cache_proto = out.File
	file_proto_message_cache_proto_goTypes = nil
	file_proto_message_cache_proto_depIdxs = nil
}
