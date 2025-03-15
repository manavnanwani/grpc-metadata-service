package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	metadata_pb "github.com/manavnanwani/grpc-metadata-service/proto/metadata"
	server_pb "github.com/manavnanwani/grpc-metadata-service/proto/server"

	"google.golang.org/grpc"
)

type server struct {
	metadata_pb.UnimplementedMetadataServiceServer
	serverID string
}

func (s *server) CollectMetadata(ctx context.Context, req *metadata_pb.MetadataRequest) (*metadata_pb.DataResponse, error) {
	fmt.Printf("Received metadata collection request from %s\n", req.ServerId)
	return &metadata_pb.DataResponse{
		Message: "Success",
		Name:    s.serverID,
		Region:  "us=east-1",
	}, nil
}

func main() {
	centralClient := "localhost:50051"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	port := r.Intn(1001) + 50000
	serverID := fmt.Sprintf("localhost:%d", port)

	fmt.Printf("Server ID: %s\n", serverID)
	fmt.Printf("Connecting to central client at %s\n", centralClient)
	conn, err := grpc.Dial(centralClient, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to central client: %v", err)
	}
	defer conn.Close()
	client := server_pb.NewServerServiceClient(conn)

	_, err = client.RegisterServer(context.Background(), &server_pb.RegisterRequest{ServerId: serverID})
	if err != nil {
		log.Fatalf("Failed to register server: %v", err)
	}

	lis, err := net.Listen("tcp", serverID)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := &server{serverID: serverID}
	grpcServer := grpc.NewServer()
	metadata_pb.RegisterMetadataServiceServer(grpcServer, s)

	fmt.Printf("Server running on %s\n", serverID)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
