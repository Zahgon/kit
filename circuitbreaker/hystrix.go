package circuitbreaker

import (
	"github.com/go-kit/kit/endpoint"
)

func Hystrix(commandName string) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}
