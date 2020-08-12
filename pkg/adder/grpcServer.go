package adder

import (
	"context"
	"github.com/shuhrat-shokirov/MyFirstGRPCProject/pkg/api"
)

type GRPCServer struct {}

func (s *GRPCServer) Add(ctx context.Context,req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX()+req.GetY()}, nil
}

