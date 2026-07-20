package zipkin

import (
	"net/http"

	"github.com/go-kit/log"
)

type TracerOption func(o *tracerOptions)

func Name(name string) TracerOption { _ = "STUB: not implemented"; return *new(TracerOption) }

func Tags(tags map[string]string) TracerOption {
	_ = "STUB: not implemented"
	return *new(TracerOption)
}

func Logger(logger log.Logger) TracerOption { _ = "STUB: not implemented"; return *new(TracerOption) }

func AllowPropagation(propagate bool) TracerOption {
	_ = "STUB: not implemented"
	return *new(TracerOption)
}

func RequestSampler(sampleFunc func(r *http.Request) bool) TracerOption {
	_ = "STUB: not implemented"
	return *new(TracerOption)
}

type tracerOptions struct {
	tags           map[string]string
	name           string
	logger         log.Logger
	propagate      bool
	requestSampler func(r *http.Request) bool
}
