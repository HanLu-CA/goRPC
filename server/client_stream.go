package main

import (
	"io"
	"log"

	pb "github.com/hanlu-ca/goRPC/proto"
)

func (s *server) SayHelloClientStream(stream pb.GreetService_SayHelloClientStreamServer) error {
	var message []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Message: message})
		}
		if err != nil {
			return err
		}
		log.Printf("Received: %v", req.Name)
		message = append(message, "Hello", req.Name)
	}
}
