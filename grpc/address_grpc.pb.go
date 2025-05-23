// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/service/address.proto

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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AddressService_Get_FullMethodName    = "/AddressService/Get"
	AddressService_List_FullMethodName   = "/AddressService/List"
	AddressService_Create_FullMethodName = "/AddressService/Create"
	AddressService_Delete_FullMethodName = "/AddressService/Delete"
	AddressService_Update_FullMethodName = "/AddressService/Update"
)

// AddressServiceClient is the client API for AddressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressServiceClient interface {
	Get(ctx context.Context, in *message.AddressGetMessage, opts ...grpc.CallOption) (*message.AddressGetResponse, error)
	List(ctx context.Context, in *message.AddressListMessage, opts ...grpc.CallOption) (*message.AddressListResponse, error)
	Create(ctx context.Context, in *message.AddressCreateMessage, opts ...grpc.CallOption) (*message.AddressCreateResponse, error)
	Delete(ctx context.Context, in *message.AddressDeleteMessage, opts ...grpc.CallOption) (*message.AddressDeleteResponse, error)
	Update(ctx context.Context, in *message.AddressUpdateMessage, opts ...grpc.CallOption) (*message.AddressUpdateResponse, error)
}

type addressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressServiceClient(cc grpc.ClientConnInterface) AddressServiceClient {
	return &addressServiceClient{cc}
}

func (c *addressServiceClient) Get(ctx context.Context, in *message.AddressGetMessage, opts ...grpc.CallOption) (*message.AddressGetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(message.AddressGetResponse)
	err := c.cc.Invoke(ctx, AddressService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) List(ctx context.Context, in *message.AddressListMessage, opts ...grpc.CallOption) (*message.AddressListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(message.AddressListResponse)
	err := c.cc.Invoke(ctx, AddressService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) Create(ctx context.Context, in *message.AddressCreateMessage, opts ...grpc.CallOption) (*message.AddressCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(message.AddressCreateResponse)
	err := c.cc.Invoke(ctx, AddressService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) Delete(ctx context.Context, in *message.AddressDeleteMessage, opts ...grpc.CallOption) (*message.AddressDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(message.AddressDeleteResponse)
	err := c.cc.Invoke(ctx, AddressService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) Update(ctx context.Context, in *message.AddressUpdateMessage, opts ...grpc.CallOption) (*message.AddressUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(message.AddressUpdateResponse)
	err := c.cc.Invoke(ctx, AddressService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressServiceServer is the server API for AddressService service.
// All implementations must embed UnimplementedAddressServiceServer
// for forward compatibility.
type AddressServiceServer interface {
	Get(context.Context, *message.AddressGetMessage) (*message.AddressGetResponse, error)
	List(context.Context, *message.AddressListMessage) (*message.AddressListResponse, error)
	Create(context.Context, *message.AddressCreateMessage) (*message.AddressCreateResponse, error)
	Delete(context.Context, *message.AddressDeleteMessage) (*message.AddressDeleteResponse, error)
	Update(context.Context, *message.AddressUpdateMessage) (*message.AddressUpdateResponse, error)
	mustEmbedUnimplementedAddressServiceServer()
}

// UnimplementedAddressServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAddressServiceServer struct{}

func (UnimplementedAddressServiceServer) Get(context.Context, *message.AddressGetMessage) (*message.AddressGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAddressServiceServer) List(context.Context, *message.AddressListMessage) (*message.AddressListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAddressServiceServer) Create(context.Context, *message.AddressCreateMessage) (*message.AddressCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAddressServiceServer) Delete(context.Context, *message.AddressDeleteMessage) (*message.AddressDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAddressServiceServer) Update(context.Context, *message.AddressUpdateMessage) (*message.AddressUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAddressServiceServer) mustEmbedUnimplementedAddressServiceServer() {}
func (UnimplementedAddressServiceServer) testEmbeddedByValue()                        {}

// UnsafeAddressServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressServiceServer will
// result in compilation errors.
type UnsafeAddressServiceServer interface {
	mustEmbedUnimplementedAddressServiceServer()
}

func RegisterAddressServiceServer(s grpc.ServiceRegistrar, srv AddressServiceServer) {
	// If the following call pancis, it indicates UnimplementedAddressServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AddressService_ServiceDesc, srv)
}

func _AddressService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.AddressGetMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AddressService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).Get(ctx, req.(*message.AddressGetMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.AddressListMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AddressService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).List(ctx, req.(*message.AddressListMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.AddressCreateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AddressService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).Create(ctx, req.(*message.AddressCreateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.AddressDeleteMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AddressService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).Delete(ctx, req.(*message.AddressDeleteMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.AddressUpdateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AddressService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).Update(ctx, req.(*message.AddressUpdateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// AddressService_ServiceDesc is the grpc.ServiceDesc for AddressService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddressService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AddressService",
	HandlerType: (*AddressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _AddressService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AddressService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AddressService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AddressService_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AddressService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service/address.proto",
}
