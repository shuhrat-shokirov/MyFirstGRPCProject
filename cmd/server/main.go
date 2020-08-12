package main

import (
	"github.com/shuhrat-shokirov/MyFirstGRPCProject/pkg/adder"
	"github.com/shuhrat-shokirov/MyFirstGRPCProject/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}


/*
	protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/adder.proto
	evans api/proto/adder.proto -p 8181
*/