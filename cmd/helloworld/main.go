package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	helloworldgrpc "github.com/jasonlimantoro/hello-microservice/internal/helloworld/gRPC/server"
	pb "github.com/jasonlimantoro/hello-microservice/internal/helloworld/proto"
	helloworldapi "github.com/jasonlimantoro/hello-microservice/internal/helloworld/rest"
	"google.golang.org/grpc"
)

func main() {
	HTTP_PORT := strconv.Itoa(8000)
	TCP_PORT := strconv.Itoa(3000)
	helloworldapi.Route()
	fmt.Println("Listening http on port " + HTTP_PORT)
	fmt.Println("Listening tcp on port " + TCP_PORT)
	lis, _ := net.Listen("tcp", ":"+TCP_PORT)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &helloworldgrpc.Server{})
	go func() {
		log.Fatal(http.ListenAndServe(":"+HTTP_PORT, nil))
	}()
	log.Fatal(s.Serve(lis))
}
