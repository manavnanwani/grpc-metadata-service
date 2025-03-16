package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	metadata_pb "github.com/manavnanwani/grpc-metadata-service/proto/metadata"
)

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
