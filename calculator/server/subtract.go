package main

import (
	"context"
	"log"

	pb "github.com/ftilie/go-grpc-servers/calculator/proto"
)

func (s *Server) Subtract(ctx context.Context, in *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	log.Printf("Subtract function was invoked with %v\n", in)
	return &pb.SubtractResponse{
		Result: in.FirstNumber - in.SecondNumber,
	}, nil
}
