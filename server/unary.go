package main

import (
	"context"

	pb "github.com/hanlu-ca/goRPC/proto"
)

func (s *server) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello from the server!",
	}, nil
}
