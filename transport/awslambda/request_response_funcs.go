package awslambda

import (
	"context"
)

type HandlerRequestFunc func(ctx context.Context, payload []byte) context.Context

type HandlerResponseFunc func(ctx context.Context, response interface{}) context.Context

type HandlerFinalizerFunc func(ctx context.Context, resp []byte, err error)
