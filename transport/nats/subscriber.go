package nats

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport"
	"github.com/go-kit/log"

	"github.com/nats-io/nats.go"
)

type Subscriber struct {
	e            endpoint.Endpoint
	dec          DecodeRequestFunc
	enc          EncodeResponseFunc
	before       []RequestFunc
	after        []SubscriberResponseFunc
	errorEncoder ErrorEncoder
	finalizer    []SubscriberFinalizerFunc
	errorHandler transport.ErrorHandler
}

func NewSubscriber(
	e endpoint.Endpoint,
	dec DecodeRequestFunc,
	enc EncodeResponseFunc,
	options ...SubscriberOption,
) *Subscriber {
	_ = "STUB: not implemented"
	return nil
}

type SubscriberOption func(*Subscriber)

func SubscriberBefore(before ...RequestFunc) SubscriberOption {
	_ = "STUB: not implemented"
	return *new(SubscriberOption)
}

func SubscriberAfter(after ...SubscriberResponseFunc) SubscriberOption {
	_ = "STUB: not implemented"
	return *new(SubscriberOption)
}

func SubscriberErrorEncoder(ee ErrorEncoder) SubscriberOption {
	_ = "STUB: not implemented"
	return *new(SubscriberOption)
}

func SubscriberErrorLogger(logger log.Logger) SubscriberOption {
	_ = "STUB: not implemented"
	return *new(SubscriberOption)
}

func SubscriberErrorHandler(errorHandler transport.ErrorHandler) SubscriberOption {
	_ = "STUB: not implemented"
	return *new(SubscriberOption)
}

func SubscriberFinalizer(f ...SubscriberFinalizerFunc) SubscriberOption {
	_ = "STUB: not implemented"
	return *new(SubscriberOption)
}

func (s Subscriber) ServeMsg(nc *nats.Conn) func(msg *nats.Msg) {
	_ = "STUB: not implemented"
	return nil
}

type ErrorEncoder func(ctx context.Context, err error, reply string, nc *nats.Conn)

type SubscriberFinalizerFunc func(ctx context.Context, msg *nats.Msg)

func NopRequestDecoder(_ context.Context, _ *nats.Msg) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeJSONResponse(_ context.Context, reply string, nc *nats.Conn, response interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func DefaultErrorEncoder(_ context.Context, err error, reply string, nc *nats.Conn) {
	_ = "STUB: not implemented"
	return
}
