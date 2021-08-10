package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/syuparn/dapr-grpc-sample/order/proto"
)

const Address = "localhost:50051"

func main() {
	listener, err := net.Listen("tcp", Address)
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	s := grpc.NewServer()
	proto.RegisterOrderServiceServer(s, NewOrderServiceServer())

	log.Printf("server started on %s", Address)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
