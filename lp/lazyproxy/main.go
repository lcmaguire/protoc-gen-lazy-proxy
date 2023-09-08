package main

import (
	example "github.com/lcmaguire/protoc-gen-lazy-proxy/example"
	exampleconnect "github.com/lcmaguire/protoc-gen-lazy-proxy/example/exampleconnect"
)

import (
	"context"
	"crypto/x509"
	"log"
	"net/http"
	"strings"
	"os"

	"github.com/rs/cors"
	"connectrpc.com/connect"
	"github.com/joho/godotenv"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.Handle(exampleconnect.NewExampleServiceHandler(newExampleService()))

	err := http.ListenAndServe(
		"localhost:8080", // todo have this be set by an env var
		// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
		// avoid x/net/http2 by using http.ListenAndServeTLS.
		h2c.NewHandler(newCORS().Handler(mux), &http2.Server{}),
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
	for k := range headers {
		headers.Get(k)
		ctx = metadata.AppendToOutgoingContext(ctx, k, headers.Get(k))
	}
	return ctx
}

func newCORS() *cors.Cors {
	// To let web developers play with the demo service from browsers, we need a
	// very permissive CORS setup.
	return cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowOriginFunc: func(origin string) bool {
			// Allow all origins, which effectively disables CORS.
			return true
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{
			// Content-Type is in the default safelist.
			"Accept",
			"Accept-Encoding",
			"Accept-Post",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Content-Encoding",
			"Grpc-Accept-Encoding",
			"Grpc-Encoding",
			"Grpc-Message",
			"Grpc-Status",
			"Grpc-Status-Details-Bin",
			"Access-Control-Allow-Origin",
		},
	})
}

func newExampleService() *ExampleService {
	targetURL := os.Getenv("ExampleService")
	cliConn, err := grpcDial(targetURL, !strings.Contains(targetURL, "localhost")) // this could be annoying for certain users.
	if err != nil {
		panic(err)
	}
	return &ExampleService{
		ExampleServiceClient: example.NewExampleServiceClient(cliConn),
	}
}

type ExampleService struct {
	exampleconnect.UnimplementedExampleServiceHandler
	example.ExampleServiceClient
}

func (s *ExampleService) Example(ctx context.Context, req *connect.Request[example.ExampleRequest]) (*connect.Response[example.ExampleResponse], error) {
	ctx = headerToContext(ctx, req.Header())
	res, err := s.ExampleServiceClient.Example(ctx, req.Msg)
	return &connect.Response[example.ExampleResponse]{
		Msg: res,
	}, err
}
