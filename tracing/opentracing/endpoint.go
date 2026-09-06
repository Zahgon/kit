package opentracing

import (
	"github.com/opentracing/opentracing-go"

	"github.com/go-kit/kit/endpoint"
)

func TraceEndpoint(tracer opentracing.Tracer, operationName string, opts ...EndpointOption) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}

func TraceServer(tracer opentracing.Tracer, operationName string, opts ...EndpointOption) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}

func TraceClient(tracer opentracing.Tracer, operationName string, opts ...EndpointOption) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}

func applyTags(span opentracing.Span, tags opentracing.Tags) { _ = "STUB: not implemented"; return }
