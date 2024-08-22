// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.17.3
// source: message.proto

package message

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Mq_Pop_FullMethodName    = "/message.mq/pop"
	Mq_Push_FullMethodName   = "/message.mq/push"
	Mq_HasMsg_FullMethodName = "/message.mq/hasMsg"
)

// MqClient is the client API for Mq service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MqClient interface {
	Pop(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Message, error)
	Push(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error)
	HasMsg(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BoolMsg, error)
}

type mqClient struct {
	cc grpc.ClientConnInterface
}

func NewMqClient(cc grpc.ClientConnInterface) MqClient {
	return &mqClient{cc}
}

func (c *mqClient) Pop(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Message, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Message)
	err := c.cc.Invoke(ctx, Mq_Pop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqClient) Push(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Mq_Push_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqClient) HasMsg(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BoolMsg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BoolMsg)
	err := c.cc.Invoke(ctx, Mq_HasMsg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MqServer is the server API for Mq service.
// All implementations must embed UnimplementedMqServer
// for forward compatibility.
type MqServer interface {
	Pop(context.Context, *emptypb.Empty) (*Message, error)
	Push(context.Context, *Message) (*emptypb.Empty, error)
	HasMsg(context.Context, *emptypb.Empty) (*BoolMsg, error)
	mustEmbedUnimplementedMqServer()
}

// UnimplementedMqServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMqServer struct{}

func (UnimplementedMqServer) Pop(context.Context, *emptypb.Empty) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pop not implemented")
}
func (UnimplementedMqServer) Push(context.Context, *Message) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedMqServer) HasMsg(context.Context, *emptypb.Empty) (*BoolMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasMsg not implemented")
}
func (UnimplementedMqServer) mustEmbedUnimplementedMqServer() {}
func (UnimplementedMqServer) testEmbeddedByValue()            {}

// UnsafeMqServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MqServer will
// result in compilation errors.
type UnsafeMqServer interface {
	mustEmbedUnimplementedMqServer()
}

func RegisterMqServer(s grpc.ServiceRegistrar, srv MqServer) {
	// If the following call pancis, it indicates UnimplementedMqServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Mq_ServiceDesc, srv)
}

func _Mq_Pop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServer).Pop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mq_Pop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServer).Pop(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mq_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mq_Push_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServer).Push(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mq_HasMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServer).HasMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mq_HasMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServer).HasMsg(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Mq_ServiceDesc is the grpc.ServiceDesc for Mq service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mq_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.mq",
	HandlerType: (*MqServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "pop",
			Handler:    _Mq_Pop_Handler,
		},
		{
			MethodName: "push",
			Handler:    _Mq_Push_Handler,
		},
		{
			MethodName: "hasMsg",
			Handler:    _Mq_HasMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
