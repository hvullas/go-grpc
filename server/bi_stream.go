package main

import (
	"io"
	"log"

	pb "github.com/hvullas/go-grpc/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Printf("Got request with name: %v", req.Message)
		res := &pb.HelloResponse{
			Message: req.Message,
		}
		if err := stream.SendMsg(res); err != nil {
			return err
		}
	}
}
