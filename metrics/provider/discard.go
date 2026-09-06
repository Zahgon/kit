package provider

import (
	"github.com/go-kit/kit/metrics"
)

type discardProvider struct{}

func NewDiscardProvider() Provider { _ = "STUB: not implemented"; return *new(Provider) }

func (discardProvider) NewCounter(string) metrics.Counter {
	_ = "STUB: not implemented"
	return *new(metrics.Counter)
}

func (discardProvider) NewGauge(string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (discardProvider) NewHistogram(string, int) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (discardProvider) Stop() { _ = "STUB: not implemented"; return }
