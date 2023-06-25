package main

import (
	"context"
	"crypto/x509"
	"log"
	"net/http"
	"strings"
	"os"

	"github.com/bufbuild/connect-go"
	// grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/lcmaguire/protoc-gen-lazy-proxy/sample/sampleconnect"

	"github.com/lcmaguire/protoc-gen-lazy-proxy/sample"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle(
		sampleconnect.NewSampleServiceHandler(newSampleService()),
	)

	mux.Handle(
		sampleconnect.NewExtraServiceHandler(newExtraService()),
	)

	err := http.ListenAndServe(
		"localhost:8080", // todo have this be set by an env var
		// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
		// avoid x/net/http2 by using http.ListenAndServeTLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: " + err.Error())
}

func grpcDial(targetURL string, secure bool) (*grpc.ClientConn, error) {
	var creds credentials.TransportCredentials
	if !secure {
		creds = insecure.NewCredentials()
	} else {
		cp, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		creds = credentials.NewClientTLSFromCert(cp, "")
	}

	return grpc.Dial(targetURL, grpc.WithTransportCredentials(creds))
}

func newSampleService() *SampleService {
	targetURL := os.Getenv("SampleService")
	cliConn, err := grpcDial(targetURL, strings.Contains(targetURL, "localhost")) // this could be annoying for certain users.
	if err != nil {
		panic(err)
	}
	return &SampleService{
		SampleServiceClient: sample.NewSampleServiceClient(cliConn),
	}
}

type SampleService struct {
	sampleconnect.UnimplementedSampleServiceHandler
	sample.SampleServiceClient
}

func (s *SampleService) Sample(ctx context.Context, req *connect.Request[sample.SampleRequest]) (*connect.Response[sample.SampleResponse], error) {
	// todo pass req.Header() -> ctx
	res, err := s.SampleServiceClient.Sample(ctx, req.Msg)
	return &connect.Response[sample.SampleResponse]{
		Msg: res,
	}, err
}

func newExtraService() *ExtraService {
	targetURL := os.Getenv("ExtraService")
	cliConn, err := grpcDial(targetURL, strings.Contains(targetURL, "localhost")) // this could be annoying for certain users.
	if err != nil {
		panic(err)
	}
	return &ExtraService{
		ExtraServiceClient: sample.NewExtraServiceClient(cliConn),
	}
}

type ExtraService struct {
	sampleconnect.UnimplementedExtraServiceHandler
	sample.ExtraServiceClient
}

func (s *ExtraService) Extra(ctx context.Context, req *connect.Request[sample.SampleRequest]) (*connect.Response[sample.SampleResponse], error) {
	// todo pass req.Header() -> ctx
	res, err := s.ExtraServiceClient.Extra(ctx, req.Msg)
	return &connect.Response[sample.SampleResponse]{
		Msg: res,
	}, err
}
