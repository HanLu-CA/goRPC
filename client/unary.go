package main

import (
	"context"
	"log"
	"time"

	pb "github.com/hanlu-ca/goRPC/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Can not greet: %v", err)
	}
	log.Printf("Greeting: %s", res.Message)
}
