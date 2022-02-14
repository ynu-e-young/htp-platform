// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: v1/capture.proto

package v1

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

// CaptureClient is the client API for Capture service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaptureClient interface {
	ReadOne(ctx context.Context, in *ReadOneRequest, opts ...grpc.CallOption) (*ImageReply, error)
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ImagesReply, error)
	ReadOneWithBinary(ctx context.Context, in *ReadOneWithBinaryRequest, opts ...grpc.CallOption) (*ImageReply, error)
	ReadAllWithBinary(ctx context.Context, in *ReadAllWithBinaryRequest, opts ...grpc.CallOption) (*ImagesReply, error)
	ReadOneWithBinaryAndCalArea(ctx context.Context, in *ReadOneWithBinaryAndCalAreaRequest, opts ...grpc.CallOption) (*ImageWithAreaReply, error)
	ReadAllWithBinaryAndCalArea(ctx context.Context, in *ReadAllWithBinaryAndCalAreaRequest, opts ...grpc.CallOption) (*ImagesWithAreaReply, error)
}

type captureClient struct {
	cc grpc.ClientConnInterface
}

func NewCaptureClient(cc grpc.ClientConnInterface) CaptureClient {
	return &captureClient{cc}
}

func (c *captureClient) ReadOne(ctx context.Context, in *ReadOneRequest, opts ...grpc.CallOption) (*ImageReply, error) {
	out := new(ImageReply)
	err := c.cc.Invoke(ctx, "/capture.service.v1.Capture/ReadOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captureClient) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ImagesReply, error) {
	out := new(ImagesReply)
	err := c.cc.Invoke(ctx, "/capture.service.v1.Capture/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captureClient) ReadOneWithBinary(ctx context.Context, in *ReadOneWithBinaryRequest, opts ...grpc.CallOption) (*ImageReply, error) {
	out := new(ImageReply)
	err := c.cc.Invoke(ctx, "/capture.service.v1.Capture/ReadOneWithBinary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captureClient) ReadAllWithBinary(ctx context.Context, in *ReadAllWithBinaryRequest, opts ...grpc.CallOption) (*ImagesReply, error) {
	out := new(ImagesReply)
	err := c.cc.Invoke(ctx, "/capture.service.v1.Capture/ReadAllWithBinary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captureClient) ReadOneWithBinaryAndCalArea(ctx context.Context, in *ReadOneWithBinaryAndCalAreaRequest, opts ...grpc.CallOption) (*ImageWithAreaReply, error) {
	out := new(ImageWithAreaReply)
	err := c.cc.Invoke(ctx, "/capture.service.v1.Capture/ReadOneWithBinaryAndCalArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captureClient) ReadAllWithBinaryAndCalArea(ctx context.Context, in *ReadAllWithBinaryAndCalAreaRequest, opts ...grpc.CallOption) (*ImagesWithAreaReply, error) {
	out := new(ImagesWithAreaReply)
	err := c.cc.Invoke(ctx, "/capture.service.v1.Capture/ReadAllWithBinaryAndCalArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaptureServer is the server API for Capture service.
// All implementations must embed UnimplementedCaptureServer
// for forward compatibility
type CaptureServer interface {
	ReadOne(context.Context, *ReadOneRequest) (*ImageReply, error)
	ReadAll(context.Context, *ReadAllRequest) (*ImagesReply, error)
	ReadOneWithBinary(context.Context, *ReadOneWithBinaryRequest) (*ImageReply, error)
	ReadAllWithBinary(context.Context, *ReadAllWithBinaryRequest) (*ImagesReply, error)
	ReadOneWithBinaryAndCalArea(context.Context, *ReadOneWithBinaryAndCalAreaRequest) (*ImageWithAreaReply, error)
	ReadAllWithBinaryAndCalArea(context.Context, *ReadAllWithBinaryAndCalAreaRequest) (*ImagesWithAreaReply, error)
	mustEmbedUnimplementedCaptureServer()
}

// UnimplementedCaptureServer must be embedded to have forward compatible implementations.
type UnimplementedCaptureServer struct {
}

func (UnimplementedCaptureServer) ReadOne(context.Context, *ReadOneRequest) (*ImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadOne not implemented")
}
func (UnimplementedCaptureServer) ReadAll(context.Context, *ReadAllRequest) (*ImagesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAll not implemented")
}
func (UnimplementedCaptureServer) ReadOneWithBinary(context.Context, *ReadOneWithBinaryRequest) (*ImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadOneWithBinary not implemented")
}
func (UnimplementedCaptureServer) ReadAllWithBinary(context.Context, *ReadAllWithBinaryRequest) (*ImagesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllWithBinary not implemented")
}
func (UnimplementedCaptureServer) ReadOneWithBinaryAndCalArea(context.Context, *ReadOneWithBinaryAndCalAreaRequest) (*ImageWithAreaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadOneWithBinaryAndCalArea not implemented")
}
func (UnimplementedCaptureServer) ReadAllWithBinaryAndCalArea(context.Context, *ReadAllWithBinaryAndCalAreaRequest) (*ImagesWithAreaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllWithBinaryAndCalArea not implemented")
}
func (UnimplementedCaptureServer) mustEmbedUnimplementedCaptureServer() {}

// UnsafeCaptureServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaptureServer will
// result in compilation errors.
type UnsafeCaptureServer interface {
	mustEmbedUnimplementedCaptureServer()
}

func RegisterCaptureServer(s grpc.ServiceRegistrar, srv CaptureServer) {
	s.RegisterService(&Capture_ServiceDesc, srv)
}

func _Capture_ReadOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptureServer).ReadOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/capture.service.v1.Capture/ReadOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptureServer).ReadOne(ctx, req.(*ReadOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Capture_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptureServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/capture.service.v1.Capture/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptureServer).ReadAll(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Capture_ReadOneWithBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadOneWithBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptureServer).ReadOneWithBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/capture.service.v1.Capture/ReadOneWithBinary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptureServer).ReadOneWithBinary(ctx, req.(*ReadOneWithBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Capture_ReadAllWithBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllWithBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptureServer).ReadAllWithBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/capture.service.v1.Capture/ReadAllWithBinary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptureServer).ReadAllWithBinary(ctx, req.(*ReadAllWithBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Capture_ReadOneWithBinaryAndCalArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadOneWithBinaryAndCalAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptureServer).ReadOneWithBinaryAndCalArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/capture.service.v1.Capture/ReadOneWithBinaryAndCalArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptureServer).ReadOneWithBinaryAndCalArea(ctx, req.(*ReadOneWithBinaryAndCalAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Capture_ReadAllWithBinaryAndCalArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllWithBinaryAndCalAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptureServer).ReadAllWithBinaryAndCalArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/capture.service.v1.Capture/ReadAllWithBinaryAndCalArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptureServer).ReadAllWithBinaryAndCalArea(ctx, req.(*ReadAllWithBinaryAndCalAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Capture_ServiceDesc is the grpc.ServiceDesc for Capture service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Capture_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "capture.service.v1.Capture",
	HandlerType: (*CaptureServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadOne",
			Handler:    _Capture_ReadOne_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _Capture_ReadAll_Handler,
		},
		{
			MethodName: "ReadOneWithBinary",
			Handler:    _Capture_ReadOneWithBinary_Handler,
		},
		{
			MethodName: "ReadAllWithBinary",
			Handler:    _Capture_ReadAllWithBinary_Handler,
		},
		{
			MethodName: "ReadOneWithBinaryAndCalArea",
			Handler:    _Capture_ReadOneWithBinaryAndCalArea_Handler,
		},
		{
			MethodName: "ReadAllWithBinaryAndCalArea",
			Handler:    _Capture_ReadAllWithBinaryAndCalArea_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/capture.proto",
}
