package main

import (
	"context"
	"log"
	"time"

	pb "github.com/manavnanwani/grpc-client-server/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("failed to get response from server %v", err)
	}

	log.Printf("%s", res.Message)
}
