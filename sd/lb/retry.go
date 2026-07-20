package lb

import (
	"time"

	"github.com/go-kit/kit/endpoint"
)

type RetryError struct {
	RawErrors []error
	Final     error
}

func (e RetryError) Error() string { _ = "STUB: not implemented"; return "" }

type Callback func(n int, received error) (keepTrying bool, replacement error)

func Retry(max int, timeout time.Duration, b Balancer) endpoint.Endpoint {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint)
}

func maxRetries(max int) Callback { _ = "STUB: not implemented"; return *new(Callback) }

func alwaysRetry(int, error) (keepTrying bool, replacement error) {
	_ = "STUB: not implemented"
	return false, nil
}

func RetryWithCallback(timeout time.Duration, b Balancer, cb Callback) endpoint.Endpoint {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint)
}
