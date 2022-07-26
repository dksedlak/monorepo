package main

import (
	"context"
	"net/http"
)

func NewMiddleware(service string) *Middleware {
	return &Middleware{
		Service: service,
	}
}

func (md *Middleware) WithRequestID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if _, err := getFromCtx(ctxTypeRequestID, ctx); err != nil {
			r = r.WithContext(ctx)
		}
		h.ServeHTTP(w, r)
	})
}

func getFromCtx(value CtxType, ctx context.Context) (string, error) {
	var result any
	if result = ctx.Value(value); result == nil {
		return "", ErrValueNotFound
	}
	return result.(string), nil
}
