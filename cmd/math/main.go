package main

import (
	"fmt"
	"net/http"

	"github.com/jasonlimantoro/hello-microservice/internal/math"
)

func main() {
	math.Route()
	fmt.Println("Listening on " + math.CONN_ADDR)
	http.ListenAndServe(math.CONN_ADDR, nil)
}
