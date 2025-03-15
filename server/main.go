package main

// import (
// 	"log"
// 	"net"

// 	greet_pb "github.com/manavnanwani/grpc-metadata-service/proto/greet"
// 	metadata_pb "github.com/manavnanwani/grpc-metadata-service/proto/metadata"

// 	"google.golang.org/grpc"
// )

// const port = ":8080"

// type helloServer struct {
// 	greet_pb.GreetServiceServer
// }

// type metadataServer struct {
// 	metadata_pb.MetadataServiceServer
// }

// func main() {
// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatalf("Failed to start the server %v", err)
// 	}

// 	grpcServer := grpc.NewServer()

// 	greet_pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
// 	metadata_pb.RegisterMetadataServiceServer(grpcServer, &metadataServer{})

// 	log.Printf("server started at %v", lis.Addr())

// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatalf("Failed to start: %v", err)
// 	}
// 	log.Printf(lis.Addr().Network())
// }
