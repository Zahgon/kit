package amqp

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RequestFunc func(context.Context, *amqp.Publishing, *amqp.Delivery) context.Context

type SubscriberResponseFunc func(context.Context,
	*amqp.Delivery,
	Channel,
	*amqp.Publishing,
) context.Context

type PublisherResponseFunc func(context.Context, *amqp.Delivery) context.Context

func SetPublishExchange(publishExchange string) RequestFunc {
	_ = "STUB: not implemented"
	return *new(RequestFunc)
}

func SetPublishKey(publishKey string) RequestFunc {
	_ = "STUB: not implemented"
	return *new(RequestFunc)
}

func SetPublishDeliveryMode(dmode uint8) RequestFunc {
	_ = "STUB: not implemented"
	return *new(RequestFunc)
}

func SetNackSleepDuration(duration time.Duration) RequestFunc {
	_ = "STUB: not implemented"
	return *new(RequestFunc)
}

func SetConsumeAutoAck(autoAck bool) RequestFunc {
	_ = "STUB: not implemented"
	return *new(RequestFunc)
}

func SetConsumeArgs(args amqp.Table) RequestFunc {
	_ = "STUB: not implemented"
	return *new(RequestFunc)
}

func SetContentType(contentType string) RequestFunc {
	_ = "STUB: not implemented"
	return *new(RequestFunc)
}

func SetContentEncoding(contentEncoding string) RequestFunc {
	_ = "STUB: not implemented"
	return *new(RequestFunc)
}

func SetCorrelationID(cid string) RequestFunc { _ = "STUB: not implemented"; return *new(RequestFunc) }

func SetAckAfterEndpoint(multiple bool) SubscriberResponseFunc {
	_ = "STUB: not implemented"
	return *new(SubscriberResponseFunc)
}

func getPublishExchange(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func getPublishKey(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func getNackSleepDuration(ctx context.Context) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func getConsumeAutoAck(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func getConsumeArgs(ctx context.Context) amqp.Table {
	_ = "STUB: not implemented"
	return *new(amqp.Table)
}

type contextKey int

const (
	ContextKeyExchange contextKey = iota

	ContextKeyPublishKey

	ContextKeyNackSleepDuration

	ContextKeyAutoAck

	ContextKeyConsumeArgs
)
