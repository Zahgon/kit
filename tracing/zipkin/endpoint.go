package zipkin

import (
	"github.com/openzipkin/zipkin-go"

	"github.com/go-kit/kit/endpoint"
)

func TraceEndpoint(tracer *zipkin.Tracer, name string) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}
