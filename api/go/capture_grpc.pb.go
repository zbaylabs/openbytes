// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: capture.proto

package openbytes

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Captures_Device_FullMethodName  = "/openbytes.Captures/Device"
	Captures_List_FullMethodName    = "/openbytes.Captures/List"
	Captures_Traffic_FullMethodName = "/openbytes.Captures/Traffic"
	Captures_Copy_FullMethodName    = "/openbytes.Captures/Copy"
)

// CapturesClient is the client API for Captures service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CapturesClient interface {
	Device(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*structpb.ListValue, error)
	List(ctx context.Context, in *Capture, opts ...grpc.CallOption) (Captures_ListClient, error)
	Traffic(ctx context.Context, in *Capture, opts ...grpc.CallOption) (Captures_TrafficClient, error)
	Copy(ctx context.Context, in *CopyRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
}

type capturesClient struct {
	cc grpc.ClientConnInterface
}

func NewCapturesClient(cc grpc.ClientConnInterface) CapturesClient {
	return &capturesClient{cc}
}

func (c *capturesClient) Device(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*structpb.ListValue, error) {
	out := new(structpb.ListValue)
	err := c.cc.Invoke(ctx, Captures_Device_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capturesClient) List(ctx context.Context, in *Capture, opts ...grpc.CallOption) (Captures_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Captures_ServiceDesc.Streams[0], Captures_List_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &capturesListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Captures_ListClient interface {
	Recv() (*Packet, error)
	grpc.ClientStream
}

type capturesListClient struct {
	grpc.ClientStream
}

func (x *capturesListClient) Recv() (*Packet, error) {
	m := new(Packet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *capturesClient) Traffic(ctx context.Context, in *Capture, opts ...grpc.CallOption) (Captures_TrafficClient, error) {
	stream, err := c.cc.NewStream(ctx, &Captures_ServiceDesc.Streams[1], Captures_Traffic_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &capturesTrafficClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Captures_TrafficClient interface {
	Recv() (*Point, error)
	grpc.ClientStream
}

type capturesTrafficClient struct {
	grpc.ClientStream
}

func (x *capturesTrafficClient) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *capturesClient) Copy(ctx context.Context, in *CopyRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, Captures_Copy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CapturesServer is the server API for Captures service.
// All implementations must embed UnimplementedCapturesServer
// for forward compatibility
type CapturesServer interface {
	Device(context.Context, *emptypb.Empty) (*structpb.ListValue, error)
	List(*Capture, Captures_ListServer) error
	Traffic(*Capture, Captures_TrafficServer) error
	Copy(context.Context, *CopyRequest) (*wrapperspb.StringValue, error)
	mustEmbedUnimplementedCapturesServer()
}

// UnimplementedCapturesServer must be embedded to have forward compatible implementations.
type UnimplementedCapturesServer struct {
}

func (UnimplementedCapturesServer) Device(context.Context, *emptypb.Empty) (*structpb.ListValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Device not implemented")
}
func (UnimplementedCapturesServer) List(*Capture, Captures_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCapturesServer) Traffic(*Capture, Captures_TrafficServer) error {
	return status.Errorf(codes.Unimplemented, "method Traffic not implemented")
}
func (UnimplementedCapturesServer) Copy(context.Context, *CopyRequest) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Copy not implemented")
}
func (UnimplementedCapturesServer) mustEmbedUnimplementedCapturesServer() {}

// UnsafeCapturesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CapturesServer will
// result in compilation errors.
type UnsafeCapturesServer interface {
	mustEmbedUnimplementedCapturesServer()
}

func RegisterCapturesServer(s grpc.ServiceRegistrar, srv CapturesServer) {
	s.RegisterService(&Captures_ServiceDesc, srv)
}

func _Captures_Device_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapturesServer).Device(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Captures_Device_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapturesServer).Device(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Captures_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Capture)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CapturesServer).List(m, &capturesListServer{stream})
}

type Captures_ListServer interface {
	Send(*Packet) error
	grpc.ServerStream
}

type capturesListServer struct {
	grpc.ServerStream
}

func (x *capturesListServer) Send(m *Packet) error {
	return x.ServerStream.SendMsg(m)
}

func _Captures_Traffic_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Capture)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CapturesServer).Traffic(m, &capturesTrafficServer{stream})
}

type Captures_TrafficServer interface {
	Send(*Point) error
	grpc.ServerStream
}

type capturesTrafficServer struct {
	grpc.ServerStream
}

func (x *capturesTrafficServer) Send(m *Point) error {
	return x.ServerStream.SendMsg(m)
}

func _Captures_Copy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapturesServer).Copy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Captures_Copy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapturesServer).Copy(ctx, req.(*CopyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Captures_ServiceDesc is the grpc.ServiceDesc for Captures service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Captures_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openbytes.Captures",
	HandlerType: (*CapturesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Device",
			Handler:    _Captures_Device_Handler,
		},
		{
			MethodName: "Copy",
			Handler:    _Captures_Copy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Captures_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Traffic",
			Handler:       _Captures_Traffic_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "capture.proto",
}
