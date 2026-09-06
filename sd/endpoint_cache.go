package sd

import (
	"io"
	"sync"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

type endpointCache struct {
	options            endpointerOptions
	mtx                sync.RWMutex
	factory            Factory
	cache              map[string]endpointCloser
	err                error
	endpoints          []endpoint.Endpoint
	logger             log.Logger
	invalidateDeadline time.Time
	timeNow            func() time.Time
}

type endpointCloser struct {
	endpoint.Endpoint
	io.Closer
}

func newEndpointCache(factory Factory, logger log.Logger, options endpointerOptions) *endpointCache {
	_ = "STUB: not implemented"
	return nil
}

func (c *endpointCache) Update(event Event) { _ = "STUB: not implemented"; return }

func (c *endpointCache) updateCache(instances []string) { _ = "STUB: not implemented"; return }

func (c *endpointCache) Endpoints() ([]endpoint.Endpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
