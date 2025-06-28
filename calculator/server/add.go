package main

import (
	"context"
	"log"

	pb "github.com/ftilie/go-grpc-servers/calculator/proto"
)

func (s *Server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Add function was invoked with %v\n", in)
	return &pb.AddResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
