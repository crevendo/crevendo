// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: proto/message/query.proto

package message

import (
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

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page        *uint32 `protobuf:"varint,1,opt,name=page,proto3,oneof" json:"page,omitempty"`
	OrderBy     *string `protobuf:"bytes,2,opt,name=orderBy,proto3,oneof" json:"orderBy,omitempty"`
	TimeLimit   *uint32 `protobuf:"varint,3,opt,name=timeLimit,proto3,oneof" json:"timeLimit,omitempty"`
	SessionUUID *string `protobuf:"bytes,4,opt,name=sessionUUID,proto3,oneof" json:"sessionUUID,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_proto_message_query_proto_rawDescGZIP(), []int{0}
}

func (x *Query) GetPage() uint32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *Query) GetOrderBy() string {
	if x != nil && x.OrderBy != nil {
		return *x.OrderBy
	}
	return ""
}

func (x *Query) GetTimeLimit() uint32 {
	if x != nil && x.TimeLimit != nil {
		return *x.TimeLimit
	}
	return 0
}

func (x *Query) GetSessionUUID() string {
	if x != nil && x.SessionUUID != nil {
		return *x.SessionUUID
	}
	return ""
}

var File_proto_message_query_proto protoreflect.FileDescriptor

var file_proto_message_query_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x05,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d,
	0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x02, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x25, 0x0a, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x55, 0x55, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_message_query_proto_rawDescOnce sync.Once
	file_proto_message_query_proto_rawDescData = file_proto_message_query_proto_rawDesc
)

func file_proto_message_query_proto_rawDescGZIP() []byte {
	file_proto_message_query_proto_rawDescOnce.Do(func() {
		file_proto_message_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_query_proto_rawDescData)
	})
	return file_proto_message_query_proto_rawDescData
}

var file_proto_message_query_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_message_query_proto_goTypes = []interface{}{
	(*Query)(nil), // 0: Query
}
var file_proto_message_query_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_message_query_proto_init() }
func file_proto_message_query_proto_init() {
	if File_proto_message_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_message_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
	file_proto_message_query_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_message_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_message_query_proto_goTypes,
		DependencyIndexes: file_proto_message_query_proto_depIdxs,
		MessageInfos:      file_proto_message_query_proto_msgTypes,
	}.Build()
	File_proto_message_query_proto = out.File
	file_proto_message_query_proto_rawDesc = nil
	file_proto_message_query_proto_goTypes = nil
	file_proto_message_query_proto_depIdxs = nil
}