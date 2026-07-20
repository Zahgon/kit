package sd

import (
	"io"

	"github.com/go-kit/kit/endpoint"
)

type Factory func(instance string) (endpoint.Endpoint, io.Closer, error)
