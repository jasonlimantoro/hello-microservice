package backend

import (
	"context"
	"net/http"

	"github.com/jasonlimantoro/hello-microservice/pkg"
	"github.com/mitchellh/mapstructure"
)

func CreateHandlerHello(svcCtx HelloServiceContext) http.HandlerFunc {
	return pkg.CreateHandler(svcCtx, pkg.DecodeJSON, pkg.EncodeJSON, Hello)
}
func Hello(ctx context.Context, request interface{}) interface{} {
	var req helloRequest
	mapstructure.Decode(request, &req)
	svcCtx := ctx.Value(pkg.SvcCtxKey).(HelloServiceContext)
	return helloResponse{Message: svcCtx.helloworldClient.SayHello(ctx, req.Name)}
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
