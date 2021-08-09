package backend

import (
	"fmt"
	"net/http"

	"github.com/jasonlimantoro/hello-microservice/internal/backend/middleware"
	"github.com/jasonlimantoro/hello-microservice/internal/checkout"
	"github.com/jasonlimantoro/hello-microservice/internal/helloworld"
	"github.com/jasonlimantoro/hello-microservice/pkg"
)

const (
	CONN_PORT = "8000"
	CONN_HOST = "localhost"
	CONN_ADDR = CONN_HOST + ":" + CONN_PORT
)

func Route() {
	helloworldClientCtx := helloworld.NewClientContext()
	checkoutClientContext := checkout.NewClientContext()
	http.HandleFunc("/", middleware.WithLogging(func(w http.ResponseWriter, r *http.Request) {
		msg := pkg.GetHello("backend")
		fmt.Fprintln(w, msg)
	}))

	http.HandleFunc("/hello",
		middleware.WithLogging(
			CreateHandlerHello(
				HelloServiceContext{
					helloworldClient: helloworldClientCtx.Client,
				},
			),
		),
	)
	http.HandleFunc("/checkout",
		middleware.WithLogging(
			CreateHandlerCreateCheckout(
				CreateCheckoutServiceContext{
					checkoutClient: checkoutClientContext.Client,
				},
			),
		),
	)
}
