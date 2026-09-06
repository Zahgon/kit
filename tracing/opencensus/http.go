package opencensus

import (
	kithttp "github.com/go-kit/kit/transport/http"
)

func HTTPClientTrace(options ...TracerOption) kithttp.ClientOption {
	_ = "STUB: not implemented"
	return *new(kithttp.ClientOption)
}

func HTTPServerTrace(options ...TracerOption) kithttp.ServerOption {
	_ = "STUB: not implemented"
	return *new(kithttp.ServerOption)
}
