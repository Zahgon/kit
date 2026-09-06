package provider

import (
	"github.com/go-kit/kit/metrics"
)

type prometheusProvider struct {
	namespace string
	subsystem string
}

func NewPrometheusProvider(namespace, subsystem string) Provider {
	_ = "STUB: not implemented"
	return *new(Provider)
}

func (p *prometheusProvider) NewCounter(name string) metrics.Counter {
	_ = "STUB: not implemented"
	return *new(metrics.Counter)
}

func (p *prometheusProvider) NewGauge(name string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (p *prometheusProvider) NewHistogram(name string, _ int) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (p *prometheusProvider) Stop() { _ = "STUB: not implemented"; return }
