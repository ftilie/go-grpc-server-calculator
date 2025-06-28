package main

import (
	"context"
	"log"

	pb "github.com/ftilie/go-grpc-servers/calculator/proto"
)

func doMultiply(c pb.CalculatorServiceClient) {
	log.Println("doMultiply was invoked")

	res, err := c.Multiply(context.Background(), &pb.MultiplyRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	})

	if err != nil {
		log.Fatalf("Error while calling Calculator RPC: %v", err)
	}

	log.Printf("Response from Calculator: %v", res.Result)
}
