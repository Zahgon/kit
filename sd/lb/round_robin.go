package lb

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
)

func NewRoundRobin(s sd.Endpointer) Balancer { _ = "STUB: not implemented"; return *new(Balancer) }

type roundRobin struct {
	s sd.Endpointer
	c uint64
}

func (rr *roundRobin) Endpoint() (endpoint.Endpoint, error) {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint), nil
}
