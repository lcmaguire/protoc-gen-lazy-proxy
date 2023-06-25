# protoc-gen-lazy-proxy


`protoc-gen-lazy-proxy` will generate a connect-go server which as a proxy for an existing grpc server(s) which will convert and forward http1 traffic to the existing grpc server.

- Allows web clients to intgegrate with gRPC services that have not set up handling serving http1 traffic.
- Allows curl to be used against gRPC services


Supports the following RPC methods.

|                     | connect-go         |
| ------------------- | ------------------ |
| unary RPC methods     | :white_check_mark: |
| streaming RPC methods |  |


## instalation
```
go install github.com/lcmaguire/protoc-gen-lazy-proxy@latest

or 

go get github.com/lcmaguire/protoc-gen-lazy-proxy@latest
```

## Config

include in an `.env` file the ServiceName with the url you would want it to forward to.

```
{{ServiceName}}={{targetURL}}
...
SampleService=localhost:8082
```

## Useage

will serve traffic on localhost:8080 for all of the services included in the proto generation

will write file to `.{{out}}/lazyproxy/main.go`

reccomended to use buf to generate proto files. see below for a sample buf.gen.yaml

```
version: v1
plugins:
  - name: go
    out: proto
    opt: paths=source_relative
  - name: go-grpc
    out: proto
    opt: paths=source_relative
  - name: connect-go
    out: proto
    opt: paths=source_relative
  - name: lazy-proxy
    out: lp # will create file in lp/lazyproxy/main.go
    opt: paths=source_relative
    strategy: all # currently will generate one executeable for all protoc files included.
```
