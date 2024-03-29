// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: proto/message/brand.proto

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

type BrandListMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BrandListMessage) Reset() {
	*x = BrandListMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_brand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrandListMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandListMessage) ProtoMessage() {}

func (x *BrandListMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_brand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brands []*data.Brand `protobuf:"bytes,1,rep,name=brands,proto3" json:"brands,omitempty"`
}

func (x *BrandListResponse) Reset() {
	*x = BrandListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_brand_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrandListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandListResponse) ProtoMessage() {}

func (x *BrandListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_brand_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BrandGetMessage) Reset() {
	*x = BrandGetMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_brand_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrandGetMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandGetMessage) ProtoMessage() {}

func (x *BrandGetMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_brand_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brand *data.Brand `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
}

func (x *BrandGetResponse) Reset() {
	*x = BrandGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_brand_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrandGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandGetResponse) ProtoMessage() {}

func (x *BrandGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_brand_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Url  string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *BrandUpdateMessage) Reset() {
	*x = BrandUpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_brand_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrandUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandUpdateMessage) ProtoMessage() {}

func (x *BrandUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_brand_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BrandUpdateResponse) Reset() {
	*x = BrandUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_brand_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrandUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandUpdateResponse) ProtoMessage() {}

func (x *BrandUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_brand_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_proto_message_brand_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x12,
	0x0a, 0x10, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x33, 0x0a, 0x11, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52,
	0x06, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x21, 0x0a, 0x0f, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x10, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x22, 0x4a, 0x0a, 0x12,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x15, 0x0a, 0x13, 0x42, 0x72, 0x61, 0x6e,
	0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72,
	0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_message_brand_proto_rawDescOnce sync.Once
	file_proto_message_brand_proto_rawDescData = file_proto_message_brand_proto_rawDesc
)

func file_proto_message_brand_proto_rawDescGZIP() []byte {
	file_proto_message_brand_proto_rawDescOnce.Do(func() {
		file_proto_message_brand_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_brand_proto_rawDescData)
	})
	return file_proto_message_brand_proto_rawDescData
}

var file_proto_message_brand_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_message_brand_proto_goTypes = []interface{}{
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
	if !protoimpl.UnsafeEnabled {
		file_proto_message_brand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrandListMessage); i {
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
		file_proto_message_brand_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrandListResponse); i {
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
		file_proto_message_brand_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrandGetMessage); i {
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
		file_proto_message_brand_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrandGetResponse); i {
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
		file_proto_message_brand_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrandUpdateMessage); i {
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
		file_proto_message_brand_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrandUpdateResponse); i {
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
			RawDescriptor: file_proto_message_brand_proto_rawDesc,
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
	file_proto_message_brand_proto_rawDesc = nil
	file_proto_message_brand_proto_goTypes = nil
	file_proto_message_brand_proto_depIdxs = nil
}
