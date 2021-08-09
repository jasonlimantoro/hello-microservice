package backend

import (
	"context"
	"net/http"

	"github.com/jasonlimantoro/hello-microservice/pkg"
	"github.com/jasonlimantoro/hello-microservice/pkg/dto"
)

func CreateHandlerCreateCheckout(svcCtx CreateCheckoutServiceContext) http.HandlerFunc {
	var createCheckoutRequest createCheckoutRequest
	return pkg.CreateHandler(svcCtx, &createCheckoutRequest, pkg.DecodeJSON, pkg.EncodeJSON, CreateCheckout)
}
func CreateCheckout(ctx context.Context, request interface{}) interface{} {
	var req = request.(*createCheckoutRequest)
	svcCtx := pkg.SvcCtxFromCtx(ctx).(CreateCheckoutServiceContext)
	created := svcCtx.checkoutClient.Create(ctx, dto.Checkout{Email: req.Email, Address: req.Address})
	return createCheckoutResponse{ID: created.Id, Email: created.Email, Address: created.Address}
}

type CreateCheckoutServiceContext struct {
	checkoutClient checkoutClient
}
type checkoutClient interface {
	Create(ctx context.Context, checkout dto.Checkout) dto.Checkout
}

type createCheckoutRequest struct {
	Address string `json:"address" validate:"required"`
	Email   string `json:"email" validate:"email,required"`
}

type createCheckoutResponse struct {
	ID      uint32 `json:"id"`
	Address string `json:"address"`
	Email   string `json:"email"`
}
