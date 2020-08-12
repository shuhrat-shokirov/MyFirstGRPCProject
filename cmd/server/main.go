package main

import (
	"google.golang.org/grpc"
	"grpc/pkg/adder"
	"grpc/pkg/api"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8181")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}


/*
	protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/adder.proto
	./evans api/proto/adder.proto -p 8181
*/