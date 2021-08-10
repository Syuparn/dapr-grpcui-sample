package main

import (
	"context"
	"log"

	"github.com/syuparn/dapr-grpc-sample/person/proto"
)

func NewPersonServer() proto.PersonServer {
	return &personServer{cnt: 1}
}

type personServer struct {
	cnt int32
	// for forward compatibility
	proto.UnimplementedPersonServer
}

func (s *personServer) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	log.Printf("request: %#v\n", req)

	res := &proto.CreateResponse{
		Id:   s.cnt,
		Age:  req.GetAge(),
		Name: req.GetName(),
	}
	log.Printf("response: %#v\n", res)

	s.cnt++
	return res, nil
}
