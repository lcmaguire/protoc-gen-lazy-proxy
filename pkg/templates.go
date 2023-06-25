package pkg

// LazyProxyService template for initialzing a grpc service and its methods for forwarding http1 traffic.
const LazyProxyService = `
func new{{.ServiceName}}() *{{.ServiceName}} {
	targetURL := os.Getenv("{{.ServiceName}}")
	cliConn, err := grpcDial(targetURL, !strings.Contains(targetURL, "localhost")) // this could be annoying for certain users.
	if err != nil {
		panic(err)
	}
	return &{{.ServiceName}}{
		{{.ServiceName}}Client: {{.Pkg}}New{{.ServiceName}}Client(cliConn),
	}
}

type {{.ServiceName}} struct {
	{{.ConnectPkg}}Unimplemented{{.ServiceName}}Handler
	{{.Pkg}}{{.ServiceName}}Client
}

{{range  .Methods}}
func (s *{{.ServiceName}}) {{.MethodName}}(ctx context.Context, req *connect.Request[{{.RequestName}}]) (*connect.Response[{{.ResponseName}}], error) {
	ctx = headerToContext(ctx, req.Header())
	res, err := s.{{.ServiceName}}Client.{{.MethodName}}(ctx, req.Msg)
	return &connect.Response[{{.ResponseName}}]{
		Msg: res,
	}, err
}
{{end}}

`

// LazyProxyServiceInfo  contains info for a service including its name, pkg and any methods.
type LazyProxyServiceInfo struct {
	ServiceName string
	Pkg         string
	ConnectPkg  string
	Methods     []LazyProxyMethodInfo
}

// LazyProxyMethodInfo contains all info required for a method.
type LazyProxyMethodInfo struct {
	ServiceName  string
	MethodName   string
	RequestName  string
	ResponseName string
}

// LazyProxyServerInfo contains the services it needs to run and any dynamic proto imports.
type LazyProxyServerInfo struct {
	Services []LazyProxyServiceInfo
	Imports  []string
}

// LazyProxyServer Template for the server that will run the lazy proxy.
const LazyProxyServer = `

import (
	"context"
	"crypto/x509"
	"log"
	"net/http"
	"strings"
	"os"

	"github.com/bufbuild/connect-go"
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
	{{range  .Services}}
		mux.Handle({{.ConnectPkg}}New{{.ServiceName}}Handler(new{{.ServiceName}}()))
	{{end}}

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
`
