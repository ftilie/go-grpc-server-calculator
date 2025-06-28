package main

import (
	"io"
	"log"

	pb "github.com/ftilie/go-grpc-servers/calculator/proto"
)

func (s *Server) GetAverage(stream pb.CalculatorService_GetAverageServer) error {
	log.Println("GetAverage function was invoked")

	var sum int32 = 0
	var count int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			log.Println("End of stream reached, sending average response")
			break
		}

		if err != nil {
			log.Printf("Error while receiving from stream: %v\n", err)
			return err
		}

		log.Printf("Received number: %d\n", req.Number)

		sum += req.Number
		count++
	}

	if count == 0 {
		return stream.SendAndClose(&pb.GetAverageResponse{Average: 0})
	}

	average := float64(sum) / float64(count)
	return stream.SendAndClose(&pb.GetAverageResponse{Average: average})
}
