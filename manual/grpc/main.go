package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	example "github.com/lcmaguire/protoc-gen-lazy-proxy/example"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8082" // this should be passed in via config
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	example.RegisterExampleServiceServer(server, &exampleService{})
	reflection.Register(server) // this should perhaps be optional

	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}

type exampleService struct {
	example.UnimplementedExampleServiceServer
}

func (s *exampleService) Example(ctx context.Context, req *example.ExampleRequest) (*example.ExampleResponse, error) {
	return &example.ExampleResponse{Name: req.Name + " is what you sent"}, nil
}
