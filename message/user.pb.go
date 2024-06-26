// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: proto/message/user.proto

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

type UserGetMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserGetMessage) Reset() {
	*x = UserGetMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGetMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGetMessage) ProtoMessage() {}

func (x *UserGetMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *data.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserGetResponse) Reset() {
	*x = UserGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGetResponse) ProtoMessage() {}

func (x *UserGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserListMessage) Reset() {
	*x = UserListMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListMessage) ProtoMessage() {}

func (x *UserListMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*data.User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UserListResponse) Reset() {
	*x = UserListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListResponse) ProtoMessage() {}

func (x *UserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_proto_message_user_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c,
	0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x11, 0x0a, 0x0f,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x2f, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x72, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x63, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_message_user_proto_rawDescOnce sync.Once
	file_proto_message_user_proto_rawDescData = file_proto_message_user_proto_rawDesc
)

func file_proto_message_user_proto_rawDescGZIP() []byte {
	file_proto_message_user_proto_rawDescOnce.Do(func() {
		file_proto_message_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_user_proto_rawDescData)
	})
	return file_proto_message_user_proto_rawDescData
}

var file_proto_message_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_message_user_proto_goTypes = []interface{}{
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
	if !protoimpl.UnsafeEnabled {
		file_proto_message_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGetMessage); i {
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
		file_proto_message_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGetResponse); i {
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
		file_proto_message_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListMessage); i {
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
		file_proto_message_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListResponse); i {
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
			RawDescriptor: file_proto_message_user_proto_rawDesc,
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
	file_proto_message_user_proto_rawDesc = nil
	file_proto_message_user_proto_goTypes = nil
	file_proto_message_user_proto_depIdxs = nil
}
