package main

import (
	"context"
	"log"
	"net"

	"github.com/mgurdal/blackmarkt/proto"
	"github.com/mgurdal/blackmarkt/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UserServer struct{}

func (srv *UserServer) MoveToMarket(ctx context.Context, item *proto.Item) (*proto.Response, error) {
	name := item.GetName()
	quantity := item.GetQuantity()
	log.Println(name, quantity)
	return &proto.Response{Result: "success"}, nil

}

func (srv *UserServer) Purchase(ctx context.Context, product *proto.Product) (*proto.Response, error) {
	name := product.GetName()
	quantity := product.GetQuantity()
	price := product.GetQuantity()
	log.Println(name, quantity, price)
	return &proto.Response{Result: "result"}, nil

}

func (srv *UserServer) Collect(ctx context.Context, factory *proto.Factory) (*proto.Response, error) {
	name := factory.GetName()
	log.Println(name)
	return &proto.Response{Result: "result"}, nil

}

func main() {
	listener, err := net.Listen("tcp", ":4004")
	if err != nil {
		log.Fatal(err)
		return
	}
	store.GetStore()
	srv := grpc.NewServer()
	proto.RegisterUserServer(srv, &UserServer{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
