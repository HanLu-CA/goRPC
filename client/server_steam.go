package main

import (
	"context"
	"io"
	"log"

	pb "github.com/hanlu-ca/goRPC/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, name *pb.NameList) {
	log.Printf("Start to call SayHelloServerStreaming RPC...")
	stream, err := client.SayHelloServerStream(context.Background(), name)

	if err != nil {
		log.Fatalf("Error while calling SayHelloServerStreaming RPC: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Println(message)
	}

	log.Printf("Streaming finished")

}
