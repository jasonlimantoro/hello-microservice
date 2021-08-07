package rest

import (
	"fmt"
	"net/http"

	hello "github.com/jasonlimantoro/hello-microservice/pkg"
)

func Route() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := hello.GetHello("math package")
		fmt.Fprintln(w, msg)
	})
}
