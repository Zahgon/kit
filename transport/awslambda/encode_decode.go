package awslambda

import (
	"context"
)

type DecodeRequestFunc func(context.Context, []byte) (interface{}, error)

type EncodeResponseFunc func(context.Context, interface{}) ([]byte, error)

type ErrorEncoder func(ctx context.Context, err error) ([]byte, error)
