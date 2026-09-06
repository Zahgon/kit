package provider

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/influx"
)

type influxProvider struct {
	in   *influx.Influx
	stop func()
}

func NewInfluxProvider(in *influx.Influx, stop func()) Provider {
	_ = "STUB: not implemented"
	return *new(Provider)
}

func (p *influxProvider) NewCounter(name string) metrics.Counter {
	_ = "STUB: not implemented"
	return *new(metrics.Counter)
}

func (p *influxProvider) NewGauge(name string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (p *influxProvider) NewHistogram(name string, buckets int) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (p *influxProvider) Stop() { _ = "STUB: not implemented"; return }
