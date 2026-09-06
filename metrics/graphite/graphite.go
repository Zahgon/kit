package graphite

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/generic"
	"github.com/go-kit/log"
)

type Graphite struct {
	mtx        sync.RWMutex
	prefix     string
	counters   map[string]*Counter
	gauges     map[string]*Gauge
	histograms map[string]*Histogram
	logger     log.Logger
}

func New(prefix string, logger log.Logger) *Graphite { _ = "STUB: not implemented"; return nil }

func (g *Graphite) NewCounter(name string) *Counter { _ = "STUB: not implemented"; return nil }

func (g *Graphite) NewGauge(name string) *Gauge { _ = "STUB: not implemented"; return nil }

func (g *Graphite) NewHistogram(name string, buckets int) *Histogram {
	_ = "STUB: not implemented"
	return nil
}

func (g *Graphite) WriteLoop(ctx context.Context, c <-chan time.Time, w io.Writer) {
	_ = "STUB: not implemented"
	return
}

func (g *Graphite) SendLoop(ctx context.Context, c <-chan time.Time, network, address string) {
	_ = "STUB: not implemented"
	return
}

func (g *Graphite) WriteTo(w io.Writer) (count int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type Counter struct {
	c *generic.Counter
}

func NewCounter(name string) *Counter { _ = "STUB: not implemented"; return nil }

func (c *Counter) With(...string) metrics.Counter {
	_ = "STUB: not implemented"
	return *new(metrics.Counter)
}

func (c *Counter) Add(delta float64) { _ = "STUB: not implemented"; return }

type Gauge struct {
	g *generic.Gauge
}

func NewGauge(name string) *Gauge { _ = "STUB: not implemented"; return nil }

func (g *Gauge) With(...string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (g *Gauge) Set(value float64) { _ = "STUB: not implemented"; return }

func (g *Gauge) Add(delta float64) { _ = "STUB: not implemented"; return }

type Histogram struct {
	h *generic.Histogram
}

func NewHistogram(name string, buckets int) *Histogram { _ = "STUB: not implemented"; return nil }

func (h *Histogram) With(...string) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (h *Histogram) Observe(value float64) { _ = "STUB: not implemented"; return }
