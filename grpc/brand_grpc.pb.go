// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/service/brand.proto

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

// BrandServiceClient is the client API for BrandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrandServiceClient interface {
	List(ctx context.Context, in *message.BrandListMessage, opts ...grpc.CallOption) (*message.BrandListResponse, error)
}

type brandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBrandServiceClient(cc grpc.ClientConnInterface) BrandServiceClient {
	return &brandServiceClient{cc}
}

func (c *brandServiceClient) List(ctx context.Context, in *message.BrandListMessage, opts ...grpc.CallOption) (*message.BrandListResponse, error) {
	out := new(message.BrandListResponse)
	err := c.cc.Invoke(ctx, "/BrandService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrandServiceServer is the server API for BrandService service.
// All implementations must embed UnimplementedBrandServiceServer
// for forward compatibility
type BrandServiceServer interface {
	List(context.Context, *message.BrandListMessage) (*message.BrandListResponse, error)
	mustEmbedUnimplementedBrandServiceServer()
}

// UnimplementedBrandServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBrandServiceServer struct {
}

func (UnimplementedBrandServiceServer) List(context.Context, *message.BrandListMessage) (*message.BrandListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedBrandServiceServer) mustEmbedUnimplementedBrandServiceServer() {}

// UnsafeBrandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrandServiceServer will
// result in compilation errors.
type UnsafeBrandServiceServer interface {
	mustEmbedUnimplementedBrandServiceServer()
}

func RegisterBrandServiceServer(s grpc.ServiceRegistrar, srv BrandServiceServer) {
	s.RegisterService(&BrandService_ServiceDesc, srv)
}

func _BrandService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.BrandListMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BrandService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).List(ctx, req.(*message.BrandListMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// BrandService_ServiceDesc is the grpc.ServiceDesc for BrandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BrandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BrandService",
	HandlerType: (*BrandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _BrandService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service/brand.proto",
}