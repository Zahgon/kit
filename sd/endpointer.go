package sd

import (
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

type Endpointer interface {
	Endpoints() ([]endpoint.Endpoint, error)
}

type FixedEndpointer []endpoint.Endpoint

func (s FixedEndpointer) Endpoints() ([]endpoint.Endpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewEndpointer(src Instancer, f Factory, logger log.Logger, options ...EndpointerOption) *DefaultEndpointer {
	_ = "STUB: not implemented"
	return nil
}

type EndpointerOption func(*endpointerOptions)

func InvalidateOnError(timeout time.Duration) EndpointerOption {
	_ = "STUB: not implemented"
	return *new(EndpointerOption)
}

type endpointerOptions struct {
	invalidateOnError bool
	invalidateTimeout time.Duration
}

type DefaultEndpointer struct {
	cache     *endpointCache
	instancer Instancer
	ch        chan Event
}

func (de *DefaultEndpointer) receive() { _ = "STUB: not implemented"; return }

func (de *DefaultEndpointer) Close() { _ = "STUB: not implemented"; return }

func (de *DefaultEndpointer) Endpoints() ([]endpoint.Endpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
