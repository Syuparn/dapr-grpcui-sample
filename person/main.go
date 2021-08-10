package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/syuparn/dapr-grpc-sample/person/proto"
)

const ADDRESS = "localhost:50051"

func main() {
	listener, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	s := grpc.NewServer()
	proto.RegisterPersonServer(s, NewPersonServer())

	log.Println("server started")

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
