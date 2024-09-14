package main

import (
	"context"
	"io"
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

func (h *GreetServer) GreetClientStream(stream pb.GreetService_GreetClientStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("Got request with name : %v", req.Name)

	}

	return nil
}

func (h *GreetServer) GreetBiDirectionalStream(stream pb.GreetService_GreetBiDirectionalStreamServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Fail when receiving %v", err)
		}

		log.Printf("Got Request with name : %v", req.Name)
	}

	return nil
}
