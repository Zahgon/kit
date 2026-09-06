package circuitbreaker

import (
	"github.com/streadway/handy/breaker"

	"github.com/go-kit/kit/endpoint"
)

func HandyBreaker(cb breaker.Breaker) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}
