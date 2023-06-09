// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// CommentClient is the client API for Comment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentClient interface {
	OperateComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentOperationResponse, error)
	ListComments(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListCommentsResponse, error)
}

type commentClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentClient(cc grpc.ClientConnInterface) CommentClient {
	return &commentClient{cc}
}

func (c *commentClient) OperateComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentOperationResponse, error) {
	out := new(CommentOperationResponse)
	err := c.cc.Invoke(ctx, "/service.Comment/OperateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) ListComments(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListCommentsResponse, error) {
	out := new(ListCommentsResponse)
	err := c.cc.Invoke(ctx, "/service.Comment/ListComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServer is the server API for Comment service.
// All implementations must embed UnimplementedCommentServer
// for forward compatibility
type CommentServer interface {
	OperateComment(context.Context, *CommentRequest) (*CommentOperationResponse, error)
	ListComments(context.Context, *ListRequest) (*ListCommentsResponse, error)
	mustEmbedUnimplementedCommentServer()
}

// UnimplementedCommentServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServer struct {
}

func (UnimplementedCommentServer) OperateComment(context.Context, *CommentRequest) (*CommentOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OperateComment not implemented")
}
func (UnimplementedCommentServer) ListComments(context.Context, *ListRequest) (*ListCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComments not implemented")
}
func (UnimplementedCommentServer) mustEmbedUnimplementedCommentServer() {}

// UnsafeCommentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServer will
// result in compilation errors.
type UnsafeCommentServer interface {
	mustEmbedUnimplementedCommentServer()
}

func RegisterCommentServer(s grpc.ServiceRegistrar, srv CommentServer) {
	s.RegisterService(&Comment_ServiceDesc, srv)
}

func _Comment_OperateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).OperateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Comment/OperateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).OperateComment(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_ListComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).ListComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Comment/ListComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).ListComments(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Comment_ServiceDesc is the grpc.ServiceDesc for Comment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Comment",
	HandlerType: (*CommentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OperateComment",
			Handler:    _Comment_OperateComment_Handler,
		},
		{
			MethodName: "ListComments",
			Handler:    _Comment_ListComments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/service/proto/comment.proto",
}
