package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	helloworldapi "github.com/jasonlimantoro/hello-microservice/internal/helloworld/rest"
)

func main() {
	PORT := strconv.Itoa(8000)
	helloworldapi.Route()
	fmt.Println("Listening on port " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
