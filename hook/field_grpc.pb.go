// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/hook/field.proto

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

// FieldHooksClient is the client API for FieldHooks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FieldHooksClient interface {
	ListAddressField(ctx context.Context, in *ListFieldHookParams, opts ...grpc.CallOption) (*ListFieldHookParams, error)
	ListUserField(ctx context.Context, in *ListFieldHookParams, opts ...grpc.CallOption) (*ListFieldHookParams, error)
}

type fieldHooksClient struct {
	cc grpc.ClientConnInterface
}

func NewFieldHooksClient(cc grpc.ClientConnInterface) FieldHooksClient {
	return &fieldHooksClient{cc}
}

func (c *fieldHooksClient) ListAddressField(ctx context.Context, in *ListFieldHookParams, opts ...grpc.CallOption) (*ListFieldHookParams, error) {
	out := new(ListFieldHookParams)
	err := c.cc.Invoke(ctx, "/FieldHooks/ListAddressField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fieldHooksClient) ListUserField(ctx context.Context, in *ListFieldHookParams, opts ...grpc.CallOption) (*ListFieldHookParams, error) {
	out := new(ListFieldHookParams)
	err := c.cc.Invoke(ctx, "/FieldHooks/ListUserField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FieldHooksServer is the server API for FieldHooks service.
// All implementations must embed UnimplementedFieldHooksServer
// for forward compatibility
type FieldHooksServer interface {
	ListAddressField(context.Context, *ListFieldHookParams) (*ListFieldHookParams, error)
	ListUserField(context.Context, *ListFieldHookParams) (*ListFieldHookParams, error)
	mustEmbedUnimplementedFieldHooksServer()
}

// UnimplementedFieldHooksServer must be embedded to have forward compatible implementations.
type UnimplementedFieldHooksServer struct {
}

func (UnimplementedFieldHooksServer) ListAddressField(context.Context, *ListFieldHookParams) (*ListFieldHookParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAddressField not implemented")
}
func (UnimplementedFieldHooksServer) ListUserField(context.Context, *ListFieldHookParams) (*ListFieldHookParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserField not implemented")
}
func (UnimplementedFieldHooksServer) mustEmbedUnimplementedFieldHooksServer() {}

// UnsafeFieldHooksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FieldHooksServer will
// result in compilation errors.
type UnsafeFieldHooksServer interface {
	mustEmbedUnimplementedFieldHooksServer()
}

func RegisterFieldHooksServer(s grpc.ServiceRegistrar, srv FieldHooksServer) {
	s.RegisterService(&FieldHooks_ServiceDesc, srv)
}

func _FieldHooks_ListAddressField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFieldHookParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FieldHooksServer).ListAddressField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FieldHooks/ListAddressField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FieldHooksServer).ListAddressField(ctx, req.(*ListFieldHookParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _FieldHooks_ListUserField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFieldHookParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FieldHooksServer).ListUserField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FieldHooks/ListUserField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FieldHooksServer).ListUserField(ctx, req.(*ListFieldHookParams))
	}
	return interceptor(ctx, in, info, handler)
}

// FieldHooks_ServiceDesc is the grpc.ServiceDesc for FieldHooks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FieldHooks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FieldHooks",
	HandlerType: (*FieldHooksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAddressField",
			Handler:    _FieldHooks_ListAddressField_Handler,
		},
		{
			MethodName: "ListUserField",
			Handler:    _FieldHooks_ListUserField_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hook/field.proto",
}
