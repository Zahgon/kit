package discard

import "github.com/go-kit/kit/metrics"

type counter struct{}

func NewCounter() metrics.Counter { _ = "STUB: not implemented"; return *new(metrics.Counter) }

func (c counter) With(labelValues ...string) metrics.Counter {
	_ = "STUB: not implemented"
	return *new(metrics.Counter)
}

func (c counter) Add(delta float64) { _ = "STUB: not implemented"; return }

type gauge struct{}

func NewGauge() metrics.Gauge { _ = "STUB: not implemented"; return *new(metrics.Gauge) }

func (g gauge) With(labelValues ...string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (g gauge) Set(value float64) { _ = "STUB: not implemented"; return }

func (g gauge) Add(delta float64) { _ = "STUB: not implemented"; return }

type histogram struct{}

func NewHistogram() metrics.Histogram { _ = "STUB: not implemented"; return *new(metrics.Histogram) }

func (h histogram) With(labelValues ...string) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (h histogram) Observe(value float64) { _ = "STUB: not implemented"; return }
