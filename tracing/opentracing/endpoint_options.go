package opentracing

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

type EndpointOptions struct {
	IgnoreBusinessError bool

	GetOperationName func(ctx context.Context, name string) string

	Tags opentracing.Tags

	GetTags func(ctx context.Context) opentracing.Tags
}

type EndpointOption func(*EndpointOptions)

func WithOptions(options EndpointOptions) EndpointOption {
	_ = "STUB: not implemented"
	return *new(EndpointOption)
}

func WithIgnoreBusinessError(ignoreBusinessError bool) EndpointOption {
	_ = "STUB: not implemented"
	return *new(EndpointOption)
}

func WithOperationNameFunc(getOperationName func(ctx context.Context, name string) string) EndpointOption {
	_ = "STUB: not implemented"
	return *new(EndpointOption)
}

func WithTags(tags opentracing.Tags) EndpointOption {
	_ = "STUB: not implemented"
	return *new(EndpointOption)
}

func WithTagsFunc(getTags func(ctx context.Context) opentracing.Tags) EndpointOption {
	_ = "STUB: not implemented"
	return *new(EndpointOption)
}
