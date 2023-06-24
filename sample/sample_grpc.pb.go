// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: sample.proto

package sample

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

// SampleServiceClient is the client API for SampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SampleServiceClient interface {
	Sample(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error)
}

type sampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSampleServiceClient(cc grpc.ClientConnInterface) SampleServiceClient {
	return &sampleServiceClient{cc}
}

func (c *sampleServiceClient) Sample(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error) {
	out := new(SampleResponse)
	err := c.cc.Invoke(ctx, "/tutorial.SampleService/Sample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SampleServiceServer is the server API for SampleService service.
// All implementations must embed UnimplementedSampleServiceServer
// for forward compatibility
type SampleServiceServer interface {
	Sample(context.Context, *SampleRequest) (*SampleResponse, error)
	mustEmbedUnimplementedSampleServiceServer()
}

// UnimplementedSampleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSampleServiceServer struct {
}

func (UnimplementedSampleServiceServer) Sample(context.Context, *SampleRequest) (*SampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sample not implemented")
}
func (UnimplementedSampleServiceServer) mustEmbedUnimplementedSampleServiceServer() {}

// UnsafeSampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SampleServiceServer will
// result in compilation errors.
type UnsafeSampleServiceServer interface {
	mustEmbedUnimplementedSampleServiceServer()
}

func RegisterSampleServiceServer(s grpc.ServiceRegistrar, srv SampleServiceServer) {
	s.RegisterService(&SampleService_ServiceDesc, srv)
}

func _SampleService_Sample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServiceServer).Sample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tutorial.SampleService/Sample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServiceServer).Sample(ctx, req.(*SampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SampleService_ServiceDesc is the grpc.ServiceDesc for SampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tutorial.SampleService",
	HandlerType: (*SampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sample",
			Handler:    _SampleService_Sample_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sample.proto",
}
