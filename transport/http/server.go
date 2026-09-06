package http

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport"
	"github.com/go-kit/log"
)

type Server struct {
	e            endpoint.Endpoint
	dec          DecodeRequestFunc
	enc          EncodeResponseFunc
	before       []RequestFunc
	after        []ServerResponseFunc
	errorEncoder ErrorEncoder
	finalizer    []ServerFinalizerFunc
	errorHandler transport.ErrorHandler
}

func NewServer(
	e endpoint.Endpoint,
	dec DecodeRequestFunc,
	enc EncodeResponseFunc,
	options ...ServerOption,
) *Server {
	_ = "STUB: not implemented"
	return nil
}

type ServerOption func(*Server)

func ServerBefore(before ...RequestFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ServerAfter(after ...ServerResponseFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ServerErrorEncoder(ee ErrorEncoder) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ServerErrorLogger(logger log.Logger) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ServerErrorHandler(errorHandler transport.ErrorHandler) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ServerFinalizer(f ...ServerFinalizerFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

type ErrorEncoder func(ctx context.Context, err error, w http.ResponseWriter)

type ServerFinalizerFunc func(ctx context.Context, code int, r *http.Request)

func NopRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeJSONResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func DefaultErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

type StatusCoder interface {
	StatusCode() int
}

type Headerer interface {
	Headers() http.Header
}
