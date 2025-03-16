package main

import (
	"context"
	"fmt"

	metadata_pb "github.com/manavnanwani/grpc-metadata-service/proto/metadata"
)

func (s *server) CollectMetadata(ctx context.Context, req *metadata_pb.MetadataRequest) (*metadata_pb.DataResponse, error) {
	fmt.Printf("Received metadata collection request from %s\n", req.ServerId)
	return &metadata_pb.DataResponse{
		Message: "Success",
		Name:    s.serverID,
		Region:  "us=east-1",
	}, nil
}
