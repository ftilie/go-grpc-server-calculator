package main

import (
	"context"
	"log"

	pb "github.com/ftilie/go-grpc-servers/calculator/proto"
)

func (s *Server) Multiply(ctx context.Context, in *pb.MultiplyRequest) (*pb.MultiplyResponse, error) {
	log.Printf("Multiply function was invoked with %v\n", in)
	return &pb.MultiplyResponse{
		Result: in.FirstNumber * in.SecondNumber,
	}, nil
}
