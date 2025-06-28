package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ftilie/go-grpc-servers/calculator/proto"
)

func doGetMaximum(c pb.CalculatorServiceClient) {
	log.Println("doMaximum was invoked")

	stream, err := c.GetMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GetMaximum: %v", err)
		return
	}

	numbers := []*pb.GetMaximumRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 7},
		{Number: 2},
	}

	waitc := make(chan struct{})

	go func() {
		for _, number := range numbers {
			log.Printf("Sending number: %v", number.Number)
			if err := stream.Send(number); err != nil {
				log.Fatalf("Error while sending number to GetMaximum: %v", err)
			}
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Println("Reached end of stream")
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving response from GetMaximum: %v", err)
				break
			}
			log.Printf("Response from GetMaximum: %v\n", res.Maximum)
		}
		close(waitc)
	}()

	<-waitc // Wait for the goroutines to finish
}
