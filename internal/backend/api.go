package backend

import (
	"context"
	"fmt"
	"net/http"
	"time"

	grpcClient "github.com/jasonlimantoro/hello-microservice/internal/helloworld/gRPC"
	pb "github.com/jasonlimantoro/hello-microservice/internal/helloworld/proto"
	hello "github.com/jasonlimantoro/hello-microservice/pkg"
)

func Route() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := hello.GetHello("backend")
		fmt.Fprintln(w, msg)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		conn, _ := grpcClient.Pool.Get(ctx)
		defer conn.Close()
		c := pb.NewGreeterClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, _ := c.SayHello(ctx, &pb.HelloRequest{Name: "Jeff"})
		fmt.Fprintf(w, "Backend says: %s", response.GetMessage())
	})
}
