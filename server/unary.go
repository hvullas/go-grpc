package main

import (
	"context"

	pb "github.com/hvullas/go-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello I am Unary mode",
	}, nil
}
