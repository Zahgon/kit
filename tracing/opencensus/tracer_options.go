package opencensus

import (
	"go.opencensus.io/plugin/ochttp/propagation/b3"
	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"
)

var defaultHTTPPropagate propagation.HTTPFormat = &b3.HTTPFormat{}

type TracerOption func(o *TracerOptions)

func WithTracerConfig(options TracerOptions) TracerOption {
	_ = "STUB: not implemented"
	return *new(TracerOption)
}

func WithSampler(sampler trace.Sampler) TracerOption {
	_ = "STUB: not implemented"
	return *new(TracerOption)
}

func WithName(name string) TracerOption { _ = "STUB: not implemented"; return *new(TracerOption) }

func IsPublic(isPublic bool) TracerOption { _ = "STUB: not implemented"; return *new(TracerOption) }

func WithHTTPPropagation(p propagation.HTTPFormat) TracerOption {
	_ = "STUB: not implemented"
	return *new(TracerOption)
}

type TracerOptions struct {
	Sampler       trace.Sampler
	Name          string
	Public        bool
	HTTPPropagate propagation.HTTPFormat
}
