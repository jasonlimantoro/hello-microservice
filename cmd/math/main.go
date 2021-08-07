package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jasonlimantoro/hello-microservice/internal/math"
)

func main() {
	PORT := strconv.Itoa(8001)
	math.Route()
	fmt.Println("Listening on port " + PORT)
	http.ListenAndServe(":"+PORT, nil)
}
