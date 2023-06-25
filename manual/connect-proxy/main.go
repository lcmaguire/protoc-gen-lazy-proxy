package main

import (
	"context"
	"crypto/x509"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bufbuild/connect-go"
	// grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/joho/godotenv"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/lcmaguire/protoc-gen-lazy-proxy/proto/sample/sampleconnect"

	"github.com/lcmaguire/protoc-gen-lazy-proxy/proto/sample"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.Handle(
		sampleconnect.NewSampleServiceHandler(newSampleService()),
	)

	err := http.ListenAndServe(
		"localhost:8080",
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
	cliConn, err := grpcDial(targetURL, !strings.Contains(targetURL, "localhost"))
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
	ctx = headerToContext(ctx, req.Header())
	res, err := s.SampleServiceClient.Sample(ctx, req.Msg)
	return &connect.Response[sample.SampleResponse]{
		Msg: res,
	}, err
}

func headerToContext(ctx context.Context, headers http.Header) context.Context {
	return metadata.NewIncomingContext(ctx, metadata.MD(headers))
}
