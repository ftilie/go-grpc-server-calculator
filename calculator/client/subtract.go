package main

import (
	"context"
	"log"

	pb "github.com/ftilie/go-grpc-servers/calculator/proto"
)

func doSubtract(c pb.CalculatorServiceClient) {
	log.Println("doSubtract was invoked")

	res, err := c.Subtract(context.Background(), &pb.SubtractRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	})

	if err != nil {
		log.Fatalf("Error while calling Calculator RPC: %v", err)
	}

	log.Printf("Response from Calculator: %v", res.Result)
}
