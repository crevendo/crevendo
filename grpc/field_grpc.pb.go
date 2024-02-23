// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: proto/service/field.proto

package grpc

import (
	context "context"
	message "github.com/crevendo/crevendo/message"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FieldServiceClient is the client API for FieldService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FieldServiceClient interface {
	UserFieldList(ctx context.Context, in *message.FieldListMessage, opts ...grpc.CallOption) (*message.FieldListResponse, error)
	AddressFieldList(ctx context.Context, in *message.FieldListMessage, opts ...grpc.CallOption) (*message.FieldListResponse, error)
}

type fieldServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFieldServiceClient(cc grpc.ClientConnInterface) FieldServiceClient {
	return &fieldServiceClient{cc}
}

func (c *fieldServiceClient) UserFieldList(ctx context.Context, in *message.FieldListMessage, opts ...grpc.CallOption) (*message.FieldListResponse, error) {
	out := new(message.FieldListResponse)
	err := c.cc.Invoke(ctx, "/FieldService/UserFieldList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fieldServiceClient) AddressFieldList(ctx context.Context, in *message.FieldListMessage, opts ...grpc.CallOption) (*message.FieldListResponse, error) {
	out := new(message.FieldListResponse)
	err := c.cc.Invoke(ctx, "/FieldService/AddressFieldList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FieldServiceServer is the server API for FieldService service.
// All implementations must embed UnimplementedFieldServiceServer
// for forward compatibility
type FieldServiceServer interface {
	UserFieldList(context.Context, *message.FieldListMessage) (*message.FieldListResponse, error)
	AddressFieldList(context.Context, *message.FieldListMessage) (*message.FieldListResponse, error)
	mustEmbedUnimplementedFieldServiceServer()
}

// UnimplementedFieldServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFieldServiceServer struct {
}

func (UnimplementedFieldServiceServer) UserFieldList(context.Context, *message.FieldListMessage) (*message.FieldListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserFieldList not implemented")
}
func (UnimplementedFieldServiceServer) AddressFieldList(context.Context, *message.FieldListMessage) (*message.FieldListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressFieldList not implemented")
}
func (UnimplementedFieldServiceServer) mustEmbedUnimplementedFieldServiceServer() {}

// UnsafeFieldServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FieldServiceServer will
// result in compilation errors.
type UnsafeFieldServiceServer interface {
	mustEmbedUnimplementedFieldServiceServer()
}

func RegisterFieldServiceServer(s grpc.ServiceRegistrar, srv FieldServiceServer) {
	s.RegisterService(&FieldService_ServiceDesc, srv)
}

func _FieldService_UserFieldList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.FieldListMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FieldServiceServer).UserFieldList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FieldService/UserFieldList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FieldServiceServer).UserFieldList(ctx, req.(*message.FieldListMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _FieldService_AddressFieldList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.FieldListMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FieldServiceServer).AddressFieldList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FieldService/AddressFieldList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FieldServiceServer).AddressFieldList(ctx, req.(*message.FieldListMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// FieldService_ServiceDesc is the grpc.ServiceDesc for FieldService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FieldService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FieldService",
	HandlerType: (*FieldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserFieldList",
			Handler:    _FieldService_UserFieldList_Handler,
		},
		{
			MethodName: "AddressFieldList",
			Handler:    _FieldService_AddressFieldList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service/field.proto",
}
