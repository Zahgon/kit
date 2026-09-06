package nats

import (
	"context"

	"github.com/nats-io/nats.go"
)

type DecodeRequestFunc func(context.Context, *nats.Msg) (request interface{}, err error)

type EncodeRequestFunc func(context.Context, *nats.Msg, interface{}) error

type EncodeResponseFunc func(context.Context, string, *nats.Conn, interface{}) error

type DecodeResponseFunc func(context.Context, *nats.Msg) (response interface{}, err error)
