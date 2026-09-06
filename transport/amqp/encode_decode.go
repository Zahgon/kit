package amqp

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type DecodeRequestFunc func(context.Context, *amqp.Delivery) (request interface{}, err error)

type EncodeRequestFunc func(context.Context, *amqp.Publishing, interface{}) error

type EncodeResponseFunc func(context.Context, *amqp.Publishing, interface{}) error

type DecodeResponseFunc func(context.Context, *amqp.Delivery) (response interface{}, err error)
