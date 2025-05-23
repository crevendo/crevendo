// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/message/address.proto

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

type AddressGetMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressGetMessage) Reset() {
	*x = AddressGetMessage{}
	mi := &file_proto_message_address_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressGetMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressGetMessage) ProtoMessage() {}

func (x *AddressGetMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_address_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressGetMessage.ProtoReflect.Descriptor instead.
func (*AddressGetMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_address_proto_rawDescGZIP(), []int{0}
}

func (x *AddressGetMessage) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AddressGetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       *data.Address          `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressGetResponse) Reset() {
	*x = AddressGetResponse{}
	mi := &file_proto_message_address_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressGetResponse) ProtoMessage() {}

func (x *AddressGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_address_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressGetResponse.ProtoReflect.Descriptor instead.
func (*AddressGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_address_proto_rawDescGZIP(), []int{1}
}

func (x *AddressGetResponse) GetAddress() *data.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type AddressDeleteMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressDeleteMessage) Reset() {
	*x = AddressDeleteMessage{}
	mi := &file_proto_message_address_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressDeleteMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressDeleteMessage) ProtoMessage() {}

func (x *AddressDeleteMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_address_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressDeleteMessage.ProtoReflect.Descriptor instead.
func (*AddressDeleteMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_address_proto_rawDescGZIP(), []int{2}
}

func (x *AddressDeleteMessage) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AddressDeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressDeleteResponse) Reset() {
	*x = AddressDeleteResponse{}
	mi := &file_proto_message_address_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressDeleteResponse) ProtoMessage() {}

func (x *AddressDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_address_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressDeleteResponse.ProtoReflect.Descriptor instead.
func (*AddressDeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_address_proto_rawDescGZIP(), []int{3}
}

type AddressListMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressListMessage) Reset() {
	*x = AddressListMessage{}
	mi := &file_proto_message_address_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressListMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressListMessage) ProtoMessage() {}

func (x *AddressListMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_address_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressListMessage.ProtoReflect.Descriptor instead.
func (*AddressListMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_address_proto_rawDescGZIP(), []int{4}
}

func (x *AddressListMessage) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AddressListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       []*data.Address        `protobuf:"bytes,1,rep,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressListResponse) Reset() {
	*x = AddressListResponse{}
	mi := &file_proto_message_address_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressListResponse) ProtoMessage() {}

func (x *AddressListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_address_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressListResponse.ProtoReflect.Descriptor instead.
func (*AddressListResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_address_proto_rawDescGZIP(), []int{5}
}

func (x *AddressListResponse) GetAddress() []*data.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type AddressCreateMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Name          string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Address       string                 `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	PostalCode    uint32                 `protobuf:"varint,6,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
	Phone         string                 `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Notes         string                 `protobuf:"bytes,8,opt,name=notes,proto3" json:"notes,omitempty"`
	Fields        []*data.Field          `protobuf:"bytes,9,rep,name=fields,proto3" json:"fields,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressCreateMessage) Reset() {
	*x = AddressCreateMessage{}
	mi := &file_proto_message_address_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressCreateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressCreateMessage) ProtoMessage() {}

func (x *AddressCreateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_address_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressCreateMessage.ProtoReflect.Descriptor instead.
func (*AddressCreateMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_address_proto_rawDescGZIP(), []int{6}
}

func (x *AddressCreateMessage) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddressCreateMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddressCreateMessage) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddressCreateMessage) GetPostalCode() uint32 {
	if x != nil {
		return x.PostalCode
	}
	return 0
}

func (x *AddressCreateMessage) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AddressCreateMessage) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *AddressCreateMessage) GetFields() []*data.Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

type AddressCreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressCreateResponse) Reset() {
	*x = AddressCreateResponse{}
	mi := &file_proto_message_address_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressCreateResponse) ProtoMessage() {}

func (x *AddressCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_address_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressCreateResponse.ProtoReflect.Descriptor instead.
func (*AddressCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_address_proto_rawDescGZIP(), []int{7}
}

type AddressUpdateMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Address       string                 `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	PostalCode    uint32                 `protobuf:"varint,6,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
	Phone         string                 `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Notes         string                 `protobuf:"bytes,8,opt,name=notes,proto3" json:"notes,omitempty"`
	Fields        []*data.Field          `protobuf:"bytes,9,rep,name=fields,proto3" json:"fields,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressUpdateMessage) Reset() {
	*x = AddressUpdateMessage{}
	mi := &file_proto_message_address_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressUpdateMessage) ProtoMessage() {}

func (x *AddressUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_address_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressUpdateMessage.ProtoReflect.Descriptor instead.
func (*AddressUpdateMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_address_proto_rawDescGZIP(), []int{8}
}

func (x *AddressUpdateMessage) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddressUpdateMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddressUpdateMessage) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddressUpdateMessage) GetPostalCode() uint32 {
	if x != nil {
		return x.PostalCode
	}
	return 0
}

func (x *AddressUpdateMessage) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AddressUpdateMessage) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *AddressUpdateMessage) GetFields() []*data.Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

type AddressUpdateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       *data.Address          `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressUpdateResponse) Reset() {
	*x = AddressUpdateResponse{}
	mi := &file_proto_message_address_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressUpdateResponse) ProtoMessage() {}

func (x *AddressUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_address_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressUpdateResponse.ProtoReflect.Descriptor instead.
func (*AddressUpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_address_proto_rawDescGZIP(), []int{9}
}

func (x *AddressUpdateResponse) GetAddress() *data.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

var File_proto_message_address_proto protoreflect.FileDescriptor

const file_proto_message_address_proto_rawDesc = "" +
	"\n" +
	"\x1bproto/message/address.proto\x1a\x13proto/address.proto\x1a\x11proto/field.proto\"#\n" +
	"\x11AddressGetMessage\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\"8\n" +
	"\x12AddressGetResponse\x12\"\n" +
	"\aaddress\x18\x01 \x01(\v2\b.AddressR\aaddress\"&\n" +
	"\x14AddressDeleteMessage\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\"\x17\n" +
	"\x15AddressDeleteResponse\",\n" +
	"\x12AddressListMessage\x12\x16\n" +
	"\x06userId\x18\x01 \x01(\rR\x06userId\"9\n" +
	"\x13AddressListResponse\x12\"\n" +
	"\aaddress\x18\x01 \x03(\v2\b.AddressR\aaddress\"\xc8\x01\n" +
	"\x14AddressCreateMessage\x12\x16\n" +
	"\x06userId\x18\x02 \x01(\rR\x06userId\x12\x12\n" +
	"\x04name\x18\x04 \x01(\tR\x04name\x12\x18\n" +
	"\aaddress\x18\x05 \x01(\tR\aaddress\x12\x1e\n" +
	"\n" +
	"postalCode\x18\x06 \x01(\rR\n" +
	"postalCode\x12\x14\n" +
	"\x05phone\x18\a \x01(\tR\x05phone\x12\x14\n" +
	"\x05notes\x18\b \x01(\tR\x05notes\x12\x1e\n" +
	"\x06fields\x18\t \x03(\v2\x06.FieldR\x06fields\"\x17\n" +
	"\x15AddressCreateResponse\"\xc0\x01\n" +
	"\x14AddressUpdateMessage\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x12\n" +
	"\x04name\x18\x04 \x01(\tR\x04name\x12\x18\n" +
	"\aaddress\x18\x05 \x01(\tR\aaddress\x12\x1e\n" +
	"\n" +
	"postalCode\x18\x06 \x01(\rR\n" +
	"postalCode\x12\x14\n" +
	"\x05phone\x18\a \x01(\tR\x05phone\x12\x14\n" +
	"\x05notes\x18\b \x01(\tR\x05notes\x12\x1e\n" +
	"\x06fields\x18\t \x03(\v2\x06.FieldR\x06fields\";\n" +
	"\x15AddressUpdateResponse\x12\"\n" +
	"\aaddress\x18\x01 \x01(\v2\b.AddressR\aaddressB&Z$github.com/crevendo/crevendo/messageb\x06proto3"

var (
	file_proto_message_address_proto_rawDescOnce sync.Once
	file_proto_message_address_proto_rawDescData []byte
)

func file_proto_message_address_proto_rawDescGZIP() []byte {
	file_proto_message_address_proto_rawDescOnce.Do(func() {
		file_proto_message_address_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_message_address_proto_rawDesc), len(file_proto_message_address_proto_rawDesc)))
	})
	return file_proto_message_address_proto_rawDescData
}

var file_proto_message_address_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_message_address_proto_goTypes = []any{
	(*AddressGetMessage)(nil),     // 0: AddressGetMessage
	(*AddressGetResponse)(nil),    // 1: AddressGetResponse
	(*AddressDeleteMessage)(nil),  // 2: AddressDeleteMessage
	(*AddressDeleteResponse)(nil), // 3: AddressDeleteResponse
	(*AddressListMessage)(nil),    // 4: AddressListMessage
	(*AddressListResponse)(nil),   // 5: AddressListResponse
	(*AddressCreateMessage)(nil),  // 6: AddressCreateMessage
	(*AddressCreateResponse)(nil), // 7: AddressCreateResponse
	(*AddressUpdateMessage)(nil),  // 8: AddressUpdateMessage
	(*AddressUpdateResponse)(nil), // 9: AddressUpdateResponse
	(*data.Address)(nil),          // 10: Address
	(*data.Field)(nil),            // 11: Field
}
var file_proto_message_address_proto_depIdxs = []int32{
	10, // 0: AddressGetResponse.address:type_name -> Address
	10, // 1: AddressListResponse.address:type_name -> Address
	11, // 2: AddressCreateMessage.fields:type_name -> Field
	11, // 3: AddressUpdateMessage.fields:type_name -> Field
	10, // 4: AddressUpdateResponse.address:type_name -> Address
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_message_address_proto_init() }
func file_proto_message_address_proto_init() {
	if File_proto_message_address_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_message_address_proto_rawDesc), len(file_proto_message_address_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_message_address_proto_goTypes,
		DependencyIndexes: file_proto_message_address_proto_depIdxs,
		MessageInfos:      file_proto_message_address_proto_msgTypes,
	}.Build()
	File_proto_message_address_proto = out.File
	file_proto_message_address_proto_goTypes = nil
	file_proto_message_address_proto_depIdxs = nil
}
