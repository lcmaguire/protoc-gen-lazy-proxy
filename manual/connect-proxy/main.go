package main

import (
	"context"
	"crypto/x509"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/lcmaguire/protoc-gen-lazy-proxy/sample"
	"github.com/lcmaguire/protoc-gen-lazy-proxy/sample/sampleconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	mux := http.NewServeMux()

	sampleCliConn, err := grpcDial("localhost:8081", false)
	if err != nil {
		panic
	}
	

	mux.Handle(
		sampleconnect.NewSampleServiceHandler(newSampleService(sampleCliConn))
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

func newSampleService(cliConn *grpc.ClientConn) *sampleService {
	return &sampleService{
		cli: sample.NewSampleServiceClient(cliConn)
	}
}

// for all services in proto gen this + grpc code.
type sampleService struct {
	sampleconnect.UnimplementedSampleServiceHandler
	cli sample.SampleServiceClient
}

func (s *sampleService) Sample(ctx context.Context, req *connect.Request[sample.SampleRequest]) (*connect.Response[sample.SampleResponse], error) {
	// todo pass req.Header() -> ctx

	res, err := s.cli.Sample(ctx, req.Msg)
	return &connect.Response[sample.SampleResponse]{
		Msg: res,
	}, err
}
