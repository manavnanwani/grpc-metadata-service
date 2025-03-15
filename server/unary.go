package main

// import (
// 	"context"
// 	"log"

// 	greet_pb "github.com/manavnanwani/grpc-metadata-service/proto/greet"
// 	metadata_pb "github.com/manavnanwani/grpc-metadata-service/proto/metadata"
// )

// // Hello Activity

// func (s *helloServer) SayHello(ctx context.Context, req *greet_pb.NoParam) (*greet_pb.HelloResponse, error) {
// 	log.Printf("Running Hello activity")
// 	return &greet_pb.HelloResponse{
// 		Message: "Hello",
// 	}, nil
// }

// // Metadata Activity

// func (s *metadataServer) GetData(ctx context.Context, req *metadata_pb.NoParam) (*metadata_pb.DataResponse, error) {
// 	log.Printf("Running Metadata Activity")
// 	response := metadata_pb.DataResponse{
// 		Message: "Success",
// 		Name:    "server-name-1",
// 		Region:  "us=east-1",
// 	}

// 	log.Printf("Activity complete!!")
// 	return &response, nil
// }
