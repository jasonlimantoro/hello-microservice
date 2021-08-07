package server

import (
	"context"
	"log"

	pb "github.com/jasonlimantoro/hello-microservice/internal/helloworld/proto"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *Server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello Again " + in.GetName()}, nil
}
