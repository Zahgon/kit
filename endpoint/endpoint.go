package endpoint

import (
	"context"
)

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

func Nop(context.Context, interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Middleware func(Endpoint) Endpoint

func Chain(outer Middleware, others ...Middleware) Middleware {
	_ = "STUB: not implemented"
	return *new(Middleware)
}

type Failer interface {
	Failed() error
}
