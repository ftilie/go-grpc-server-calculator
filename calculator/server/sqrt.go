package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/ftilie/go-grpc-servers/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SquareRoot(ctx context.Context, in *pb.SquareRootRequest) (*pb.SquareRootResponse, error) {
	log.Printf("SquareRoot function was invoked with number: %v", in.Number)

	if in.Number < 0 {
		log.Println("Received negative number, returning error")
		return nil, status.Errorf(
			codes.InvalidArgument,
			"%s", fmt.Sprintf("Cannot calculate square root of a negative number: %d", in.Number),
		)
	}

	result := &pb.SquareRootResponse{
		Result: math.Sqrt(float64(in.Number)),
	}

	log.Printf("Returning result: %f", result.Result)
	return result, nil
}
