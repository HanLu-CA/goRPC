package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/hanlu-ca/goRPC/proto"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, name *pb.NameList) {
	log.Printf("Start to call SayHelloBidirectionalStream RPC...")
	stream, err := client.SayHelloBidirectionalStream(context.Background())
	if err != nil {
		log.Fatalf("Error while sending name: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming: %v", err)
			}

			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range name.Name {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending request: %v", err)
		}

		time.Sleep(1 * time.Second)
	}

	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional stream finished")
}
