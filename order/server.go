package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/syuparn/dapr-grpc-sample/order/proto"
)

func NewOrderServiceServer() proto.OrderServiceServer {
	return &OrderServer{}
}

type OrderServer struct {
	// for forward compatibility
	proto.UnimplementedOrderServiceServer
}

func (s *OrderServer) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	log.Printf("request: %#v\n", req)

	order := &Order{
		ID:    newID(),
		Drink: req.GetDrink().String(),
		Time:  time.Now(),
	}
	Orders = append(Orders, order)

	res := &proto.CreateResponse{
		Order: &proto.Order{
			Id:    order.ID,
			Drink: proto.Drink(proto.Drink_value[order.Drink]),
			Time:  timestamppb.New(order.Time),
		},
	}
	log.Printf("response: %#v\n", res)

	return res, nil
}

func (s *OrderServer) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	orders := []*proto.Order{}

	for _, o := range Orders {
		order := &proto.Order{
			Id:    o.ID,
			Drink: proto.Drink(proto.Drink_value[o.Drink]),
			Time:  timestamppb.New(o.Time),
		}
		orders = append(orders, order)
	}

	return &proto.ListResponse{
		Orders: orders,
	}, nil
}
