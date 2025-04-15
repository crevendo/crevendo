// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/hook/field.proto

package hook

import (
	data "github.com/crevendo/crevendo/data"
	message "github.com/crevendo/crevendo/message"
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

type FieldHandleHookParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Fields        []*data.Field          `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FieldHandleHookParams) Reset() {
	*x = FieldHandleHookParams{}
	mi := &file_proto_hook_field_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldHandleHookParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldHandleHookParams) ProtoMessage() {}

func (x *FieldHandleHookParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hook_field_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldHandleHookParams.ProtoReflect.Descriptor instead.
func (*FieldHandleHookParams) Descriptor() ([]byte, []int) {
	return file_proto_hook_field_proto_rawDescGZIP(), []int{0}
}

func (x *FieldHandleHookParams) GetFields() []*data.Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

type FieldListHookParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Fields        []*data.Field          `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FieldListHookParams) Reset() {
	*x = FieldListHookParams{}
	mi := &file_proto_hook_field_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldListHookParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldListHookParams) ProtoMessage() {}

func (x *FieldListHookParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hook_field_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldListHookParams.ProtoReflect.Descriptor instead.
func (*FieldListHookParams) Descriptor() ([]byte, []int) {
	return file_proto_hook_field_proto_rawDescGZIP(), []int{1}
}

func (x *FieldListHookParams) GetFields() []*data.Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

var File_proto_hook_field_proto protoreflect.FileDescriptor

const file_proto_hook_field_proto_rawDesc = "" +
	"\n" +
	"\x16proto/hook/field.proto\x1a\x11proto/field.proto\x1a\x19proto/message/field.proto\"7\n" +
	"\x15FieldHandleHookParams\x12\x1e\n" +
	"\x06fields\x18\x01 \x03(\v2\x06.FieldR\x06fields\"5\n" +
	"\x13FieldListHookParams\x12\x1e\n" +
	"\x06fields\x18\x01 \x03(\v2\x06.FieldR\x06fields2\xc8\x01\n" +
	"\n" +
	"FieldHooks\x12>\n" +
	"\x10AddressFieldList\x12\x14.FieldListHookParams\x1a\x12.FieldListResponse\"\x00\x12;\n" +
	"\rUserFieldList\x12\x14.FieldListHookParams\x1a\x12.FieldListResponse\"\x00\x12=\n" +
	"\vFieldHandle\x12\x14.FieldListHookParams\x1a\x16.FieldHandleHookParams\"\x00B#Z!github.com/crevendo/crevendo/hookb\x06proto3"

var (
	file_proto_hook_field_proto_rawDescOnce sync.Once
	file_proto_hook_field_proto_rawDescData []byte
)

func file_proto_hook_field_proto_rawDescGZIP() []byte {
	file_proto_hook_field_proto_rawDescOnce.Do(func() {
		file_proto_hook_field_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_hook_field_proto_rawDesc), len(file_proto_hook_field_proto_rawDesc)))
	})
	return file_proto_hook_field_proto_rawDescData
}

var file_proto_hook_field_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_hook_field_proto_goTypes = []any{
	(*FieldHandleHookParams)(nil),     // 0: FieldHandleHookParams
	(*FieldListHookParams)(nil),       // 1: FieldListHookParams
	(*data.Field)(nil),                // 2: Field
	(*message.FieldListResponse)(nil), // 3: FieldListResponse
}
var file_proto_hook_field_proto_depIdxs = []int32{
	2, // 0: FieldHandleHookParams.fields:type_name -> Field
	2, // 1: FieldListHookParams.fields:type_name -> Field
	1, // 2: FieldHooks.AddressFieldList:input_type -> FieldListHookParams
	1, // 3: FieldHooks.UserFieldList:input_type -> FieldListHookParams
	1, // 4: FieldHooks.FieldHandle:input_type -> FieldListHookParams
	3, // 5: FieldHooks.AddressFieldList:output_type -> FieldListResponse
	3, // 6: FieldHooks.UserFieldList:output_type -> FieldListResponse
	0, // 7: FieldHooks.FieldHandle:output_type -> FieldHandleHookParams
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_hook_field_proto_init() }
func file_proto_hook_field_proto_init() {
	if File_proto_hook_field_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_hook_field_proto_rawDesc), len(file_proto_hook_field_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hook_field_proto_goTypes,
		DependencyIndexes: file_proto_hook_field_proto_depIdxs,
		MessageInfos:      file_proto_hook_field_proto_msgTypes,
	}.Build()
	File_proto_hook_field_proto = out.File
	file_proto_hook_field_proto_goTypes = nil
	file_proto_hook_field_proto_depIdxs = nil
}
