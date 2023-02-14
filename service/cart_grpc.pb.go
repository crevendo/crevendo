// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/service/cart.proto

package service

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

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartServiceClient interface {
	Get(ctx context.Context, in *message.CartGetMessage, opts ...grpc.CallOption) (*message.CartGetResponse, error)
	Create(ctx context.Context, in *message.CartCreateMessage, opts ...grpc.CallOption) (*message.CartCreateResponse, error)
	AddItem(ctx context.Context, in *message.AddItemMessage, opts ...grpc.CallOption) (*message.AddItemResponse, error)
	RemoveItem(ctx context.Context, in *message.RemoveItemMessage, opts ...grpc.CallOption) (*message.RemoveItemResponse, error)
	UpdateItemQuantity(ctx context.Context, in *message.UpdateItemQuantityMessage, opts ...grpc.CallOption) (*message.UpdateItemQuantityResponse, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) Get(ctx context.Context, in *message.CartGetMessage, opts ...grpc.CallOption) (*message.CartGetResponse, error) {
	out := new(message.CartGetResponse)
	err := c.cc.Invoke(ctx, "/CartService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) Create(ctx context.Context, in *message.CartCreateMessage, opts ...grpc.CallOption) (*message.CartCreateResponse, error) {
	out := new(message.CartCreateResponse)
	err := c.cc.Invoke(ctx, "/CartService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) AddItem(ctx context.Context, in *message.AddItemMessage, opts ...grpc.CallOption) (*message.AddItemResponse, error) {
	out := new(message.AddItemResponse)
	err := c.cc.Invoke(ctx, "/CartService/AddItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) RemoveItem(ctx context.Context, in *message.RemoveItemMessage, opts ...grpc.CallOption) (*message.RemoveItemResponse, error) {
	out := new(message.RemoveItemResponse)
	err := c.cc.Invoke(ctx, "/CartService/RemoveItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) UpdateItemQuantity(ctx context.Context, in *message.UpdateItemQuantityMessage, opts ...grpc.CallOption) (*message.UpdateItemQuantityResponse, error) {
	out := new(message.UpdateItemQuantityResponse)
	err := c.cc.Invoke(ctx, "/CartService/UpdateItemQuantity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations must embed UnimplementedCartServiceServer
// for forward compatibility
type CartServiceServer interface {
	Get(context.Context, *message.CartGetMessage) (*message.CartGetResponse, error)
	Create(context.Context, *message.CartCreateMessage) (*message.CartCreateResponse, error)
	AddItem(context.Context, *message.AddItemMessage) (*message.AddItemResponse, error)
	RemoveItem(context.Context, *message.RemoveItemMessage) (*message.RemoveItemResponse, error)
	UpdateItemQuantity(context.Context, *message.UpdateItemQuantityMessage) (*message.UpdateItemQuantityResponse, error)
	mustEmbedUnimplementedCartServiceServer()
}

// UnimplementedCartServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (UnimplementedCartServiceServer) Get(context.Context, *message.CartGetMessage) (*message.CartGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCartServiceServer) Create(context.Context, *message.CartCreateMessage) (*message.CartCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCartServiceServer) AddItem(context.Context, *message.AddItemMessage) (*message.AddItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (UnimplementedCartServiceServer) RemoveItem(context.Context, *message.RemoveItemMessage) (*message.RemoveItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveItem not implemented")
}
func (UnimplementedCartServiceServer) UpdateItemQuantity(context.Context, *message.UpdateItemQuantityMessage) (*message.UpdateItemQuantityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItemQuantity not implemented")
}
func (UnimplementedCartServiceServer) mustEmbedUnimplementedCartServiceServer() {}

// UnsafeCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServiceServer will
// result in compilation errors.
type UnsafeCartServiceServer interface {
	mustEmbedUnimplementedCartServiceServer()
}

func RegisterCartServiceServer(s grpc.ServiceRegistrar, srv CartServiceServer) {
	s.RegisterService(&CartService_ServiceDesc, srv)
}

func _CartService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.CartGetMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).Get(ctx, req.(*message.CartGetMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.CartCreateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).Create(ctx, req.(*message.CartCreateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.AddItemMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/AddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddItem(ctx, req.(*message.AddItemMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_RemoveItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.RemoveItemMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).RemoveItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/RemoveItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).RemoveItem(ctx, req.(*message.RemoveItemMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_UpdateItemQuantity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.UpdateItemQuantityMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).UpdateItemQuantity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/UpdateItemQuantity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).UpdateItemQuantity(ctx, req.(*message.UpdateItemQuantityMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// CartService_ServiceDesc is the grpc.ServiceDesc for CartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CartService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CartService_Create_Handler,
		},
		{
			MethodName: "AddItem",
			Handler:    _CartService_AddItem_Handler,
		},
		{
			MethodName: "RemoveItem",
			Handler:    _CartService_RemoveItem_Handler,
		},
		{
			MethodName: "UpdateItemQuantity",
			Handler:    _CartService_UpdateItemQuantity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service/cart.proto",
}
