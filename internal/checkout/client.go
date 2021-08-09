package checkout

import (
	"context"
	"log"

	pb "github.com/jasonlimantoro/hello-microservice/internal/checkout/proto"
	"github.com/jasonlimantoro/hello-microservice/pkg/dto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:3001"
)

type ClientContext struct {
	Conn   *grpc.ClientConn
	Client Client
}

type Client struct {
	pb.CheckoutServiceClient
}

func (cl Client) Create(ctx context.Context, checkout dto.Checkout) dto.Checkout {
	response, _ := cl.CheckoutServiceClient.Create(ctx, &pb.CreateCheckoutRequest{Address: checkout.Address, Email: checkout.Email})
	return dto.Checkout{
		Id:      response.GetId(),
		Address: response.GetAddress(),
		Email:   response.GetEmail(),
	}
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
		CheckoutServiceClient: pb.NewCheckoutServiceClient(conn),
	}
	ctx := &ClientContext{
		Client: cl,
		Conn:   conn,
	}
	return ctx
}
