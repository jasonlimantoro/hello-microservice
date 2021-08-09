package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jasonlimantoro/hello-microservice/internal/helloworld"
	pb "github.com/jasonlimantoro/hello-microservice/internal/helloworld/proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Listening on " + helloworld.CONN_ADDR)
	lis, _ := net.Listen("tcp", ":"+helloworld.CONN_PORT)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &helloworld.Server{})
	log.Fatal(s.Serve(lis))
}
