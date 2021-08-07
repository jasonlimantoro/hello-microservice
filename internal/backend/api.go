package backend

import (
	"fmt"
	"net/http"

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
		response, _ := helloworldServerCtx.Client.SayHello(r.Context(), &pb.HelloRequest{Name: "Jeff"})
		fmt.Fprintf(w, "Backend says: %s", response.GetMessage())
	})
}
