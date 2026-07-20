package pcp

import (
	"github.com/performancecopilot/speed/v4"

	"github.com/go-kit/kit/metrics"
)

type Reporter struct {
	c *speed.PCPClient
}

func NewReporter(appname string) (*Reporter, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Reporter) Start() { _ = "STUB: not implemented"; return }

func (r *Reporter) Stop() { _ = "STUB: not implemented"; return }

type Counter struct {
	c speed.Counter
}

func (r *Reporter) NewCounter(name string, desc ...string) (*Counter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Counter) With(labelValues ...string) metrics.Counter {
	_ = "STUB: not implemented"
	return *new(metrics.Counter)
}

func (c *Counter) Add(delta float64) { _ = "STUB: not implemented"; return }

type Gauge struct {
	g speed.Gauge
}

func (r *Reporter) NewGauge(name string, desc ...string) (*Gauge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Gauge) With(labelValues ...string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (g *Gauge) Set(value float64) { _ = "STUB: not implemented"; return }

func (g *Gauge) Add(delta float64) { _ = "STUB: not implemented"; return }

type Histogram struct {
	h speed.Histogram
}

func (r *Reporter) NewHistogram(name string, min, max int64, unit speed.MetricUnit, desc ...string) (*Histogram, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *Histogram) With(labelValues ...string) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (h *Histogram) Observe(value float64) { _ = "STUB: not implemented"; return }

func (h *Histogram) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (h *Histogram) Percentile(p float64) int64 { _ = "STUB: not implemented"; return 0 }
