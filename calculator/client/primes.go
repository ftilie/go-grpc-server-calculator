package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ftilie/go-grpc-servers/calculator/proto"
)

func doGetPrimes(c pb.CalculatorServiceClient) {
	log.Println("doGetPrimes was invoked")

	stream, err := c.GetPrimes(context.Background(), &pb.GetPrimesRequest{
		Number: 120,
	})
	if err != nil {
		log.Fatalf("Error while calling GetPrimes: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			// We have reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving response from GetPrimes: %v", err)
		}
		log.Printf("Response from GetPrimes: %v\n", res.Prime)
	}
}
