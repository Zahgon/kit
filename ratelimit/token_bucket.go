package ratelimit

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

var ErrLimited = errors.New("rate limit exceeded")

type Allower interface {
	Allow() bool
}

func NewErroringLimiter(limit Allower) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}

type Waiter interface {
	Wait(ctx context.Context) error
}

func NewDelayingLimiter(limit Waiter) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}

type AllowerFunc func() bool

func (f AllowerFunc) Allow() bool { _ = "STUB: not implemented"; return false }

type WaiterFunc func(ctx context.Context) error

func (f WaiterFunc) Wait(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
