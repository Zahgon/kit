package influxstatsd

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/generic"
	"github.com/go-kit/kit/metrics/internal/lv"
	"github.com/go-kit/kit/metrics/internal/ratemap"
	"github.com/go-kit/log"
)

type Influxstatsd struct {
	mtx        sync.RWMutex
	prefix     string
	rates      *ratemap.RateMap
	counters   *lv.Space
	gauges     map[string]*gaugeNode
	timings    *lv.Space
	histograms *lv.Space
	logger     log.Logger
	lvs        lv.LabelValues
}

func New(prefix string, logger log.Logger, lvs ...string) *Influxstatsd {
	_ = "STUB: not implemented"
	return nil
}

func (d *Influxstatsd) NewCounter(name string, sampleRate float64) *Counter {
	_ = "STUB: not implemented"
	return nil
}

func (d *Influxstatsd) NewGauge(name string) *Gauge { _ = "STUB: not implemented"; return nil }

func (d *Influxstatsd) NewTiming(name string, sampleRate float64) *Timing {
	_ = "STUB: not implemented"
	return nil
}

func (d *Influxstatsd) NewHistogram(name string, sampleRate float64) *Histogram {
	_ = "STUB: not implemented"
	return nil
}

func (d *Influxstatsd) WriteLoop(ctx context.Context, c <-chan time.Time, w io.Writer) {
	_ = "STUB: not implemented"
	return
}

func (d *Influxstatsd) SendLoop(ctx context.Context, c <-chan time.Time, network, address string) {
	_ = "STUB: not implemented"
	return
}

func (d *Influxstatsd) WriteTo(w io.Writer) (count int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func sum(a []float64) float64 { _ = "STUB: not implemented"; return 0 }

func sampling(r float64) string { _ = "STUB: not implemented"; return "" }

func (d *Influxstatsd) tagValues(labelValues []string) string { _ = "STUB: not implemented"; return "" }

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
	g      *generic.Gauge
	influx *Influxstatsd
	set    int32
}

func (g *Gauge) With(labelValues ...string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (g *Gauge) Set(value float64) { _ = "STUB: not implemented"; return }

func (g *Gauge) Add(delta float64) { _ = "STUB: not implemented"; return }

type Timing struct {
	name string
	lvs  lv.LabelValues
	obs  observeFunc
}

func (t *Timing) With(labelValues ...string) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (t *Timing) Observe(value float64) { _ = "STUB: not implemented"; return }

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

type pair struct{ label, value string }

type gaugeNode struct {
	mtx      sync.RWMutex
	gauge    *Gauge
	children map[pair]*gaugeNode
}

func (n *gaugeNode) addGauge(g *Gauge, lvs lv.LabelValues) *Gauge {
	_ = "STUB: not implemented"
	return nil
}

func (n *gaugeNode) walk(fn func(string, lv.LabelValues, float64) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *Gauge) touch() { _ = "STUB: not implemented"; return }

func (g *Gauge) read() (float64, bool) { _ = "STUB: not implemented"; return 0, false }
