package checkout

import (
	"context"

	pb "github.com/jasonlimantoro/hello-microservice/internal/checkout/proto"
)

type Server struct {
	pb.UnimplementedCheckoutServiceServer
}

var checkouts []pb.Checkout

func (s *Server) Create(ctx context.Context, in *pb.CreateCheckoutRequest) (*pb.Checkout, error) {
	newCheckout := &pb.Checkout{
		Id:      uint32(len(checkouts) + 1),
		Email:   in.GetEmail(),
		Address: in.GetAddress(),
	}
	checkouts = append(checkouts, *newCheckout)
	return newCheckout, nil
}
