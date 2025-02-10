// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: api/grpc/hello_world_api.proto

package hello_world

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

const (
	HelloWorldAPI_Version_FullMethodName = "/hello_world_api.helloWorldAPI/Version"
	HelloWorldAPI_Get_FullMethodName     = "/hello_world_api.helloWorldAPI/Get"
	HelloWorldAPI_Set_FullMethodName     = "/hello_world_api.helloWorldAPI/Set"
)

// HelloWorldAPIClient is the client API for HelloWorldAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloWorldAPIClient interface {
	Version(ctx context.Context, in *Version_Request, opts ...grpc.CallOption) (*Version_Response, error)
	Get(ctx context.Context, in *Get_Request, opts ...grpc.CallOption) (*Value, error)
	Set(ctx context.Context, in *Set_Request, opts ...grpc.CallOption) (*Set_Response, error)
}

type helloWorldAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloWorldAPIClient(cc grpc.ClientConnInterface) HelloWorldAPIClient {
	return &helloWorldAPIClient{cc}
}

func (c *helloWorldAPIClient) Version(ctx context.Context, in *Version_Request, opts ...grpc.CallOption) (*Version_Response, error) {
	out := new(Version_Response)
	err := c.cc.Invoke(ctx, HelloWorldAPI_Version_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloWorldAPIClient) Get(ctx context.Context, in *Get_Request, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, HelloWorldAPI_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloWorldAPIClient) Set(ctx context.Context, in *Set_Request, opts ...grpc.CallOption) (*Set_Response, error) {
	out := new(Set_Response)
	err := c.cc.Invoke(ctx, HelloWorldAPI_Set_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloWorldAPIServer is the server API for HelloWorldAPI service.
// All implementations must embed UnimplementedHelloWorldAPIServer
// for forward compatibility
type HelloWorldAPIServer interface {
	Version(context.Context, *Version_Request) (*Version_Response, error)
	Get(context.Context, *Get_Request) (*Value, error)
	Set(context.Context, *Set_Request) (*Set_Response, error)
	mustEmbedUnimplementedHelloWorldAPIServer()
}

// UnimplementedHelloWorldAPIServer must be embedded to have forward compatible implementations.
type UnimplementedHelloWorldAPIServer struct {
}

func (UnimplementedHelloWorldAPIServer) Version(context.Context, *Version_Request) (*Version_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedHelloWorldAPIServer) Get(context.Context, *Get_Request) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedHelloWorldAPIServer) Set(context.Context, *Set_Request) (*Set_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedHelloWorldAPIServer) mustEmbedUnimplementedHelloWorldAPIServer() {}

// UnsafeHelloWorldAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloWorldAPIServer will
// result in compilation errors.
type UnsafeHelloWorldAPIServer interface {
	mustEmbedUnimplementedHelloWorldAPIServer()
}

func RegisterHelloWorldAPIServer(s grpc.ServiceRegistrar, srv HelloWorldAPIServer) {
	s.RegisterService(&HelloWorldAPI_ServiceDesc, srv)
}

func _HelloWorldAPI_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Version_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldAPIServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloWorldAPI_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldAPIServer).Version(ctx, req.(*Version_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloWorldAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Get_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloWorldAPI_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldAPIServer).Get(ctx, req.(*Get_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloWorldAPI_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Set_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldAPIServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloWorldAPI_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldAPIServer).Set(ctx, req.(*Set_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloWorldAPI_ServiceDesc is the grpc.ServiceDesc for HelloWorldAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloWorldAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello_world_api.helloWorldAPI",
	HandlerType: (*HelloWorldAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _HelloWorldAPI_Version_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _HelloWorldAPI_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _HelloWorldAPI_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/hello_world_api.proto",
}
