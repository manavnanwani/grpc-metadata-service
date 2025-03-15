package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	greet_pb "github.com/manavnanwani/grpc-metadata-service/proto/greet"
	metadata_pb "github.com/manavnanwani/grpc-metadata-service/proto/metadata"
)

const port = ":8080"

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials((insecure.NewCredentials())))
	if err != nil {
		log.Fatalf("Failed to connect to the server %v", err)
	}

	defer conn.Close()

	client := greet_pb.NewGreetServiceClient(conn)
	callSayHello(client)

	metadata_client := metadata_pb.NewMetadataServiceClient(conn)
	getMetadataWorkflow(metadata_client)
}
