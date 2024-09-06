package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/kevinjuliow/grpc-demo-go/proto"
)

func callGreet(client pb.GreetServiceClient) {
	ctx, timeOut := context.WithTimeout(context.Background(), time.Second)

	defer timeOut()

	res, err := client.Greet(ctx, &pb.NoParams{})
	if err != nil {
		log.Printf("Greet  failed: %v", err)
	}

	log.Printf(res.Message)

}

func callGreetServerStream(client pb.GreetServiceClient, names *pb.NameLists) {
	log.Printf("Stream Started")

	stream, err := client.GreetServerStream(context.Background(), names)
	if err != nil {
		log.Fatalf("error stream %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF { //if End of file
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming %v", err)
		}

		log.Println(res)
	}

	log.Println("Stream Finished")
}
