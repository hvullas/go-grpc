package main

import (
	"log"

	pb "github.com/hvullas/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"Ullas", "Sachin", "Alice"},
	}

	/* unary api demonstration */
	//callSayHello(client)

	/* server stream demonstaration */
	//callSayHelloServerStream(client, names)

	/* client stream demonstaration */
	// callSayHelloClientStream(client, names)

	/* biderctional stream demonstaration */
	callSayHelloBidirectionalStream(client, names)

}
