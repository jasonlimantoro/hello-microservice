package main

import (
	"fmt"
	"net/http"
	"strconv"

	mathapi "github.com/jasonlimantoro/hello-microservice/internal/math/rest"
)

func main() {
	PORT := strconv.Itoa(8001)
	mathapi.Route()
	fmt.Println("Listening on port " + PORT)
	http.ListenAndServe(":"+PORT, nil)
}
