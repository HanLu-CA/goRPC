package main

import (
	"context"
	"log"
	"time"

	pb "github.com/hanlu-ca/goRPC/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, name *pb.NameList) {
	log.Printf("Start to call SayHelloClientStreaming RPC...")
	stream, err := client.SayHelloClientStream(context.Background())
	if err != nil {
		log.Fatalf("Error while calling SayHelloClientStreaming RPC: %v", err)
	}

	for _, name := range name.Name {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending request: %v", err)
		}
		log.Printf("Sending: %v", req.Name)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client stream finished")
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}
	log.Printf("Response: %v", res.Message)
}
