package influx

import (
	"context"
	"time"

	influxdb "github.com/influxdata/influxdb1-client/v2"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/internal/lv"
	"github.com/go-kit/log"
)

type Influx struct {
	counters   *lv.Space
	gauges     *lv.Space
	histograms *lv.Space
	tags       map[string]string
	conf       influxdb.BatchPointsConfig
	logger     log.Logger
}

func New(tags map[string]string, conf influxdb.BatchPointsConfig, logger log.Logger) *Influx {
	_ = "STUB: not implemented"
	return nil
}

func (in *Influx) NewCounter(name string) *Counter { _ = "STUB: not implemented"; return nil }

func (in *Influx) NewGauge(name string) *Gauge { _ = "STUB: not implemented"; return nil }

func (in *Influx) NewHistogram(name string) *Histogram { _ = "STUB: not implemented"; return nil }

type BatchPointsWriter interface {
	Write(influxdb.BatchPoints) error
}

func (in *Influx) WriteLoop(ctx context.Context, c <-chan time.Time, w BatchPointsWriter) {
	_ = "STUB: not implemented"
	return
}

func (in *Influx) WriteTo(w BatchPointsWriter) (err error) { _ = "STUB: not implemented"; return nil }

func mergeTags(tags map[string]string, labelValues []string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func sum(a []float64) float64 { _ = "STUB: not implemented"; return 0 }

func last(a []float64) float64 { _ = "STUB: not implemented"; return 0 }

type observeFunc func(name string, lvs lv.LabelValues, value float64)

type Counter struct {
	name string
	lvs  lv.LabelValues
	obs  observeFunc
}

func (c *Counter) With(labelValues ...string) metrics.Counter {
	_ = "STUB: not implemented"
	return *new(metrics.Counter)
}

func (c *Counter) Add(delta float64) { _ = "STUB: not implemented"; return }

type Gauge struct {
	name string
	lvs  lv.LabelValues
	obs  observeFunc
	add  observeFunc
}

func (g *Gauge) With(labelValues ...string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (g *Gauge) Set(value float64) { _ = "STUB: not implemented"; return }

func (g *Gauge) Add(delta float64) { _ = "STUB: not implemented"; return }

type Histogram struct {
	name string
	lvs  lv.LabelValues
	obs  observeFunc
}

func (h *Histogram) With(labelValues ...string) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (h *Histogram) Observe(value float64) { _ = "STUB: not implemented"; return }
