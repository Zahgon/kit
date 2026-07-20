package nats

import (
	"context"

	"github.com/nats-io/nats.go"
)

type RequestFunc func(context.Context, *nats.Msg) context.Context

type SubscriberResponseFunc func(context.Context, *nats.Conn) context.Context

type PublisherResponseFunc func(context.Context, *nats.Msg) context.Context
