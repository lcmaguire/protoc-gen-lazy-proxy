package pkg

// TODO decide either have one struct with n number of connections, or N number of structs with 1 connection.
const LazyProxyService = `
func new{{.ServiceName}}Service(cliConn *grpc.ClientConn) *{{.ServiceName}} {
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
