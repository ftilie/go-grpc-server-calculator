package main

import (
	"context"
	"log"

	pb "github.com/ftilie/go-grpc-servers/calculator/proto"
)

func (s *Server) Divide(ctx context.Context, in *pb.DivideRequest) (*pb.DivideResponse, error) {
	log.Printf("Divide function was invoked with %v\n", in)
	return &pb.DivideResponse{
		Result: float64(in.FirstNumber) / float64(in.SecondNumber),
	}, nil
}
