package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jasonlimantoro/hello-microservice/internal/backend"
)

func main() {
	HTTP_PORT := strconv.Itoa(8000)
	backend.Route()
	fmt.Println("Listening http on port " + HTTP_PORT)
	log.Fatal(http.ListenAndServe(":"+HTTP_PORT, nil))
}
