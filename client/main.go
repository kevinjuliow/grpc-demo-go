package main

import (
	"log"

	pb "github.com/kevinjuliow/grpc-demo-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost"+port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("Failed to Connect %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameLists{
		Names: []string{"Kevin", "Julio", "Walter"},
	}

	// callGreet(client)
	// callGreetServerStream(client, names)
	// callGreetClientStream(client, names)
	callBiDirectionalStream(client, names)
}
