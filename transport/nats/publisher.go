package nats

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/nats-io/nats.go"
)

type Publisher struct {
	publisher *nats.Conn
	subject   string
	enc       EncodeRequestFunc
	dec       DecodeResponseFunc
	before    []RequestFunc
	after     []PublisherResponseFunc
	timeout   time.Duration
}

func NewPublisher(
	publisher *nats.Conn,
	subject string,
	enc EncodeRequestFunc,
	dec DecodeResponseFunc,
	options ...PublisherOption,
) *Publisher {
	_ = "STUB: not implemented"
	return nil
}

type PublisherOption func(*Publisher)

func PublisherBefore(before ...RequestFunc) PublisherOption {
	_ = "STUB: not implemented"
	return *new(PublisherOption)
}

func PublisherAfter(after ...PublisherResponseFunc) PublisherOption {
	_ = "STUB: not implemented"
	return *new(PublisherOption)
}

func PublisherTimeout(timeout time.Duration) PublisherOption {
	_ = "STUB: not implemented"
	return *new(PublisherOption)
}

func (p Publisher) Endpoint() endpoint.Endpoint {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint)
}

func EncodeJSONRequest(_ context.Context, msg *nats.Msg, request interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
