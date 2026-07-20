package opencensus

import (
	"context"

	"go.opencensus.io/trace"
)

type EndpointOptions struct {
	IgnoreBusinessError bool

	Attributes []trace.Attribute

	GetName func(ctx context.Context, name string) string

	GetAttributes func(ctx context.Context) []trace.Attribute
}

type EndpointOption func(*EndpointOptions)

func WithEndpointConfig(options EndpointOptions) EndpointOption {
	_ = "STUB: not implemented"
	return *new(EndpointOption)
}

func WithEndpointAttributes(attrs ...trace.Attribute) EndpointOption {
	_ = "STUB: not implemented"
	return *new(EndpointOption)
}

func WithIgnoreBusinessError(val bool) EndpointOption {
	_ = "STUB: not implemented"
	return *new(EndpointOption)
}

func WithSpanName(fn func(ctx context.Context, name string) string) EndpointOption {
	_ = "STUB: not implemented"
	return *new(EndpointOption)
}

func WithSpanAttributes(fn func(ctx context.Context) []trace.Attribute) EndpointOption {
	_ = "STUB: not implemented"
	return *new(EndpointOption)
}
