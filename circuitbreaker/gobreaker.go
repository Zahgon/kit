package circuitbreaker

import (
	"github.com/sony/gobreaker"

	"github.com/go-kit/kit/endpoint"
)

func Gobreaker(cb *gobreaker.CircuitBreaker) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}
