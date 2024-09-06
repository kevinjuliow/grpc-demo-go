package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kevinjuliow/grpc-demo-go/proto"
)

func (h *GreetServer) Greet(ctx context.Context, req *pb.NoParams) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Message: "Hello ",
	}, nil
}

func (h *GreetServer) GreetServerStream(req *pb.NameLists, stream pb.GreetService_GreetServerStreamServer) error {
	log.Printf("Got Request with names : %v", req.Names)

	for _, names := range req.Names {

		response := &pb.GreetResponse{
			Message: "Hello " + names,
		}
		if err := stream.Send(response); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
