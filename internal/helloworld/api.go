package helloworld

import (
	"context"
	"log"

	pb "github.com/jasonlimantoro/hello-microservice/internal/helloworld/proto"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("SayHello Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *Server) SayBye(ctx context.Context, in *pb.ByeRequest) (*pb.ByeReply, error) {
	log.Printf("SayBye Received: %v", in.GetName())
	return &pb.ByeReply{Message: "Bye " + in.GetName()}, nil
}
