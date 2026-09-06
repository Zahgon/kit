package test

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type metaContext string

const (
	correlationID     metaContext = "correlation-id"
	responseHDR       metaContext = "my-response-header"
	responseTRLR      metaContext = "my-response-trailer"
	correlationIDTRLR metaContext = "correlation-id-consumed"
)

func injectCorrelationID(ctx context.Context, md *metadata.MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func displayClientRequestHeaders(ctx context.Context, md *metadata.MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func extractCorrelationID(ctx context.Context, md metadata.MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func displayServerRequestHeaders(ctx context.Context, md metadata.MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func injectResponseHeader(ctx context.Context, md *metadata.MD, _ *metadata.MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func displayServerResponseHeaders(ctx context.Context, md *metadata.MD, _ *metadata.MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func injectResponseTrailer(ctx context.Context, _ *metadata.MD, md *metadata.MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func injectConsumedCorrelationID(ctx context.Context, _ *metadata.MD, md *metadata.MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func displayServerResponseTrailers(ctx context.Context, _ *metadata.MD, md *metadata.MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func displayClientResponseHeaders(ctx context.Context, md metadata.MD, _ metadata.MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func displayClientResponseTrailers(ctx context.Context, _ metadata.MD, md metadata.MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func extractConsumedCorrelationID(ctx context.Context, _ metadata.MD, md metadata.MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func SetCorrelationID(ctx context.Context, v string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetConsumedCorrelationID(ctx context.Context) string { _ = "STUB: not implemented"; return "" }
