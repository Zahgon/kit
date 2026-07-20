package amqp

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport"
	"github.com/go-kit/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Subscriber struct {
	e                 endpoint.Endpoint
	dec               DecodeRequestFunc
	enc               EncodeResponseFunc
	before            []RequestFunc
	after             []SubscriberResponseFunc
	responsePublisher ResponsePublisher
	errorEncoder      ErrorEncoder
	errorHandler      transport.ErrorHandler
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

func SubscriberResponsePublisher(rp ResponsePublisher) SubscriberOption {
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

func (s Subscriber) ServeDelivery(ch Channel) func(deliv *amqp.Delivery) {
	_ = "STUB: not implemented"
	return nil
}

func EncodeJSONResponse(
	ctx context.Context,
	pub *amqp.Publishing,
	response interface{},
) error {
	_ = "STUB: not implemented"
	return nil
}

func EncodeNopResponse(
	ctx context.Context,
	pub *amqp.Publishing,
	response interface{},
) error {
	_ = "STUB: not implemented"
	return nil
}

type ResponsePublisher func(
	context.Context,
	*amqp.Delivery,
	Channel,
	*amqp.Publishing,
) error

func DefaultResponsePublisher(
	ctx context.Context,
	deliv *amqp.Delivery,
	ch Channel,
	pub *amqp.Publishing,
) error {
	_ = "STUB: not implemented"
	return nil
}

func NopResponsePublisher(
	ctx context.Context,
	deliv *amqp.Delivery,
	ch Channel,
	pub *amqp.Publishing,
) error {
	_ = "STUB: not implemented"
	return nil
}

type ErrorEncoder func(ctx context.Context,
	err error, deliv *amqp.Delivery, ch Channel, pub *amqp.Publishing)

func DefaultErrorEncoder(ctx context.Context,
	err error, deliv *amqp.Delivery, ch Channel, pub *amqp.Publishing) {
	_ = "STUB: not implemented"
	return
}

func SingleNackRequeueErrorEncoder(ctx context.Context,
	err error, deliv *amqp.Delivery, ch Channel, pub *amqp.Publishing) {
	_ = "STUB: not implemented"
	return
}

func ReplyErrorEncoder(
	ctx context.Context,
	err error,
	deliv *amqp.Delivery,
	ch Channel,
	pub *amqp.Publishing,
) {
	_ = "STUB: not implemented"
	return
}

func ReplyAndAckErrorEncoder(ctx context.Context, err error, deliv *amqp.Delivery, ch Channel, pub *amqp.Publishing) {
	_ = "STUB: not implemented"
	return
}

type DefaultErrorResponse struct {
	Error string `json:"err"`
}

type Channel interface {
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWail bool, args amqp.Table) (<-chan amqp.Delivery, error)
}
