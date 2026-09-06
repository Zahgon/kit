package generic

import (
	"io"
	"sync"

	"github.com/VividCortex/gohistogram"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/internal/lv"
)

type Counter struct {
	bits uint64
	Name string
	lvs  lv.LabelValues
}

func NewCounter(name string) *Counter { _ = "STUB: not implemented"; return nil }

func (c *Counter) With(labelValues ...string) metrics.Counter {
	_ = "STUB: not implemented"
	return *new(metrics.Counter)
}

func (c *Counter) Add(delta float64) { _ = "STUB: not implemented"; return }

func (c *Counter) Value() float64 { _ = "STUB: not implemented"; return 0 }

func (c *Counter) ValueReset() float64 { _ = "STUB: not implemented"; return 0 }

func (c *Counter) LabelValues() []string { _ = "STUB: not implemented"; return nil }

type Gauge struct {
	bits uint64
	Name string
	lvs  lv.LabelValues
}

func NewGauge(name string) *Gauge { _ = "STUB: not implemented"; return nil }

func (g *Gauge) With(labelValues ...string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (g *Gauge) Set(value float64) { _ = "STUB: not implemented"; return }

func (g *Gauge) Add(delta float64) { _ = "STUB: not implemented"; return }

func (g *Gauge) Value() float64 { _ = "STUB: not implemented"; return 0 }

func (g *Gauge) LabelValues() []string { _ = "STUB: not implemented"; return nil }

type Histogram struct {
	Name string
	lvs  lv.LabelValues
	h    *safeHistogram
}

func NewHistogram(name string, buckets int) *Histogram { _ = "STUB: not implemented"; return nil }

func (h *Histogram) With(labelValues ...string) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (h *Histogram) Observe(value float64) { _ = "STUB: not implemented"; return }

func (h *Histogram) Quantile(q float64) float64 { _ = "STUB: not implemented"; return 0 }

func (h *Histogram) LabelValues() []string { _ = "STUB: not implemented"; return nil }

func (h *Histogram) Print(w io.Writer) { _ = "STUB: not implemented"; return }

type safeHistogram struct {
	sync.RWMutex
	gohistogram.Histogram
}

type Bucket struct {
	From, To, Count int64
}

type Quantile struct {
	Quantile int
	Value    int64
}

type SimpleHistogram struct {
	mtx sync.RWMutex
	lvs lv.LabelValues
	avg float64
	n   uint64
}

func NewSimpleHistogram() *SimpleHistogram { _ = "STUB: not implemented"; return nil }

func (h *SimpleHistogram) With(labelValues ...string) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (h *SimpleHistogram) Observe(value float64) { _ = "STUB: not implemented"; return }

func (h *SimpleHistogram) ApproximateMovingAverage() float64 { _ = "STUB: not implemented"; return 0 }

func (h *SimpleHistogram) LabelValues() []string { _ = "STUB: not implemented"; return nil }
