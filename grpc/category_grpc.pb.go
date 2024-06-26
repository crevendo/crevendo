// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: proto/service/category.proto

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

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryServiceClient interface {
	Get(ctx context.Context, in *message.CategoryGetMessage, opts ...grpc.CallOption) (*message.CategoryGetResponse, error)
	List(ctx context.Context, in *message.CategoryListMessage, opts ...grpc.CallOption) (*message.CategoryListResponse, error)
	Create(ctx context.Context, in *message.CategoryCreateMessage, opts ...grpc.CallOption) (*message.CategoryCreateResponse, error)
	Update(ctx context.Context, in *message.CategoryUpdateMessage, opts ...grpc.CallOption) (*message.CategoryUpdateResponse, error)
	Delete(ctx context.Context, in *message.CategoryDeleteMessage, opts ...grpc.CallOption) (*message.CategoryDeleteResponse, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) Get(ctx context.Context, in *message.CategoryGetMessage, opts ...grpc.CallOption) (*message.CategoryGetResponse, error) {
	out := new(message.CategoryGetResponse)
	err := c.cc.Invoke(ctx, "/CategoryService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) List(ctx context.Context, in *message.CategoryListMessage, opts ...grpc.CallOption) (*message.CategoryListResponse, error) {
	out := new(message.CategoryListResponse)
	err := c.cc.Invoke(ctx, "/CategoryService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) Create(ctx context.Context, in *message.CategoryCreateMessage, opts ...grpc.CallOption) (*message.CategoryCreateResponse, error) {
	out := new(message.CategoryCreateResponse)
	err := c.cc.Invoke(ctx, "/CategoryService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) Update(ctx context.Context, in *message.CategoryUpdateMessage, opts ...grpc.CallOption) (*message.CategoryUpdateResponse, error) {
	out := new(message.CategoryUpdateResponse)
	err := c.cc.Invoke(ctx, "/CategoryService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) Delete(ctx context.Context, in *message.CategoryDeleteMessage, opts ...grpc.CallOption) (*message.CategoryDeleteResponse, error) {
	out := new(message.CategoryDeleteResponse)
	err := c.cc.Invoke(ctx, "/CategoryService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServiceServer is the server API for CategoryService service.
// All implementations must embed UnimplementedCategoryServiceServer
// for forward compatibility
type CategoryServiceServer interface {
	Get(context.Context, *message.CategoryGetMessage) (*message.CategoryGetResponse, error)
	List(context.Context, *message.CategoryListMessage) (*message.CategoryListResponse, error)
	Create(context.Context, *message.CategoryCreateMessage) (*message.CategoryCreateResponse, error)
	Update(context.Context, *message.CategoryUpdateMessage) (*message.CategoryUpdateResponse, error)
	Delete(context.Context, *message.CategoryDeleteMessage) (*message.CategoryDeleteResponse, error)
	mustEmbedUnimplementedCategoryServiceServer()
}

// UnimplementedCategoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryServiceServer struct {
}

func (UnimplementedCategoryServiceServer) Get(context.Context, *message.CategoryGetMessage) (*message.CategoryGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCategoryServiceServer) List(context.Context, *message.CategoryListMessage) (*message.CategoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCategoryServiceServer) Create(context.Context, *message.CategoryCreateMessage) (*message.CategoryCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCategoryServiceServer) Update(context.Context, *message.CategoryUpdateMessage) (*message.CategoryUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCategoryServiceServer) Delete(context.Context, *message.CategoryDeleteMessage) (*message.CategoryDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCategoryServiceServer) mustEmbedUnimplementedCategoryServiceServer() {}

// UnsafeCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServiceServer will
// result in compilation errors.
type UnsafeCategoryServiceServer interface {
	mustEmbedUnimplementedCategoryServiceServer()
}

func RegisterCategoryServiceServer(s grpc.ServiceRegistrar, srv CategoryServiceServer) {
	s.RegisterService(&CategoryService_ServiceDesc, srv)
}

func _CategoryService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.CategoryGetMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CategoryService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Get(ctx, req.(*message.CategoryGetMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.CategoryListMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CategoryService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).List(ctx, req.(*message.CategoryListMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.CategoryCreateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CategoryService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Create(ctx, req.(*message.CategoryCreateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.CategoryUpdateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CategoryService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Update(ctx, req.(*message.CategoryUpdateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.CategoryDeleteMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CategoryService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Delete(ctx, req.(*message.CategoryDeleteMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryService_ServiceDesc is the grpc.ServiceDesc for CategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CategoryService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CategoryService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CategoryService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CategoryService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CategoryService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service/category.proto",
}
