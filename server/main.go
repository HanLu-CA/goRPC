package main

import (
	"log"
	"net"

	pb "github.com/hanlu-ca/goRPC/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(grpcServer, &server{})
	log.Printf("Starting gRPC listener on port %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
