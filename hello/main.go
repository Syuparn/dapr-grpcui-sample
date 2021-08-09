package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/syuparn/dapr-grpc-sample/hello/proto"
)

const ADDRESS = "0.0.0.0:50051"

func main() {
	listener, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	s := grpc.NewServer()
	proto.RegisterPersonServer(s, NewPersonServer())

	log.Println("server started")

	s.Serve(listener)
}
