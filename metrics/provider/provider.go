package provider

import (
	"github.com/go-kit/kit/metrics"
)

type Provider interface {
	NewCounter(name string) metrics.Counter
	NewGauge(name string) metrics.Gauge
	NewHistogram(name string, buckets int) metrics.Histogram
	Stop()
}
