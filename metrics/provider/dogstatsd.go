package provider

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/dogstatsd"
)

type dogstatsdProvider struct {
	d    *dogstatsd.Dogstatsd
	stop func()
}

func NewDogstatsdProvider(d *dogstatsd.Dogstatsd, stop func()) Provider {
	_ = "STUB: not implemented"
	return *new(Provider)
}

func (p *dogstatsdProvider) NewCounter(name string) metrics.Counter {
	_ = "STUB: not implemented"
	return *new(metrics.Counter)
}

func (p *dogstatsdProvider) NewGauge(name string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (p *dogstatsdProvider) NewHistogram(name string, _ int) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (p *dogstatsdProvider) Stop() { _ = "STUB: not implemented"; return }
