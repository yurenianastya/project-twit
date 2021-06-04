// Code generated by protoc-gen-go-methods. DO NOT EDIT.

package twit

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TwitServiceClient is the client API for TwitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TwitServiceClient interface {
	GetTwits(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (TwitService_GetTwitsClient, error)
	WriteTwit(ctx context.Context, in *Twit, opts ...grpc.CallOption) (*ResponseTwit, error)
	GetTwit(ctx context.Context, in *TwitUUID, opts ...grpc.CallOption) (*Twit, error)
	DeleteTwit(ctx context.Context, in *TwitUUID, opts ...grpc.CallOption) (*ResponseTwit, error)
}

type twitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTwitServiceClient(cc grpc.ClientConnInterface) TwitServiceClient {
	return &twitServiceClient{cc}
}

func (c *twitServiceClient) GetTwits(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (TwitService_GetTwitsClient, error) {
	stream, err := c.cc.NewStream(ctx, &TwitService_ServiceDesc.Streams[0], "/main.TwitService/getTwits", opts...)
	if err != nil {
		return nil, err
	}
	x := &twitServiceGetTwitsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TwitService_GetTwitsClient interface {
	Recv() (*Twit, error)
	grpc.ClientStream
}

type twitServiceGetTwitsClient struct {
	grpc.ClientStream
}

func (x *twitServiceGetTwitsClient) Recv() (*Twit, error) {
	m := new(Twit)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *twitServiceClient) WriteTwit(ctx context.Context, in *Twit, opts ...grpc.CallOption) (*ResponseTwit, error) {
	out := new(ResponseTwit)
	err := c.cc.Invoke(ctx, "/main.TwitService/writeTwit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitServiceClient) GetTwit(ctx context.Context, in *TwitUUID, opts ...grpc.CallOption) (*Twit, error) {
	out := new(Twit)
	err := c.cc.Invoke(ctx, "/main.TwitService/getTwit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitServiceClient) DeleteTwit(ctx context.Context, in *TwitUUID, opts ...grpc.CallOption) (*ResponseTwit, error) {
	out := new(ResponseTwit)
	err := c.cc.Invoke(ctx, "/main.TwitService/deleteTwit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TwitServiceServer is the server API for TwitService service.
// All implementations must embed UnimplementedTwitServiceServer
// for forward compatibility
type TwitServiceServer interface {
	GetTwits(*emptypb.Empty, TwitService_GetTwitsServer) error
	WriteTwit(context.Context, *Twit) (*ResponseTwit, error)
	GetTwit(context.Context, *TwitUUID) (*Twit, error)
	DeleteTwit(context.Context, *TwitUUID) (*ResponseTwit, error)
	mustEmbedUnimplementedTwitServiceServer()
}

// UnimplementedTwitServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTwitServiceServer struct {
}

func (UnimplementedTwitServiceServer) GetTwits(*emptypb.Empty, TwitService_GetTwitsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTwits not implemented")
}
func (UnimplementedTwitServiceServer) WriteTwit(context.Context, *Twit) (*ResponseTwit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteTwit not implemented")
}
func (UnimplementedTwitServiceServer) GetTwit(context.Context, *TwitUUID) (*Twit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTwit not implemented")
}
func (UnimplementedTwitServiceServer) DeleteTwit(context.Context, *TwitUUID) (*ResponseTwit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTwit not implemented")
}
func (UnimplementedTwitServiceServer) mustEmbedUnimplementedTwitServiceServer() {}

// UnsafeTwitServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TwitServiceServer will
// result in compilation errors.
type UnsafeTwitServiceServer interface {
	mustEmbedUnimplementedTwitServiceServer()
}

func RegisterTwitServiceServer(s grpc.ServiceRegistrar, srv TwitServiceServer) {
	s.RegisterService(&TwitService_ServiceDesc, srv)
}

func _TwitService_GetTwits_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TwitServiceServer).GetTwits(m, &twitServiceGetTwitsServer{stream})
}

type TwitService_GetTwitsServer interface {
	Send(*Twit) error
	grpc.ServerStream
}

type twitServiceGetTwitsServer struct {
	grpc.ServerStream
}

func (x *twitServiceGetTwitsServer) Send(m *Twit) error {
	return x.ServerStream.SendMsg(m)
}

func _TwitService_WriteTwit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Twit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceServer).WriteTwit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TwitService/writeTwit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceServer).WriteTwit(ctx, req.(*Twit))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitService_GetTwit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwitUUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceServer).GetTwit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TwitService/getTwit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceServer).GetTwit(ctx, req.(*TwitUUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitService_DeleteTwit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwitUUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceServer).DeleteTwit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TwitService/deleteTwit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceServer).DeleteTwit(ctx, req.(*TwitUUID))
	}
	return interceptor(ctx, in, info, handler)
}

// TwitService_ServiceDesc is the grpc.ServiceDesc for TwitService service.
// It's only intended for direct use with methods.RegisterService,
// and not to be introspected or modified (even as a copy)
var TwitService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.TwitService",
	HandlerType: (*TwitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "writeTwit",
			Handler:    _TwitService_WriteTwit_Handler,
		},
		{
			MethodName: "getTwit",
			Handler:    _TwitService_GetTwit_Handler,
		},
		{
			MethodName: "deleteTwit",
			Handler:    _TwitService_DeleteTwit_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getTwits",
			Handler:       _TwitService_GetTwits_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/twit/twit.proto",
}
