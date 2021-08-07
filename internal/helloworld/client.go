package helloworld

import (
	"log"

	pb "github.com/jasonlimantoro/hello-microservice/internal/helloworld/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:3000"
)

type ServerContext struct {
	Conn   *grpc.ClientConn
	Client pb.GreeterClient
}

func createConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}
	log.Printf("Connected to helloworld at %s", address)
	return conn, err
}

func NewServerContext() *ServerContext {
	conn, _ := createConn()
	ctx := &ServerContext{
		Client: pb.NewGreeterClient(conn),
		Conn:   conn,
	}
	return ctx
}
