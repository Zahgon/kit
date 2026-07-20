package provider

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/graphite"
)

type graphiteProvider struct {
	g    *graphite.Graphite
	stop func()
}

func NewGraphiteProvider(g *graphite.Graphite, stop func()) Provider {
	_ = "STUB: not implemented"
	return *new(Provider)
}

func (p *graphiteProvider) NewCounter(name string) metrics.Counter {
	_ = "STUB: not implemented"
	return *new(metrics.Counter)
}

func (p *graphiteProvider) NewGauge(name string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (p *graphiteProvider) NewHistogram(name string, buckets int) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (p *graphiteProvider) Stop() { _ = "STUB: not implemented"; return }
