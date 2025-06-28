package main

import (
	"io"
	"log"

	pb "github.com/ftilie/go-grpc-servers/calculator/proto"
)

func (s *Server) GetMaximum(stream pb.CalculatorService_GetMaximumServer) error {
	log.Printf("GetMaximum function was invoked")
	var max int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("End of stream reached")
			return nil
		}

		if err != nil {
			log.Printf("Error receiving from client: %v", err)
			return err
		}

		log.Printf("Received number: %d", req.Number)

		if req.Number > max || max == 0 {
			max = req.Number
		}

		err = stream.Send(&pb.GetMaximumResponse{
			Maximum: max,
		})
		if err != nil {
			log.Fatalf("Error while sending response: %v\n", err)
			return err
		}
	}
}
