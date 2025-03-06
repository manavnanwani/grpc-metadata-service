package main

import (
	"log"

	pb "github.com/manavnanwani/grpc-client-server/proto"

	"context"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	log.Printf("Running activity from client")
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
