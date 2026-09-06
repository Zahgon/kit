package grpc

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	binHdrSuffix = "-bin"
)

type ClientRequestFunc func(context.Context, *metadata.MD) context.Context

type ServerRequestFunc func(context.Context, metadata.MD) context.Context

type ServerResponseFunc func(ctx context.Context, header *metadata.MD, trailer *metadata.MD) context.Context

type ClientResponseFunc func(ctx context.Context, header metadata.MD, trailer metadata.MD) context.Context

func SetRequestHeader(key, val string) ClientRequestFunc {
	_ = "STUB: not implemented"
	return *new(ClientRequestFunc)
}

func SetResponseHeader(key, val string) ServerResponseFunc {
	_ = "STUB: not implemented"
	return *new(ServerResponseFunc)
}

func SetResponseTrailer(key, val string) ServerResponseFunc {
	_ = "STUB: not implemented"
	return *new(ServerResponseFunc)
}

func EncodeKeyValue(key, val string) (string, string) { _ = "STUB: not implemented"; return "", "" }

type contextKey int

const (
	ContextKeyRequestMethod contextKey = iota
)
