// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/category.proto

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

type Category struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Parent        *Category              `protobuf:"bytes,4,opt,name=parent,proto3" json:"parent,omitempty"`
	ParentId      uint32                 `protobuf:"varint,5,opt,name=parentId,proto3" json:"parentId,omitempty"`
	Children      []*Category            `protobuf:"bytes,6,rep,name=children,proto3" json:"children,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Category) Reset() {
	*x = Category{}
	mi := &file_proto_category_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_proto_category_proto_rawDescGZIP(), []int{0}
}

func (x *Category) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetParent() *Category {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *Category) GetParentId() uint32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Category) GetChildren() []*Category {
	if x != nil {
		return x.Children
	}
	return nil
}

var File_proto_category_proto protoreflect.FileDescriptor

const file_proto_category_proto_rawDesc = "" +
	"\n" +
	"\x14proto/category.proto\"\x94\x01\n" +
	"\bCategory\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12!\n" +
	"\x06parent\x18\x04 \x01(\v2\t.CategoryR\x06parent\x12\x1a\n" +
	"\bparentId\x18\x05 \x01(\rR\bparentId\x12%\n" +
	"\bchildren\x18\x06 \x03(\v2\t.CategoryR\bchildrenB#Z!github.com/crevendo/crevendo/datab\x06proto3"

var (
	file_proto_category_proto_rawDescOnce sync.Once
	file_proto_category_proto_rawDescData []byte
)

func file_proto_category_proto_rawDescGZIP() []byte {
	file_proto_category_proto_rawDescOnce.Do(func() {
		file_proto_category_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_category_proto_rawDesc), len(file_proto_category_proto_rawDesc)))
	})
	return file_proto_category_proto_rawDescData
}

var file_proto_category_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_category_proto_goTypes = []any{
	(*Category)(nil), // 0: Category
}
var file_proto_category_proto_depIdxs = []int32{
	0, // 0: Category.parent:type_name -> Category
	0, // 1: Category.children:type_name -> Category
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_category_proto_init() }
func file_proto_category_proto_init() {
	if File_proto_category_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_category_proto_rawDesc), len(file_proto_category_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_category_proto_goTypes,
		DependencyIndexes: file_proto_category_proto_depIdxs,
		MessageInfos:      file_proto_category_proto_msgTypes,
	}.Build()
	File_proto_category_proto = out.File
	file_proto_category_proto_goTypes = nil
	file_proto_category_proto_depIdxs = nil
}
