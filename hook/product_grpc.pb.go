// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/hook/product.proto

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

// ProductHooksClient is the client API for ProductHooks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductHooksClient interface {
	ProductList(ctx context.Context, in *ProductListHookParams, opts ...grpc.CallOption) (*ProductListHookParams, error)
	ProductGet(ctx context.Context, in *ProductGetHookParams, opts ...grpc.CallOption) (*ProductGetHookParams, error)
}

type productHooksClient struct {
	cc grpc.ClientConnInterface
}

func NewProductHooksClient(cc grpc.ClientConnInterface) ProductHooksClient {
	return &productHooksClient{cc}
}

func (c *productHooksClient) ProductList(ctx context.Context, in *ProductListHookParams, opts ...grpc.CallOption) (*ProductListHookParams, error) {
	out := new(ProductListHookParams)
	err := c.cc.Invoke(ctx, "/ProductHooks/ProductList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productHooksClient) ProductGet(ctx context.Context, in *ProductGetHookParams, opts ...grpc.CallOption) (*ProductGetHookParams, error) {
	out := new(ProductGetHookParams)
	err := c.cc.Invoke(ctx, "/ProductHooks/ProductGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductHooksServer is the server API for ProductHooks service.
// All implementations must embed UnimplementedProductHooksServer
// for forward compatibility
type ProductHooksServer interface {
	ProductList(context.Context, *ProductListHookParams) (*ProductListHookParams, error)
	ProductGet(context.Context, *ProductGetHookParams) (*ProductGetHookParams, error)
	mustEmbedUnimplementedProductHooksServer()
}

// UnimplementedProductHooksServer must be embedded to have forward compatible implementations.
type UnimplementedProductHooksServer struct {
}

func (UnimplementedProductHooksServer) ProductList(context.Context, *ProductListHookParams) (*ProductListHookParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductList not implemented")
}
func (UnimplementedProductHooksServer) ProductGet(context.Context, *ProductGetHookParams) (*ProductGetHookParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductGet not implemented")
}
func (UnimplementedProductHooksServer) mustEmbedUnimplementedProductHooksServer() {}

// UnsafeProductHooksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductHooksServer will
// result in compilation errors.
type UnsafeProductHooksServer interface {
	mustEmbedUnimplementedProductHooksServer()
}

func RegisterProductHooksServer(s grpc.ServiceRegistrar, srv ProductHooksServer) {
	s.RegisterService(&ProductHooks_ServiceDesc, srv)
}

func _ProductHooks_ProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductListHookParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductHooksServer).ProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductHooks/ProductList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductHooksServer).ProductList(ctx, req.(*ProductListHookParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductHooks_ProductGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductGetHookParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductHooksServer).ProductGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductHooks/ProductGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductHooksServer).ProductGet(ctx, req.(*ProductGetHookParams))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductHooks_ServiceDesc is the grpc.ServiceDesc for ProductHooks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductHooks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProductHooks",
	HandlerType: (*ProductHooksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProductList",
			Handler:    _ProductHooks_ProductList_Handler,
		},
		{
			MethodName: "ProductGet",
			Handler:    _ProductHooks_ProductGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hook/product.proto",
}
