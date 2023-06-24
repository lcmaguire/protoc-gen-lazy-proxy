package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"

	"github.com/lcmaguire/protoc-gen-lazy-proxy/sample"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8081" // this should be passed in via config
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	// services in your protoFile

	sample.RegisterSampleServiceServer(server, &sampleService{})
	reflection.Register(server) // this should perhaps be optional

	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}

type sampleService struct {
	sample.UnimplementedSampleServiceServer
}

func (s *sampleService) Sample(ctx context.Context, req *sample.SampleRequest) (*sample.SampleResponse, error) {
	log.Println("in sample")

	md, found := metadata.FromIncomingContext(ctx)
	log.Println("found ? ", found)

	for _, v := range md.Get("Content-Type") {
		log.Println(v)
	}

	return &sample.SampleResponse{Name: req.Name + " is what you sent"}, nil
}
