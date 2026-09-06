package provider

import (
	"github.com/go-kit/kit/metrics"
)

type expvarProvider struct{}

func NewExpvarProvider() Provider { _ = "STUB: not implemented"; return *new(Provider) }

func (p expvarProvider) NewCounter(name string) metrics.Counter {
	_ = "STUB: not implemented"
	return *new(metrics.Counter)
}

func (p expvarProvider) NewGauge(name string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (p expvarProvider) NewHistogram(name string, buckets int) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (p expvarProvider) Stop() { _ = "STUB: not implemented"; return }
