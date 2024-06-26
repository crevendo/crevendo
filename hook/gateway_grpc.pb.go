// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: proto/hook/gateway.proto

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

// GatewayHooksClient is the client API for GatewayHooks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayHooksClient interface {
	List(ctx context.Context, in *GatewayListHookParams, opts ...grpc.CallOption) (*GatewayListHookParams, error)
}

type gatewayHooksClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayHooksClient(cc grpc.ClientConnInterface) GatewayHooksClient {
	return &gatewayHooksClient{cc}
}

func (c *gatewayHooksClient) List(ctx context.Context, in *GatewayListHookParams, opts ...grpc.CallOption) (*GatewayListHookParams, error) {
	out := new(GatewayListHookParams)
	err := c.cc.Invoke(ctx, "/GatewayHooks/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayHooksServer is the server API for GatewayHooks service.
// All implementations must embed UnimplementedGatewayHooksServer
// for forward compatibility
type GatewayHooksServer interface {
	List(context.Context, *GatewayListHookParams) (*GatewayListHookParams, error)
	mustEmbedUnimplementedGatewayHooksServer()
}

// UnimplementedGatewayHooksServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayHooksServer struct {
}

func (UnimplementedGatewayHooksServer) List(context.Context, *GatewayListHookParams) (*GatewayListHookParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedGatewayHooksServer) mustEmbedUnimplementedGatewayHooksServer() {}

// UnsafeGatewayHooksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayHooksServer will
// result in compilation errors.
type UnsafeGatewayHooksServer interface {
	mustEmbedUnimplementedGatewayHooksServer()
}

func RegisterGatewayHooksServer(s grpc.ServiceRegistrar, srv GatewayHooksServer) {
	s.RegisterService(&GatewayHooks_ServiceDesc, srv)
}

func _GatewayHooks_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayListHookParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayHooksServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GatewayHooks/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayHooksServer).List(ctx, req.(*GatewayListHookParams))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayHooks_ServiceDesc is the grpc.ServiceDesc for GatewayHooks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayHooks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GatewayHooks",
	HandlerType: (*GatewayHooksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _GatewayHooks_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hook/gateway.proto",
}
