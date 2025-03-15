package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	metadata_pb "github.com/manavnanwani/grpc-metadata-service/proto/metadata"
	server_pb "github.com/manavnanwani/grpc-metadata-service/proto/server"

	"google.golang.org/grpc"
)

type server struct {
	metadata_pb.UnimplementedMetadataServiceServer
	server_pb.UnimplementedServerServiceServer

	mu      sync.Mutex
	servers map[string]metadata_pb.MetadataServiceClient
}

func (s *server) RegisterServer(ctx context.Context, req *server_pb.RegisterRequest) (*server_pb.RegisterResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	conn, err := grpc.Dial(req.ServerId, grpc.WithInsecure())
	if err != nil {
		return &server_pb.RegisterResponse{Success: false}, err
	}
	s.servers[req.ServerId] = metadata_pb.NewMetadataServiceClient(conn)
	fmt.Printf("**** Server %s registered\n", req.ServerId)
	return &server_pb.RegisterResponse{Success: true}, nil
}

func (s *server) collectMetadata() {
	for {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		s.mu.Lock()
		if len(s.servers) == 0 {
			s.mu.Unlock()
			continue
		}
		keys := make([]string, 0, len(s.servers))
		for k := range s.servers {
			keys = append(keys, k)
		}
		selected := keys[rand.Intn(len(keys))]
		client := s.servers[selected]
		s.mu.Unlock()

		fmt.Printf("Initiating metadata collection from server %s", selected)
		resp, err := client.CollectMetadata(context.Background(), &metadata_pb.MetadataRequest{ServerId: selected})
		if err != nil {
			log.Printf("Failed to collect metadata from %s: %v", selected, err)
			continue
		}
		fmt.Printf("Received metadata from %s: %s\n", selected, resp)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := &server{servers: make(map[string]metadata_pb.MetadataServiceClient)}
	grpcServer := grpc.NewServer()
	metadata_pb.RegisterMetadataServiceServer(grpcServer, s)
	server_pb.RegisterServerServiceServer(grpcServer, s)

	go s.collectMetadata()
	fmt.Println("Client listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
