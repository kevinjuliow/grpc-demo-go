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

func callGreetClientStream(client pb.GreetServiceClient, names *pb.NameLists) {
	log.Printf("Client Stream started")

	var messages []string

	stream, err := client.GreetClientStream(context.Background())
	if err != nil {
		log.Fatalf("Failed while Client Stream %v", err)
	}

	for _, name := range names.Names {
		if err := stream.Send(&pb.GreetRequest{Name: name}); err != nil {
			log.Fatalf("Failed while sending %v", err)
		}
		log.Printf("Sent Request with name %v", name)
		messages = append(messages, name)
		time.Sleep(1 * time.Second)
	}

	resp, errClose := stream.CloseAndRecv()
	if errClose == io.EOF {
		log.Println(messages)
		log.Fatal("Client Stream Finished")
	}
	if errClose != nil {
		log.Fatalf("Error While Receiving %v", err)
	}

	log.Println(resp)
}

func callBiDirectionalStream(client pb.GreetServiceClient, names *pb.NameLists) {

	log.Println("BiDirectional Stream Started ")
	stream, err := client.GreetBiDirectionalStream(context.Background())
	if err != nil {
		log.Fatal("Could not send names")
	}
	channel := make(chan struct{})

	go func() {
		for {
			messages, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while Receiving %v", err)
			}
			log.Println(messages)
		}
		close(channel)
	}()

	for _, name := range names.Names {
		req := &pb.GreetRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	stream.CloseSend()
	<-channel
	log.Printf("BiDirectional Stream ended.")
}
