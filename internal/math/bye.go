package math

import (
	"context"
	"net/http"

	"github.com/jasonlimantoro/hello-microservice/pkg"
)

func createHandleBye(svcCtx ByeServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoded := &byeRequest{}
		pkg.DecodeJSON(r.Body, decoded)
		response := svcCtx.helloworldClient.SayBye(r.Context(), decoded.Name)
		pkg.EncodeJSON(w, byeResponse{Message: response})
	}
}

type ByeServiceContext struct {
	helloworldClient helloworldClienter
}
type helloworldClienter interface {
	SayBye(ctx context.Context, name string) string
}

type byeRequest struct {
	Name string `json:"name"`
}

type byeResponse struct {
	Message string `json:"message"`
}
