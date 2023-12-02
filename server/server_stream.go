package main

import (
	"log"
	"time"

	pb "github.com/hanlu-ca/goRPC/proto"
)

func (s *server) SayHelloServerStream(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamServer) error {
	log.Printf("Received: %v", req.Name)
	for _, name := range req.Name {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
