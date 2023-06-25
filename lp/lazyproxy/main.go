package main

import "github.com/lcmaguire/protoc-gen-lazy-proxy/sample/sampleconnect"
import "github.com/lcmaguire/protoc-gen-lazy-proxy/sample"
import "github.com/bufbuild/connect-go"
import "google.golang.org/grpc"
import "context"

func newSampleServiceService(cliConn *grpc.ClientConn) *SampleService {
	return &SampleService{
		cli: sample.NewSampleServiceClient(cliConn),
	}
}

type SampleService struct {
	sampleconnect.UnimplementedSampleServiceHandler
	cli sample.SampleServiceClient
}

func (s *SampleService) Sample(ctx context.Context, req *connect.Request[sample.SampleRequest]) (*connect.Response[sample.SampleResponse], error) {
	// todo pass req.Header() -> ctx
	// for headers desired, get -> write to outgoing metadata
	res, err := s.cli.Sample(ctx, req.Msg)
	return &connect.Response[sample.SampleResponse]{
		Msg: res,
	}, err
}
