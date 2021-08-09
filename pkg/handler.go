package pkg

import (
	"context"
	"io"
	"net/http"
)

type Endpoint = func(ctx context.Context, request interface{}) (response interface{})

type SvcCtxKeyType string

const SvcCtxKey SvcCtxKeyType = "svcCtx"

func CreateHandler(
	svcCtx interface{},
	decode func(io.Reader, interface{}) error,
	encode func(io.Writer, interface{}) error,
	endpoint Endpoint,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var decoded interface{}
		decode(r.Body, &decoded)
		newContext := context.WithValue(r.Context(), SvcCtxKey, svcCtx)
		rtn := endpoint(newContext, decoded)
		encode(w, rtn)
	}
}
