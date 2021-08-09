package pkg

import (
	"context"
	"io"
	"net/http"
)

type Endpoint = func(ctx context.Context, request interface{}) (response interface{})

type svcCtxKeyType string

const svcCtxKey svcCtxKeyType = "svcCtx"

func SvcCtxFromCtx(ctx context.Context) interface{} {
	return ctx.Value(svcCtxKey)
}

func CreateHandler(
	svcCtx interface{},
	requestObjectPtr interface{},
	decode func(io.Reader, interface{}) error,
	encode func(io.Writer, interface{}) error,
	endpoint Endpoint,
) http.HandlerFunc {
	type ErrorResponse struct {
		Error string `json:"error"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		decode(r.Body, requestObjectPtr)
		newContext := context.WithValue(r.Context(), svcCtxKey, svcCtx)

		w.Header().Add("Content-Type", "application/json")

		if err := Validate(newContext, requestObjectPtr); err != nil {
			errorResponse := ErrorResponse{Error: err.Error()}
			encode(w, errorResponse)
			return
		}
		rtn := endpoint(newContext, requestObjectPtr)
		encode(w, rtn)
	}
}
