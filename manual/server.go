package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"github.com/lcmaguire/protoc-gen-lazy-proxy/sample"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8080" // this should be passed in via config
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	// services in your protoFile

	sample.RegisterSampleService(server, &sampleService{})
	reflection.Register(server) // this should perhaps be optional

	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}

type sampleService struct{}

func (s *sampleService) Sample(ctx context.Context)
