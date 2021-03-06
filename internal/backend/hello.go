package backend

import (
	"context"
	"net/http"

	"github.com/jasonlimantoro/hello-microservice/pkg"
)

func CreateHandlerHello(svcCtx HelloServiceContext) http.HandlerFunc {
	var req helloRequest
	return pkg.CreateHandler(
		svcCtx,
		&req,
		pkg.DecodeJSON,
		pkg.EncodeJSON,
		Hello,
	)
}
func Hello(ctx context.Context, request interface{}) interface{} {
	var req = request.(*helloRequest)
	svcCtx := pkg.SvcCtxFromCtx(ctx).(HelloServiceContext)
	return helloResponse{Message: svcCtx.helloworldClient.SayHello(ctx, req.Name)}
}

type HelloServiceContext struct {
	helloworldClient helloworldClienter
}
type helloworldClienter interface {
	SayHello(ctx context.Context, name string) string
}

type helloRequest struct {
	Name string `json:"name" validate:"required"`
}

type helloResponse struct {
	Message string `json:"message"`
}
