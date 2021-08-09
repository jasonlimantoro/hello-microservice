package math

import (
	"fmt"
	"net/http"

	"github.com/jasonlimantoro/hello-microservice/internal/helloworld"
	hello "github.com/jasonlimantoro/hello-microservice/pkg"
)

const (
	CONN_PORT = "8001"
	CONN_HOST = "localhost"
	CONN_ADDR = CONN_HOST + ":" + CONN_PORT
)

func Route() {
	helloworldClientCtx := helloworld.NewClientContext()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := hello.GetHello("math package")
		fmt.Fprintln(w, msg)
	})

	http.HandleFunc("/bye", createHandleBye(ByeServiceContext{helloworldClient: helloworldClientCtx.Client}))
}
