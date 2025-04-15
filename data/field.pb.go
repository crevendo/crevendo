// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/field.proto

package data

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

type Field struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Label         string                 `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Value         string                 `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Field) Reset() {
	*x = Field{}
	mi := &file_proto_field_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_proto_field_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_proto_field_proto_rawDescGZIP(), []int{0}
}

func (x *Field) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Field) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Field) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Field) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_proto_field_proto protoreflect.FileDescriptor

const file_proto_field_proto_rawDesc = "" +
	"\n" +
	"\x11proto/field.proto\"[\n" +
	"\x05Field\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x12\n" +
	"\x04type\x18\x02 \x01(\tR\x04type\x12\x14\n" +
	"\x05label\x18\x03 \x01(\tR\x05label\x12\x14\n" +
	"\x05value\x18\x04 \x01(\tR\x05valueB#Z!github.com/crevendo/crevendo/datab\x06proto3"

var (
	file_proto_field_proto_rawDescOnce sync.Once
	file_proto_field_proto_rawDescData []byte
)

func file_proto_field_proto_rawDescGZIP() []byte {
	file_proto_field_proto_rawDescOnce.Do(func() {
		file_proto_field_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_field_proto_rawDesc), len(file_proto_field_proto_rawDesc)))
	})
	return file_proto_field_proto_rawDescData
}

var file_proto_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_field_proto_goTypes = []any{
	(*Field)(nil), // 0: Field
}
var file_proto_field_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_field_proto_init() }
func file_proto_field_proto_init() {
	if File_proto_field_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_field_proto_rawDesc), len(file_proto_field_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_field_proto_goTypes,
		DependencyIndexes: file_proto_field_proto_depIdxs,
		MessageInfos:      file_proto_field_proto_msgTypes,
	}.Build()
	File_proto_field_proto = out.File
	file_proto_field_proto_goTypes = nil
	file_proto_field_proto_depIdxs = nil
}
