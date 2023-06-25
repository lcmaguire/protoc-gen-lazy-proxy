package pkg

// TODO decide either have one struct with n number of connections, or N number of structs with 1 connection.

// todo handle creating grpcDial within this struct

const LazyProxyService = `
func new{{.ServiceName}}() *{{.ServiceName}} {
	cliConn, err := grpcDial("localhost:8081", false)
	if err != nil {
		panic(err)
	}
	return &{{.ServiceName}}{
		cli: {{.Pkg}}.New{{.ServiceName}}Client(cliConn),
	}
}

type {{.ServiceName}} struct {
	{{.Pkg}}connect.Unimplemented{{.ServiceName}}Handler
	cli {{.Pkg}}.{{.ServiceName}}Client
}
`

/*
	read config
		ServiceName
		URL
		bool

	INIT {}

	struct {
		.... all services
	}
*/

type LazyProxyServiceInfo struct {
	ServiceName string
	Pkg         string
}

const LazyProxyMethod = `
func (s *{{.ServiceName}}) {{.MethodName}}(ctx context.Context, req *connect.Request[{{.RequestName}}]) (*connect.Response[{{.ResponseName}}], error) {
	// todo pass req.Header() -> ctx
	// for headers desired, get -> write to outgoing metadata
	res, err := s.cli.Sample(ctx, req.Msg)
	return &connect.Response[{{.ResponseName}}]{
		Msg: res,
	}, err
}

`

type LazyProxyMethodInfo struct {
	ServiceName  string
	MethodName   string
	RequestName  string
	ResponseName string
}

type LazyProxyServerInfo struct {
	Services []LazyProxyServiceInfo
	Imports  []string
}

const LazyProxyServer = `

import (
	"context"
	"crypto/x509"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	{{range  .Imports}}
	{{.}}
	{{end}}
)

func main() {
	mux := http.NewServeMux()

	{{range  .Services}}
		mux.Handle(
			{{.Pkg}}connect.New{{.ServiceName}}Handler(new{{.ServiceName}}()),
			// sampleconnect.NewSampleServiceHandler(newSampleService(sampleCliConn)),
		)
	{{end}}

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
`
