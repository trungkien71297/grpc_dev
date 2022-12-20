package main

import (
	"grpc_dev/server/pb/address"
	"grpc_dev/server/pb/bookshop"
	"grpc_dev/server/service_rpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	bookshop.RegisterInventoryServer(s, &service_rpc.BookshopServer{})
	address.RegisterAddressServer(s, &service_rpc.AddressServer{})
	if err := s.Serve(listener); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}
