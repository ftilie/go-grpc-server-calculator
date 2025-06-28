package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ftilie/go-grpc-servers/calculator/proto"
)

func doGetAverage(c pb.CalculatorServiceClient) {
	log.Println("GetAverage function was invoked")

	stream, err := c.GetAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
		return
	}

	numbers := []int32{1, 2, 3, 4, 5}

	for _, number := range numbers {
		log.Printf("Sending number: %v", number)
		if err := stream.Send(&pb.GetAverageRequest{
			Number: number,
		}); err != nil {
			log.Fatalf("Error while sending number to GetAverage: %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from GetAverage: %v", err)
	}

	log.Printf("Response from GetAverage: %v\n", res.Average)

}
