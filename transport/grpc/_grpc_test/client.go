package test

import (
	"context"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/endpoint"
)

type clientBinding struct {
	test endpoint.Endpoint
}

func (c *clientBinding) Test(ctx context.Context, a string, b int64) (context.Context, string, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), "", nil
}

func NewClient(cc *grpc.ClientConn) Service { _ = "STUB: not implemented"; return *new(Service) }
