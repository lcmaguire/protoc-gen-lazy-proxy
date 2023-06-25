// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: sample/sample.proto

package sampleconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-lazy-proxy/proto/sample"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// SampleServiceName is the fully-qualified name of the SampleService service.
	SampleServiceName = "sample.SampleService"
)

// SampleServiceClient is a client for the sample.SampleService service.
type SampleServiceClient interface {
	Sample(context.Context, *connect_go.Request[sample.SampleRequest]) (*connect_go.Response[sample.SampleResponse], error)
}

// NewSampleServiceClient constructs a client for the sample.SampleService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSampleServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) SampleServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &sampleServiceClient{
		sample: connect_go.NewClient[sample.SampleRequest, sample.SampleResponse](
			httpClient,
			baseURL+"/sample.SampleService/Sample",
			opts...,
		),
	}
}

// sampleServiceClient implements SampleServiceClient.
type sampleServiceClient struct {
	sample *connect_go.Client[sample.SampleRequest, sample.SampleResponse]
}

// Sample calls sample.SampleService.Sample.
func (c *sampleServiceClient) Sample(ctx context.Context, req *connect_go.Request[sample.SampleRequest]) (*connect_go.Response[sample.SampleResponse], error) {
	return c.sample.CallUnary(ctx, req)
}

// SampleServiceHandler is an implementation of the sample.SampleService service.
type SampleServiceHandler interface {
	Sample(context.Context, *connect_go.Request[sample.SampleRequest]) (*connect_go.Response[sample.SampleResponse], error)
}

// NewSampleServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSampleServiceHandler(svc SampleServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/sample.SampleService/Sample", connect_go.NewUnaryHandler(
		"/sample.SampleService/Sample",
		svc.Sample,
		opts...,
	))
	return "/sample.SampleService/", mux
}

// UnimplementedSampleServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSampleServiceHandler struct{}

func (UnimplementedSampleServiceHandler) Sample(context.Context, *connect_go.Request[sample.SampleRequest]) (*connect_go.Response[sample.SampleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sample.SampleService.Sample is not implemented"))
}
