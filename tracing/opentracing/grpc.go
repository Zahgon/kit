package opentracing

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/metadata"

	"github.com/go-kit/log"
)

func ContextToGRPC(tracer opentracing.Tracer, logger log.Logger) func(ctx context.Context, md *metadata.MD) context.Context {
	_ = "STUB: not implemented"
	return nil
}

func GRPCToContext(tracer opentracing.Tracer, operationName string, logger log.Logger) func(ctx context.Context, md metadata.MD) context.Context {
	_ = "STUB: not implemented"
	return nil
}

type metadataReaderWriter struct {
	*metadata.MD
}

func (w metadataReaderWriter) Set(key, val string) { _ = "STUB: not implemented"; return }

func (w metadataReaderWriter) ForeachKey(handler func(key, val string) error) error {
	_ = "STUB: not implemented"
	return nil
}
