package lb

import (
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type Balancer interface {
	Endpoint() (endpoint.Endpoint, error)
}

var ErrNoEndpoints = errors.New("no endpoints available")
