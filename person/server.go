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

	person := &Person{
		ID:   s.cnt,
		Age:  req.GetAge(),
		Name: req.GetName(),
	}
	People = append(People, person)

	res := &proto.CreateResponse{
		Person: &proto.Person{
			Id:   person.ID,
			Age:  person.Age,
			Name: person.Name,
		},
	}
	log.Printf("response: %#v\n", res)

	s.cnt++
	return res, nil
}

func (s *personServer) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	people := []*proto.Person{}

	for _, p := range People {
		person := &proto.Person{
			Id:   p.ID,
			Age:  p.Age,
			Name: p.Name,
		}
		people = append(people, person)
	}

	return &proto.ListResponse{
		People: people,
	}, nil
}
