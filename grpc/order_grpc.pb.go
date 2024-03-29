// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: proto/service/order.proto

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

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	Get(ctx context.Context, in *message.OrderGetMessage, opts ...grpc.CallOption) (*message.OrderGetResponse, error)
	Create(ctx context.Context, in *message.OrderCreateMessage, opts ...grpc.CallOption) (*message.OrderCreateResponse, error)
	List(ctx context.Context, in *message.OrderListMessage, opts ...grpc.CallOption) (*message.OrderListResponse, error)
	Update(ctx context.Context, in *message.OrderUpdateMessage, opts ...grpc.CallOption) (*message.OrderUpdateResponse, error)
	StatusList(ctx context.Context, in *message.OrderStatusListMessage, opts ...grpc.CallOption) (*message.OrderStatusListResponse, error)
	OrderItemStatusList(ctx context.Context, in *message.OrderItemStatusListMessage, opts ...grpc.CallOption) (*message.OrderItemStatusListResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) Get(ctx context.Context, in *message.OrderGetMessage, opts ...grpc.CallOption) (*message.OrderGetResponse, error) {
	out := new(message.OrderGetResponse)
	err := c.cc.Invoke(ctx, "/OrderService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) Create(ctx context.Context, in *message.OrderCreateMessage, opts ...grpc.CallOption) (*message.OrderCreateResponse, error) {
	out := new(message.OrderCreateResponse)
	err := c.cc.Invoke(ctx, "/OrderService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) List(ctx context.Context, in *message.OrderListMessage, opts ...grpc.CallOption) (*message.OrderListResponse, error) {
	out := new(message.OrderListResponse)
	err := c.cc.Invoke(ctx, "/OrderService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) Update(ctx context.Context, in *message.OrderUpdateMessage, opts ...grpc.CallOption) (*message.OrderUpdateResponse, error) {
	out := new(message.OrderUpdateResponse)
	err := c.cc.Invoke(ctx, "/OrderService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) StatusList(ctx context.Context, in *message.OrderStatusListMessage, opts ...grpc.CallOption) (*message.OrderStatusListResponse, error) {
	out := new(message.OrderStatusListResponse)
	err := c.cc.Invoke(ctx, "/OrderService/StatusList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) OrderItemStatusList(ctx context.Context, in *message.OrderItemStatusListMessage, opts ...grpc.CallOption) (*message.OrderItemStatusListResponse, error) {
	out := new(message.OrderItemStatusListResponse)
	err := c.cc.Invoke(ctx, "/OrderService/OrderItemStatusList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	Get(context.Context, *message.OrderGetMessage) (*message.OrderGetResponse, error)
	Create(context.Context, *message.OrderCreateMessage) (*message.OrderCreateResponse, error)
	List(context.Context, *message.OrderListMessage) (*message.OrderListResponse, error)
	Update(context.Context, *message.OrderUpdateMessage) (*message.OrderUpdateResponse, error)
	StatusList(context.Context, *message.OrderStatusListMessage) (*message.OrderStatusListResponse, error)
	OrderItemStatusList(context.Context, *message.OrderItemStatusListMessage) (*message.OrderItemStatusListResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) Get(context.Context, *message.OrderGetMessage) (*message.OrderGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedOrderServiceServer) Create(context.Context, *message.OrderCreateMessage) (*message.OrderCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrderServiceServer) List(context.Context, *message.OrderListMessage) (*message.OrderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedOrderServiceServer) Update(context.Context, *message.OrderUpdateMessage) (*message.OrderUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOrderServiceServer) StatusList(context.Context, *message.OrderStatusListMessage) (*message.OrderStatusListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatusList not implemented")
}
func (UnimplementedOrderServiceServer) OrderItemStatusList(context.Context, *message.OrderItemStatusListMessage) (*message.OrderItemStatusListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderItemStatusList not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.OrderGetMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Get(ctx, req.(*message.OrderGetMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.OrderCreateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Create(ctx, req.(*message.OrderCreateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.OrderListMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).List(ctx, req.(*message.OrderListMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.OrderUpdateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Update(ctx, req.(*message.OrderUpdateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_StatusList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.OrderStatusListMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).StatusList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/StatusList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).StatusList(ctx, req.(*message.OrderStatusListMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_OrderItemStatusList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.OrderItemStatusListMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).OrderItemStatusList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/OrderItemStatusList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).OrderItemStatusList(ctx, req.(*message.OrderItemStatusListMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _OrderService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _OrderService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _OrderService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OrderService_Update_Handler,
		},
		{
			MethodName: "StatusList",
			Handler:    _OrderService_StatusList_Handler,
		},
		{
			MethodName: "OrderItemStatusList",
			Handler:    _OrderService_OrderItemStatusList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service/order.proto",
}
