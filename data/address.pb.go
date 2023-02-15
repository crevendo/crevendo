// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/address.proto

package data

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

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreId    uint32 `protobuf:"varint,2,opt,name=storeId,proto3" json:"storeId,omitempty"`
	UserId     uint32 `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Name       string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Address    string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	PostalCode uint32 `protobuf:"varint,6,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
	Phone      uint32 `protobuf:"varint,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Notes      string `protobuf:"bytes,8,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_proto_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (x *Address) GetStoreId() uint32 {
	if x != nil {
		return x.StoreId
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

func (x *Address) GetPhone() uint32 {
	if x != nil {
		return x.Phone
	}
	return 0
}

func (x *Address) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

var File_proto_address_proto protoreflect.FileDescriptor

var file_proto_address_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x42, 0x23, 0x5a,
	0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x65, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_address_proto_rawDescOnce sync.Once
	file_proto_address_proto_rawDescData = file_proto_address_proto_rawDesc
)

func file_proto_address_proto_rawDescGZIP() []byte {
	file_proto_address_proto_rawDescOnce.Do(func() {
		file_proto_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_address_proto_rawDescData)
	})
	return file_proto_address_proto_rawDescData
}

var file_proto_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_address_proto_goTypes = []interface{}{
	(*Address)(nil), // 0: Address
}
var file_proto_address_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_address_proto_init() }
func file_proto_address_proto_init() {
	if File_proto_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_proto_address_proto_rawDesc,
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
	file_proto_address_proto_rawDesc = nil
	file_proto_address_proto_goTypes = nil
	file_proto_address_proto_depIdxs = nil
}