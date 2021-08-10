package main

import (
	"context"
	"log"

	"github.com/syuparn/dapr-grpc-sample/person/proto"
)

func NewPersonServiceServer() proto.PersonServiceServer {
	return &personServer{cnt: 1}
}

type personServer struct {
	cnt int32
	// for forward compatibility
	proto.UnimplementedPersonServiceServer
}

func (s *personServer) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	log.Printf("request: %#v\n", req)

	res := &proto.CreateResponse{
		Person: &proto.Person{
			Id:   s.cnt,
			Age:  req.GetAge(),
			Name: req.GetName(),
		},
	}
	log.Printf("response: %#v\n", res)

	s.cnt++
	return res, nil
}
