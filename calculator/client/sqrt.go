package main

import (
	"context"
	"log"

	pb "github.com/ftilie/go-grpc-servers/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Printf("doSqrt function was invoked")

	re, err := c.SquareRoot(context.Background(), &pb.SquareRootRequest{
		Number: n,
	})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error from server: %s, code: %v", e.Message(), e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Printf("Invalid argument error: %s", e.Message())
			}
		} else {
			log.Fatalf("Error while calling SquareRoot: %v\n", err)
		}
	} else {
		log.Printf("SquareRoot Response: %v", re)
		if re != nil {
			log.Printf("Result: %f", re.Result)
		} else {
			log.Println("Received nil response")
		}
	}
}
