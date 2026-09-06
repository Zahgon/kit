package zipkin

import (
	zipkin "github.com/openzipkin/zipkin-go"

	kithttp "github.com/go-kit/kit/transport/http"
)

func HTTPClientTrace(tracer *zipkin.Tracer, options ...TracerOption) kithttp.ClientOption {
	_ = "STUB: not implemented"
	return *new(kithttp.ClientOption)
}

func HTTPServerTrace(tracer *zipkin.Tracer, options ...TracerOption) kithttp.ServerOption {
	_ = "STUB: not implemented"
	return *new(kithttp.ServerOption)
}
