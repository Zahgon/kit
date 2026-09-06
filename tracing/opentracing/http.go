package opentracing

import (
	opentracing "github.com/opentracing/opentracing-go"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

func ContextToHTTP(tracer opentracing.Tracer, logger log.Logger) kithttp.RequestFunc {
	_ = "STUB: not implemented"
	return *new(kithttp.RequestFunc)
}

func HTTPToContext(tracer opentracing.Tracer, operationName string, logger log.Logger) kithttp.RequestFunc {
	_ = "STUB: not implemented"
	return *new(kithttp.RequestFunc)
}
