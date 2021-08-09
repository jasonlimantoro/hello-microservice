package helloworld

import (
	"context"
	"log"

	pb "github.com/jasonlimantoro/hello-microservice/internal/helloworld/proto"
	"google.golang.org/grpc"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3000"
	CONN_ADDR = CONN_HOST + ":" + CONN_PORT
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

func (cl Client) SayBye(ctx context.Context, name string) string {
	response, _ := cl.GreeterClient.SayBye(ctx, &pb.ByeRequest{Name: name})
	return response.GetMessage()
}

func createConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(CONN_ADDR, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}
	log.Printf("Connected to helloworld at %s", CONN_ADDR)
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
