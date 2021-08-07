package grpc

import (
	"log"
	"time"

	grpcpool "github.com/processout/grpc-go-pool"
	"google.golang.org/grpc"
)

var address = "localhost:3000"
var factory = func() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}
	log.Printf("Connected to employee at %s", address)
	return conn, err
}
var Pool, _ = grpcpool.New(factory, 5, 5, time.Second)
