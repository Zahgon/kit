package multi

import "github.com/go-kit/kit/metrics"

type Counter []metrics.Counter

func NewCounter(c ...metrics.Counter) Counter { _ = "STUB: not implemented"; return *new(Counter) }

func (c Counter) Add(delta float64) { _ = "STUB: not implemented"; return }

func (c Counter) With(labelValues ...string) metrics.Counter {
	_ = "STUB: not implemented"
	return *new(metrics.Counter)
}

type Gauge []metrics.Gauge

func NewGauge(g ...metrics.Gauge) Gauge { _ = "STUB: not implemented"; return *new(Gauge) }

func (g Gauge) Set(value float64) { _ = "STUB: not implemented"; return }

func (g Gauge) With(labelValues ...string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (g Gauge) Add(delta float64) { _ = "STUB: not implemented"; return }

type Histogram []metrics.Histogram

func NewHistogram(h ...metrics.Histogram) Histogram {
	_ = "STUB: not implemented"
	return *new(Histogram)
}

func (h Histogram) Observe(value float64) { _ = "STUB: not implemented"; return }

func (h Histogram) With(labelValues ...string) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}
