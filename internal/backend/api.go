package backend

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jasonlimantoro/hello-microservice/internal/helloworld"
	pb "github.com/jasonlimantoro/hello-microservice/internal/helloworld/proto"
	hello "github.com/jasonlimantoro/hello-microservice/pkg"
)

func Route() {
	helloworldServerCtx := helloworld.NewServerContext()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := hello.GetHello("backend")
		fmt.Fprintln(w, msg)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		defer helloworldServerCtx.Conn.Close()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		response, _ := helloworldServerCtx.Client.SayHello(ctx, &pb.HelloRequest{Name: "Jeff"})
		fmt.Fprintf(w, "Backend says: %s", response.GetMessage())
	})
}
