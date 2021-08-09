package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jasonlimantoro/hello-microservice/internal/backend"
)

func main() {
	backend.Route()
	fmt.Println("Listening on " + backend.CONN_ADDR)
	log.Fatal(http.ListenAndServe(backend.CONN_ADDR, nil))
}
