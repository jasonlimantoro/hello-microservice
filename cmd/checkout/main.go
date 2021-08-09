package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/jasonlimantoro/hello-microservice/internal/checkout"
	pb "github.com/jasonlimantoro/hello-microservice/internal/checkout/proto"
	"google.golang.org/grpc"
)

func main() {
	TCP_PORT := strconv.Itoa(3001)
	fmt.Println("Listening tcp on port " + TCP_PORT)
	lis, _ := net.Listen("tcp", ":"+TCP_PORT)
	s := grpc.NewServer()
	pb.RegisterCheckoutServiceServer(s, &checkout.Server{})
	log.Fatal(s.Serve(lis))
}
