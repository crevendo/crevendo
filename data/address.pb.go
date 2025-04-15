// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/address.proto

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

type Address struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        uint32                 `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Name          string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Address       string                 `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	PostalCode    uint32                 `protobuf:"varint,6,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
	Phone         string                 `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Notes         string                 `protobuf:"bytes,8,opt,name=notes,proto3" json:"notes,omitempty"`
	Fields        []*Field               `protobuf:"bytes,9,rep,name=fields,proto3" json:"fields,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address) Reset() {
	*x = Address{}
	mi := &file_proto_address_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_proto_address_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_proto_address_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Address) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Address) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Address) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Address) GetPostalCode() uint32 {
	if x != nil {
		return x.PostalCode
	}
	return 0
}

func (x *Address) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Address) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *Address) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

var File_proto_address_proto protoreflect.FileDescriptor

const file_proto_address_proto_rawDesc = "" +
	"\n" +
	"\x13proto/address.proto\x1a\x11proto/field.proto\"\xcb\x01\n" +
	"\aAddress\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x16\n" +
	"\x06userId\x18\x03 \x01(\rR\x06userId\x12\x12\n" +
	"\x04name\x18\x04 \x01(\tR\x04name\x12\x18\n" +
	"\aaddress\x18\x05 \x01(\tR\aaddress\x12\x1e\n" +
	"\n" +
	"postalCode\x18\x06 \x01(\rR\n" +
	"postalCode\x12\x14\n" +
	"\x05phone\x18\a \x01(\tR\x05phone\x12\x14\n" +
	"\x05notes\x18\b \x01(\tR\x05notes\x12\x1e\n" +
	"\x06fields\x18\t \x03(\v2\x06.FieldR\x06fieldsB#Z!github.com/crevendo/crevendo/datab\x06proto3"

var (
	file_proto_address_proto_rawDescOnce sync.Once
	file_proto_address_proto_rawDescData []byte
)

func file_proto_address_proto_rawDescGZIP() []byte {
	file_proto_address_proto_rawDescOnce.Do(func() {
		file_proto_address_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_address_proto_rawDesc), len(file_proto_address_proto_rawDesc)))
	})
	return file_proto_address_proto_rawDescData
}

var file_proto_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_address_proto_goTypes = []any{
	(*Address)(nil), // 0: Address
	(*Field)(nil),   // 1: Field
}
var file_proto_address_proto_depIdxs = []int32{
	1, // 0: Address.fields:type_name -> Field
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_address_proto_init() }
func file_proto_address_proto_init() {
	if File_proto_address_proto != nil {
		return
	}
	file_proto_field_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_address_proto_rawDesc), len(file_proto_address_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_address_proto_goTypes,
		DependencyIndexes: file_proto_address_proto_depIdxs,
		MessageInfos:      file_proto_address_proto_msgTypes,
	}.Build()
	File_proto_address_proto = out.File
	file_proto_address_proto_goTypes = nil
	file_proto_address_proto_depIdxs = nil
}
