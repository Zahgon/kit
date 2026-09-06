package opencensus

import (
	"github.com/go-kit/kit/endpoint"
)

const TraceEndpointDefaultName = "gokit/endpoint"

func TraceEndpoint(name string, options ...EndpointOption) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}
