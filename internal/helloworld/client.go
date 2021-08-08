package helloworld

import (
	"context"
	"log"

	pb "github.com/jasonlimantoro/hello-microservice/internal/helloworld/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:3000"
)

type ClientContext struct {
	Conn   *grpc.ClientConn
	Client Client
}

type Client struct {
	pb.GreeterClient
}

func (cl Client) SayHello(ctx context.Context, name string) string {
	response, _ := cl.GreeterClient.SayHello(ctx, &pb.HelloRequest{Name: name})
	return response.GetMessage()
}

func createConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}
	log.Printf("Connected to helloworld at %s", address)
	return conn, err
}

func NewClientContext() *ClientContext {
	conn, _ := createConn()
	cl := Client{
		GreeterClient: pb.NewGreeterClient(conn),
	}
	ctx := &ClientContext{
		Client: cl,
		Conn:   conn,
	}
	return ctx
}
