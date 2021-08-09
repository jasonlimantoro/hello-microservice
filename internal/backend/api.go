package backend

import (
	"fmt"
	"net/http"

	"github.com/jasonlimantoro/hello-microservice/internal/checkout"
	"github.com/jasonlimantoro/hello-microservice/internal/helloworld"
	"github.com/jasonlimantoro/hello-microservice/pkg"
)

func Route() {
	helloworldClientCtx := helloworld.NewClientContext()
	checkoutClientContext := checkout.NewClientContext()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := pkg.GetHello("backend")
		fmt.Fprintln(w, msg)
	})

	http.HandleFunc("/hello", CreateHandlerHello(HelloServiceContext{helloworldClient: helloworldClientCtx.Client}))
	http.HandleFunc("/checkout", CreateHandlerCreateCheckout(
		CreateCheckoutServiceContext{checkoutClient: checkoutClientContext.Client},
	))
}
