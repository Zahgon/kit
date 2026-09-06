package provider

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/statsd"
)

type statsdProvider struct {
	s    *statsd.Statsd
	stop func()
}

func NewStatsdProvider(s *statsd.Statsd, stop func()) Provider {
	_ = "STUB: not implemented"
	return *new(Provider)
}

func (p *statsdProvider) NewCounter(name string) metrics.Counter {
	_ = "STUB: not implemented"
	return *new(metrics.Counter)
}

func (p *statsdProvider) NewGauge(name string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (p *statsdProvider) NewHistogram(name string, _ int) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (p *statsdProvider) Stop() { _ = "STUB: not implemented"; return }
