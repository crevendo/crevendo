// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/message/user.proto

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

type UserGetMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserGetMessage) Reset() {
	*x = UserGetMessage{}
	mi := &file_proto_message_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserGetMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGetMessage) ProtoMessage() {}

func (x *UserGetMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGetMessage.ProtoReflect.Descriptor instead.
func (*UserGetMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserGetMessage) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserGetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *data.User             `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserGetResponse) Reset() {
	*x = UserGetResponse{}
	mi := &file_proto_message_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGetResponse) ProtoMessage() {}

func (x *UserGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGetResponse.ProtoReflect.Descriptor instead.
func (*UserGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserGetResponse) GetUser() *data.User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserListMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserListMessage) Reset() {
	*x = UserListMessage{}
	mi := &file_proto_message_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserListMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListMessage) ProtoMessage() {}

func (x *UserListMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListMessage.ProtoReflect.Descriptor instead.
func (*UserListMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_user_proto_rawDescGZIP(), []int{2}
}

type UserListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*data.User           `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserListResponse) Reset() {
	*x = UserListResponse{}
	mi := &file_proto_message_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListResponse) ProtoMessage() {}

func (x *UserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListResponse.ProtoReflect.Descriptor instead.
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserListResponse) GetUsers() []*data.User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_proto_message_user_proto protoreflect.FileDescriptor

const file_proto_message_user_proto_rawDesc = "" +
	"\n" +
	"\x18proto/message/user.proto\x1a\x10proto/user.proto\" \n" +
	"\x0eUserGetMessage\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\",\n" +
	"\x0fUserGetResponse\x12\x19\n" +
	"\x04user\x18\x01 \x01(\v2\x05.UserR\x04user\"\x11\n" +
	"\x0fUserListMessage\"/\n" +
	"\x10UserListResponse\x12\x1b\n" +
	"\x05users\x18\x01 \x03(\v2\x05.UserR\x05usersB&Z$github.com/crevendo/crevendo/messageb\x06proto3"

var (
	file_proto_message_user_proto_rawDescOnce sync.Once
	file_proto_message_user_proto_rawDescData []byte
)

func file_proto_message_user_proto_rawDescGZIP() []byte {
	file_proto_message_user_proto_rawDescOnce.Do(func() {
		file_proto_message_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_message_user_proto_rawDesc), len(file_proto_message_user_proto_rawDesc)))
	})
	return file_proto_message_user_proto_rawDescData
}

var file_proto_message_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_message_user_proto_goTypes = []any{
	(*UserGetMessage)(nil),   // 0: UserGetMessage
	(*UserGetResponse)(nil),  // 1: UserGetResponse
	(*UserListMessage)(nil),  // 2: UserListMessage
	(*UserListResponse)(nil), // 3: UserListResponse
	(*data.User)(nil),        // 4: User
}
var file_proto_message_user_proto_depIdxs = []int32{
	4, // 0: UserGetResponse.user:type_name -> User
	4, // 1: UserListResponse.users:type_name -> User
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_message_user_proto_init() }
func file_proto_message_user_proto_init() {
	if File_proto_message_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_message_user_proto_rawDesc), len(file_proto_message_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_message_user_proto_goTypes,
		DependencyIndexes: file_proto_message_user_proto_depIdxs,
		MessageInfos:      file_proto_message_user_proto_msgTypes,
	}.Build()
	File_proto_message_user_proto = out.File
	file_proto_message_user_proto_goTypes = nil
	file_proto_message_user_proto_depIdxs = nil
}
