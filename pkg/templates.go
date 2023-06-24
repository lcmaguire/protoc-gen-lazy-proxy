package pkg

// sampled from https://connect.build/docs/go/getting-started & demo connect repo

// ServerTemplate template for a connect-go gRPC / HTTP server.
const ServerTemplate = `
package main 

import (
	"log" 
	"net/http"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"

	// your protoPathHere
	"{{.GenImportPath}}connect"

	// your services
	{{.ServiceImports}}
)


func main() {
	mux := http.NewServeMux()
	
	reflector := grpcreflect.NewStaticReflector(
		{{.FullName}}
	  )
	
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	// The generated constructors return a path and a plain net/http
	// handler.
	{{.Services}}
	err := http.ListenAndServe(
	  "localhost:8080",
	  // For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
	  // avoid x/net/http2 by using http.ListenAndServeTLS.
	  h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: " + err.Error())
  }
  
`

// TODO add in health. (or try using for loops within templates)
const ServiceHandleTemplate = `

mux.Handle({{.Pkg}}connect.New{{.ServiceName}}Handler(&{{.ServiceStruct}}{}))
`

// ServiceTemplate template for the body of a file that creates a struct for your service handler + a constructor.
const ServiceTemplate = `
package {{.GoPkgName}}

import (
	{{.Imports}}
)

// {{.ServiceName}} implements {{.FullName}}.
type {{.ServiceName}} struct { 
	{{.Pkg}}.Unimplemented{{.ServiceName}}Handler
}
		
func New{{.ServiceName}} () *{{.ServiceName}} {
	return &{{.ServiceName}}{}
}
`

// MethodTemplate template for an unimplemented unary connect-go gRPC method.
const MethodTemplate = `

package {{.GoPkgName}}

import (
	"context"
	"errors"
	connect_go "github.com/bufbuild/connect-go"

	{{.Imports}}
)

// {{.MethodName}} implements {{.FullName}}.
func (s * {{.ServiceName}}) {{.MethodName}}(ctx context.Context, req *connect_go.Request[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	res := connect_go.NewResponse(&{{.ResponseType}}{})
	return res, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("not yet implemented"))
}

`
