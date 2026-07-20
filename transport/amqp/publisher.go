package amqp

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	amqp "github.com/rabbitmq/amqp091-go"
)

const maxCorrelationIdLength = 255

type Publisher struct {
	ch        Channel
	q         *amqp.Queue
	enc       EncodeRequestFunc
	dec       DecodeResponseFunc
	before    []RequestFunc
	after     []PublisherResponseFunc
	deliverer Deliverer
	timeout   time.Duration
}

func NewPublisher(
	ch Channel,
	q *amqp.Queue,
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

func PublisherDeliverer(deliverer Deliverer) PublisherOption {
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

type Deliverer func(
	context.Context,
	Publisher,
	*amqp.Publishing,
) (*amqp.Delivery, error)

func DefaultDeliverer(
	ctx context.Context,
	p Publisher,
	pub *amqp.Publishing,
) (*amqp.Delivery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SendAndForgetDeliverer(
	ctx context.Context,
	p Publisher,
	pub *amqp.Publishing,
) (*amqp.Delivery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
