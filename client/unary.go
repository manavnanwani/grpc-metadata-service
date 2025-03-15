package main

import (
	"context"
	"log"
	"time"

	greet_pb "github.com/manavnanwani/grpc-metadata-service/proto/greet"
	metadata_pb "github.com/manavnanwani/grpc-metadata-service/proto/metadata"
)

// Hello Workflow

func callSayHello(client greet_pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	log.Printf("Initiating hello request to server")
	res, err := client.SayHello(ctx, &greet_pb.NoParam{})
	if err != nil {
		log.Fatalf("failed to get response from server %v", err)
	}

	log.Printf("Response: %s", res.Message)
}

// Metadata Workflow

func getMetadataWorkflow(client metadata_pb.MetadataServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	log.Printf("Initiating metadata request to server")
	res, err := client.GetData(ctx, &metadata_pb.NoParam{})

	if err != nil {
		log.Fatalf("failed to get response from server %v", err)
	}

	log.Printf("Response: %s", res)
}
