package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/hanlu-ca/goRPC/proto"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	name := &pb.NameList{
		Name: []string{"AlphaStream",
			"Beta Bridge",
			"Gamma Grove",
			"Delta Drive",
			"Epsilon Estate",
			"Zeta Zone",
			"Eta Echo",
			"Theta Thrive",
			"Iota Isle",
			"Kappa Key",
		},
	}

	// callSayHello(client)

	// callSayHelloServerStream(client, name)
	// callSayHelloClientStream(client, name)
	callSayHelloBidirectionalStream(client, name)
}
