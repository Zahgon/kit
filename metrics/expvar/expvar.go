package expvar

import (
	"expvar"
	"sync"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/generic"
)

type Counter struct {
	f *expvar.Float
}

func NewCounter(name string) *Counter { _ = "STUB: not implemented"; return nil }

func (c *Counter) With(labelValues ...string) metrics.Counter {
	_ = "STUB: not implemented"
	return *new(metrics.Counter)
}

func (c *Counter) Add(delta float64) { _ = "STUB: not implemented"; return }

type Gauge struct {
	f *expvar.Float
}

func NewGauge(name string) *Gauge { _ = "STUB: not implemented"; return nil }

func (g *Gauge) With(labelValues ...string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (g *Gauge) Set(value float64) { _ = "STUB: not implemented"; return }

func (g *Gauge) Add(delta float64) { _ = "STUB: not implemented"; return }

type Histogram struct {
	mtx sync.Mutex
	h   *generic.Histogram
	p50 *expvar.Float
	p90 *expvar.Float
	p95 *expvar.Float
	p99 *expvar.Float
}

func NewHistogram(name string, buckets int) *Histogram { _ = "STUB: not implemented"; return nil }

func (h *Histogram) With(labelValues ...string) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (h *Histogram) Observe(value float64) { _ = "STUB: not implemented"; return }
