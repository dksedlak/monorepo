package main

import "errors"

type CtxType string

var (
	ErrValueNotFound = errors.New("context doesn't have the value")
	ErrLogNotFound   = errors.New("context doesn't have any logger on it")

	ctxTypeRequestID = CtxType("requestId")
)

type Middleware struct {
	Service string
}
