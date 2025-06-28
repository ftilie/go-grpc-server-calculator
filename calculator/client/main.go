package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ftilie/go-grpc-servers/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	doAdd(c)
	// doSubtract(c)
	// doMultiply(c)
	// doDivide(c)
	// doGetPrimes(c)
	// doGetAverage(c)
	// doGetMaximum(c)
	// doSqrt(c, -100)
}
