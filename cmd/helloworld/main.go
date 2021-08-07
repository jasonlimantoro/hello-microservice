package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/jasonlimantoro/hello-microservice/internal/helloworld"
	pb "github.com/jasonlimantoro/hello-microservice/internal/helloworld/proto"
	"google.golang.org/grpc"
)

func main() {
	TCP_PORT := strconv.Itoa(3000)
	fmt.Println("Listening tcp on port " + TCP_PORT)
	lis, _ := net.Listen("tcp", ":"+TCP_PORT)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &helloworld.Server{})
	log.Fatal(s.Serve(lis))
}
