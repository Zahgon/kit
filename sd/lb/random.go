package lb

import (
	"math/rand"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
)

func NewRandom(s sd.Endpointer, seed int64) Balancer {
	_ = "STUB: not implemented"
	return *new(Balancer)
}

type random struct {
	s sd.Endpointer
	r *rand.Rand
}

func (r *random) Endpoint() (endpoint.Endpoint, error) {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint), nil
}
