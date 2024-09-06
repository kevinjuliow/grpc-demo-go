package main

import (
	"log"
	"net"

	pb "github.com/kevinjuliow/grpc-demo-go/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

type GreetServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(grpcServer, &GreetServer{})
	log.Printf("Server Started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start grpc server %v", err)
	}

}
