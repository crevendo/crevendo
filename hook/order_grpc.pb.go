// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/hook/order.proto

package hook

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderHooksClient is the client API for OrderHooks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderHooksClient interface {
	OrderCreated(ctx context.Context, in *OrderCreateHookParams, opts ...grpc.CallOption) (*OrderCreateHookParams, error)
	OrderStatuses(ctx context.Context, in *OrderStatusesHookParams, opts ...grpc.CallOption) (*OrderStatusesHookParams, error)
	OrderItemStatuses(ctx context.Context, in *OrderStatusesHookParams, opts ...grpc.CallOption) (*OrderStatusesHookParams, error)
}

type orderHooksClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderHooksClient(cc grpc.ClientConnInterface) OrderHooksClient {
	return &orderHooksClient{cc}
}

func (c *orderHooksClient) OrderCreated(ctx context.Context, in *OrderCreateHookParams, opts ...grpc.CallOption) (*OrderCreateHookParams, error) {
	out := new(OrderCreateHookParams)
	err := c.cc.Invoke(ctx, "/OrderHooks/OrderCreated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderHooksClient) OrderStatuses(ctx context.Context, in *OrderStatusesHookParams, opts ...grpc.CallOption) (*OrderStatusesHookParams, error) {
	out := new(OrderStatusesHookParams)
	err := c.cc.Invoke(ctx, "/OrderHooks/OrderStatuses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderHooksClient) OrderItemStatuses(ctx context.Context, in *OrderStatusesHookParams, opts ...grpc.CallOption) (*OrderStatusesHookParams, error) {
	out := new(OrderStatusesHookParams)
	err := c.cc.Invoke(ctx, "/OrderHooks/OrderItemStatuses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderHooksServer is the server API for OrderHooks service.
// All implementations must embed UnimplementedOrderHooksServer
// for forward compatibility
type OrderHooksServer interface {
	OrderCreated(context.Context, *OrderCreateHookParams) (*OrderCreateHookParams, error)
	OrderStatuses(context.Context, *OrderStatusesHookParams) (*OrderStatusesHookParams, error)
	OrderItemStatuses(context.Context, *OrderStatusesHookParams) (*OrderStatusesHookParams, error)
	mustEmbedUnimplementedOrderHooksServer()
}

// UnimplementedOrderHooksServer must be embedded to have forward compatible implementations.
type UnimplementedOrderHooksServer struct {
}

func (UnimplementedOrderHooksServer) OrderCreated(context.Context, *OrderCreateHookParams) (*OrderCreateHookParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderCreated not implemented")
}
func (UnimplementedOrderHooksServer) OrderStatuses(context.Context, *OrderStatusesHookParams) (*OrderStatusesHookParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderStatuses not implemented")
}
func (UnimplementedOrderHooksServer) OrderItemStatuses(context.Context, *OrderStatusesHookParams) (*OrderStatusesHookParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderItemStatuses not implemented")
}
func (UnimplementedOrderHooksServer) mustEmbedUnimplementedOrderHooksServer() {}

// UnsafeOrderHooksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderHooksServer will
// result in compilation errors.
type UnsafeOrderHooksServer interface {
	mustEmbedUnimplementedOrderHooksServer()
}

func RegisterOrderHooksServer(s grpc.ServiceRegistrar, srv OrderHooksServer) {
	s.RegisterService(&OrderHooks_ServiceDesc, srv)
}

func _OrderHooks_OrderCreated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderCreateHookParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderHooksServer).OrderCreated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderHooks/OrderCreated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderHooksServer).OrderCreated(ctx, req.(*OrderCreateHookParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderHooks_OrderStatuses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderStatusesHookParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderHooksServer).OrderStatuses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderHooks/OrderStatuses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderHooksServer).OrderStatuses(ctx, req.(*OrderStatusesHookParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderHooks_OrderItemStatuses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderStatusesHookParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderHooksServer).OrderItemStatuses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderHooks/OrderItemStatuses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderHooksServer).OrderItemStatuses(ctx, req.(*OrderStatusesHookParams))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderHooks_ServiceDesc is the grpc.ServiceDesc for OrderHooks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderHooks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderHooks",
	HandlerType: (*OrderHooksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrderCreated",
			Handler:    _OrderHooks_OrderCreated_Handler,
		},
		{
			MethodName: "OrderStatuses",
			Handler:    _OrderHooks_OrderStatuses_Handler,
		},
		{
			MethodName: "OrderItemStatuses",
			Handler:    _OrderHooks_OrderItemStatuses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hook/order.proto",
}
