package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/internal/lv"
)

type Counter struct {
	cv  *prometheus.CounterVec
	lvs lv.LabelValues
}

func NewCounterFrom(opts prometheus.CounterOpts, labelNames []string) *Counter {
	_ = "STUB: not implemented"
	return nil
}

func NewCounter(cv *prometheus.CounterVec) *Counter { _ = "STUB: not implemented"; return nil }

func (c *Counter) With(labelValues ...string) metrics.Counter {
	_ = "STUB: not implemented"
	return *new(metrics.Counter)
}

func (c *Counter) Add(delta float64) { _ = "STUB: not implemented"; return }

type Gauge struct {
	gv  *prometheus.GaugeVec
	lvs lv.LabelValues
}

func NewGaugeFrom(opts prometheus.GaugeOpts, labelNames []string) *Gauge {
	_ = "STUB: not implemented"
	return nil
}

func NewGauge(gv *prometheus.GaugeVec) *Gauge { _ = "STUB: not implemented"; return nil }

func (g *Gauge) With(labelValues ...string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (g *Gauge) Set(value float64) { _ = "STUB: not implemented"; return }

func (g *Gauge) Add(delta float64) { _ = "STUB: not implemented"; return }

type Summary struct {
	sv  *prometheus.SummaryVec
	lvs lv.LabelValues
}

func NewSummaryFrom(opts prometheus.SummaryOpts, labelNames []string) *Summary {
	_ = "STUB: not implemented"
	return nil
}

func NewSummary(sv *prometheus.SummaryVec) *Summary { _ = "STUB: not implemented"; return nil }

func (s *Summary) With(labelValues ...string) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (s *Summary) Observe(value float64) { _ = "STUB: not implemented"; return }

type Histogram struct {
	hv  *prometheus.HistogramVec
	lvs lv.LabelValues
}

func NewHistogramFrom(opts prometheus.HistogramOpts, labelNames []string) *Histogram {
	_ = "STUB: not implemented"
	return nil
}

func NewHistogram(hv *prometheus.HistogramVec) *Histogram { _ = "STUB: not implemented"; return nil }

func (h *Histogram) With(labelValues ...string) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (h *Histogram) Observe(value float64) { _ = "STUB: not implemented"; return }

func makeLabels(labelValues ...string) prometheus.Labels {
	_ = "STUB: not implemented"
	return *new(prometheus.Labels)
}
