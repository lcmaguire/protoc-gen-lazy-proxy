package main

import (
	"context"
	"crypto/x509"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
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

	/* todo see if reflect can be added in too.
	reflector := grpcreflect.NewStaticReflector(
		"sample.SampleService",
	)

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	*/

	mux.Handle(
		sampleconnect.NewSampleServiceHandler(newSampleService()),
		// sampleconnect.NewSampleServiceHandler(newSampleService(sampleCliConn)),
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
	cliConn, err := grpcDial("localhost:8081", false)
	if err != nil {
		panic(err)
	}

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
