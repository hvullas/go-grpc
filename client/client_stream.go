package main

import (
	"context"
	"log"
	"time"

	pb "github.com/hvullas/go-grpc/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Client streaming started")

	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for _, names := range names.Names {
		req := &pb.HelloRequest{
			Message: names,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}

		log.Printf("Sent the request with name: %s", names)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished")
	if err != nil {
		log.Fatalf("Error while recieving %v", err)
	}

	log.Printf("%v", res.Messages)
}
