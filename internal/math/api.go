package math

import (
	"fmt"
	"net/http"

	"github.com/jasonlimantoro/hello-microservice/internal/helloworld"
	hello "github.com/jasonlimantoro/hello-microservice/pkg"
)

func Route() {
	helloworldClientCtx := helloworld.NewClientContext()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := hello.GetHello("math package")
		fmt.Fprintln(w, msg)
	})

	http.HandleFunc("/bye", createHandleBye(ByeServiceContext{helloworldClient: helloworldClientCtx.Client}))
}
