package main

import (
	"log"

	pb "github.com/ftilie/go-grpc-servers/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) GetPrimes(in *pb.GetPrimesRequest, stream grpc.ServerStreamingServer[pb.GetPrimesResponse]) error {
	log.Printf("Primes function was invoked with %v\n", in)

	// Calculate the prime factores
	var primes []int32
	var k = 2
	for in.Number > 1 {
		if in.Number%int32(k) == 0 {
			primes = append(primes, int32(k))
			in.Number /= int32(k)
		} else {
			k++
		}
	}

	// Send each prime number to the client
	for _, prime := range primes {
		res := &pb.GetPrimesResponse{
			Prime: prime,
		}
		if err := stream.Send(res); err != nil {
			log.Printf("Error sending prime number: %v\n", err)
			return err
		}
	}

	return nil
}
