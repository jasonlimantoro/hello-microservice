package backend

import (
	"context"
	"log"
	"net/http"

	"github.com/jasonlimantoro/hello-microservice/pkg"
	"github.com/jasonlimantoro/hello-microservice/pkg/dto"
	"github.com/mitchellh/mapstructure"
)

func CreateHandlerCreateCheckout(svcCtx CreateCheckoutServiceContext) http.HandlerFunc {
	return pkg.CreateHandler(svcCtx, pkg.DecodeJSON, pkg.EncodeJSON, CreateCheckout)
}
func CreateCheckout(ctx context.Context, request interface{}) interface{} {
	var req createCheckoutRequest
	err := mapstructure.Decode(request, &req)
	if err != nil {
		log.Fatal(err.Error())
	}
	svcCtx := pkg.SvcCtxFromCtx(ctx).(CreateCheckoutServiceContext)
	created := svcCtx.checkoutClient.Create(ctx, dto.Checkout{Email: req.Email, Address: req.Address})
	return createCheckoutResponse{Email: created.Email, Address: req.Address}
}

type CreateCheckoutServiceContext struct {
	checkoutClient checkoutClient
}
type checkoutClient interface {
	Create(ctx context.Context, checkout dto.Checkout) dto.Checkout
}

type createCheckoutRequest struct {
	Address string `json:"address"`
	Email   string `json:"email"`
}

type createCheckoutResponse struct {
	Address string `json:"address"`
	Email   string `json:"email"`
}
