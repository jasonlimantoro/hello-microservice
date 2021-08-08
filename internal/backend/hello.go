package backend

import (
	"context"
	"net/http"

	"github.com/jasonlimantoro/hello-microservice/pkg"
)

func createHandleHello(svcCtx HelloServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoded := &helloRequest{}
		pkg.DecodeJSON(r.Body, decoded)
		response := svcCtx.helloworldClient.SayHello(r.Context(), decoded.Name)
		pkg.EncodeJSON(w, helloResponse{Message: response})
	}
}

type HelloServiceContext struct {
	helloworldClient helloworldClienter
}
type helloworldClienter interface {
	SayHello(ctx context.Context, name string) string
}

type helloRequest struct {
	Name string `json:"name"`
}

type helloResponse struct {
	Message string `json:"message"`
}
