package main

import (
	"context"
	"crypto/x509"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bufbuild/connect-go"
	"github.com/joho/godotenv"
	v1 "github.com/lcmaguire/protoc-gen-lazy-proxy/proto/sample/v1"
	samplev1connect "github.com/lcmaguire/protoc-gen-lazy-proxy/proto/sample/v1/samplev1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// connectFileName ? ident samplev1connect
// generatedFilenamePrefixToSlash ? ident sample/v1/sample
// connectPath ? ident sample/v1/samplev1connect
// proto ident v1.

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.Handle(samplev1connect.NewSampleServiceHandler(newSampleService()))

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

// this should probably be handled by middleware, but lazy implementation for a lazy proxy.
func headerToContext(ctx context.Context, headers http.Header) context.Context {
	return metadata.NewIncomingContext(ctx, metadata.MD(headers))
}

func newSampleService() *SampleService {
	targetURL := os.Getenv("SampleService")
	cliConn, err := grpcDial(targetURL, !strings.Contains(targetURL, "localhost")) // this could be annoying for certain users.
	if err != nil {
		panic(err)
	}
	return &SampleService{
		SampleServiceClient: v1.NewSampleServiceClient(cliConn),
	}
}

type SampleService struct {
	samplev1connect.UnimplementedSampleServiceHandler
	v1.SampleServiceClient
}

func (s *SampleService) Sample(ctx context.Context, req *connect.Request[v1.SampleRequest]) (*connect.Response[v1.SampleResponse], error) {
	ctx = headerToContext(ctx, req.Header())
	res, err := s.SampleServiceClient.Sample(ctx, req.Msg)
	return &connect.Response[v1.SampleResponse]{
		Msg: res,
	}, err
}
