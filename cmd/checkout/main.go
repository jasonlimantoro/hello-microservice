package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jasonlimantoro/hello-microservice/internal/checkout"
	pb "github.com/jasonlimantoro/hello-microservice/internal/checkout/proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Listening on " + checkout.CONN_ADDR)
	lis, _ := net.Listen("tcp", ":"+checkout.CONN_PORT)
	s := grpc.NewServer()
	pb.RegisterCheckoutServiceServer(s, &checkout.Server{})
	log.Fatal(s.Serve(lis))
}
