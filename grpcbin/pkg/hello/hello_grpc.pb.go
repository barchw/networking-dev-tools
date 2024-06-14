// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: hello.proto

package hello

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

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	StreamGoats(ctx context.Context, in *Pen, opts ...grpc.CallOption) (HelloService_StreamGoatsClient, error)
	ListGoats(ctx context.Context, in *Pen, opts ...grpc.CallOption) (*GoatList, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/HelloService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) StreamGoats(ctx context.Context, in *Pen, opts ...grpc.CallOption) (HelloService_StreamGoatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[0], "/HelloService/StreamGoats", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceStreamGoatsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloService_StreamGoatsClient interface {
	Recv() (*Goat, error)
	grpc.ClientStream
}

type helloServiceStreamGoatsClient struct {
	grpc.ClientStream
}

func (x *helloServiceStreamGoatsClient) Recv() (*Goat, error) {
	m := new(Goat)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) ListGoats(ctx context.Context, in *Pen, opts ...grpc.CallOption) (*GoatList, error) {
	out := new(GoatList)
	err := c.cc.Invoke(ctx, "/HelloService/ListGoats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
type HelloServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	StreamGoats(*Pen, HelloService_StreamGoatsServer) error
	ListGoats(context.Context, *Pen) (*GoatList, error)
	mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (UnimplementedHelloServiceServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedHelloServiceServer) StreamGoats(*Pen, HelloService_StreamGoatsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamGoats not implemented")
}
func (UnimplementedHelloServiceServer) ListGoats(context.Context, *Pen) (*GoatList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGoats not implemented")
}
func (UnimplementedHelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_StreamGoats_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Pen)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServiceServer).StreamGoats(m, &helloServiceStreamGoatsServer{stream})
}

type HelloService_StreamGoatsServer interface {
	Send(*Goat) error
	grpc.ServerStream
}

type helloServiceStreamGoatsServer struct {
	grpc.ServerStream
}

func (x *helloServiceStreamGoatsServer) Send(m *Goat) error {
	return x.ServerStream.SendMsg(m)
}

func _HelloService_ListGoats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pen)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).ListGoats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HelloService/ListGoats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).ListGoats(ctx, req.(*Pen))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
		{
			MethodName: "ListGoats",
			Handler:    _HelloService_ListGoats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamGoats",
			Handler:       _HelloService_StreamGoats_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hello.proto",
}