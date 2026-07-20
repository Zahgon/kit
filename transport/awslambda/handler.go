package awslambda

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport"
	"github.com/go-kit/log"
)

type Handler struct {
	e            endpoint.Endpoint
	dec          DecodeRequestFunc
	enc          EncodeResponseFunc
	before       []HandlerRequestFunc
	after        []HandlerResponseFunc
	errorEncoder ErrorEncoder
	finalizer    []HandlerFinalizerFunc
	errorHandler transport.ErrorHandler
}

func NewHandler(
	e endpoint.Endpoint,
	dec DecodeRequestFunc,
	enc EncodeResponseFunc,
	options ...HandlerOption,
) *Handler {
	_ = "STUB: not implemented"
	return nil
}

type HandlerOption func(*Handler)

func HandlerBefore(before ...HandlerRequestFunc) HandlerOption {
	_ = "STUB: not implemented"
	return *new(HandlerOption)
}

func HandlerAfter(after ...HandlerResponseFunc) HandlerOption {
	_ = "STUB: not implemented"
	return *new(HandlerOption)
}

func HandlerErrorLogger(logger log.Logger) HandlerOption {
	_ = "STUB: not implemented"
	return *new(HandlerOption)
}

func HandlerErrorHandler(errorHandler transport.ErrorHandler) HandlerOption {
	_ = "STUB: not implemented"
	return *new(HandlerOption)
}

func HandlerErrorEncoder(ee ErrorEncoder) HandlerOption {
	_ = "STUB: not implemented"
	return *new(HandlerOption)
}

func HandlerFinalizer(f ...HandlerFinalizerFunc) HandlerOption {
	_ = "STUB: not implemented"
	return *new(HandlerOption)
}

func DefaultErrorEncoder(ctx context.Context, err error) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *Handler) Invoke(
	ctx context.Context,
	payload []byte,
) (resp []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
