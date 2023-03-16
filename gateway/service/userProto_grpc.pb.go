// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: service/proto/userProto.proto

package service

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

// UserProtoClient is the client API for UserProto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserProtoClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	IsLogin(ctx context.Context, in *IsLoginRequest, opts ...grpc.CallOption) (*IsLoginResponse, error)
	GetUserMsg(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userProtoClient struct {
	cc grpc.ClientConnInterface
}

func NewUserProtoClient(cc grpc.ClientConnInterface) UserProtoClient {
	return &userProtoClient{cc}
}

func (c *userProtoClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/service.UserProto/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProtoClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/service.UserProto/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProtoClient) IsLogin(ctx context.Context, in *IsLoginRequest, opts ...grpc.CallOption) (*IsLoginResponse, error) {
	out := new(IsLoginResponse)
	err := c.cc.Invoke(ctx, "/service.UserProto/IsLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProtoClient) GetUserMsg(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/service.UserProto/GetUserMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserProtoServer is the server API for UserProto service.
// All implementations must embed UnimplementedUserProtoServer
// for forward compatibility
type UserProtoServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	IsLogin(context.Context, *IsLoginRequest) (*IsLoginResponse, error)
	GetUserMsg(context.Context, *UserRequest) (*UserResponse, error)
	mustEmbedUnimplementedUserProtoServer()
}

// UnimplementedUserProtoServer must be embedded to have forward compatible implementations.
type UnimplementedUserProtoServer struct {
}

func (UnimplementedUserProtoServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserProtoServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserProtoServer) IsLogin(context.Context, *IsLoginRequest) (*IsLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsLogin not implemented")
}
func (UnimplementedUserProtoServer) GetUserMsg(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMsg not implemented")
}
func (UnimplementedUserProtoServer) mustEmbedUnimplementedUserProtoServer() {}

// UnsafeUserProtoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserProtoServer will
// result in compilation errors.
type UnsafeUserProtoServer interface {
	mustEmbedUnimplementedUserProtoServer()
}

func RegisterUserProtoServer(s grpc.ServiceRegistrar, srv UserProtoServer) {
	s.RegisterService(&UserProto_ServiceDesc, srv)
}

func _UserProto_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProtoServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserProto/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProtoServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProto_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProtoServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserProto/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProtoServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProto_IsLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProtoServer).IsLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserProto/IsLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProtoServer).IsLogin(ctx, req.(*IsLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProto_GetUserMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProtoServer).GetUserMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserProto/GetUserMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProtoServer).GetUserMsg(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserProto_ServiceDesc is the grpc.ServiceDesc for UserProto service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserProto_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.UserProto",
	HandlerType: (*UserProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserProto_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserProto_Register_Handler,
		},
		{
			MethodName: "IsLogin",
			Handler:    _UserProto_IsLogin_Handler,
		},
		{
			MethodName: "GetUserMsg",
			Handler:    _UserProto_GetUserMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/proto/userProto.proto",
}
