package statsd

import (
	"context"
	"io"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/internal/lv"
	"github.com/go-kit/kit/metrics/internal/ratemap"
	"github.com/go-kit/log"
)

type Statsd struct {
	prefix string
	rates  *ratemap.RateMap

	counters *lv.Space
	gauges   *lv.Space
	timings  *lv.Space

	logger log.Logger
}

func New(prefix string, logger log.Logger) *Statsd { _ = "STUB: not implemented"; return nil }

func (s *Statsd) NewCounter(name string, sampleRate float64) *Counter {
	_ = "STUB: not implemented"
	return nil
}

func (s *Statsd) NewGauge(name string) *Gauge { _ = "STUB: not implemented"; return nil }

func (s *Statsd) NewTiming(name string, sampleRate float64) *Timing {
	_ = "STUB: not implemented"
	return nil
}

func (s *Statsd) WriteLoop(ctx context.Context, c <-chan time.Time, w io.Writer) {
	_ = "STUB: not implemented"
	return
}

func (s *Statsd) SendLoop(ctx context.Context, c <-chan time.Time, network, address string) {
	_ = "STUB: not implemented"
	return
}

func (s *Statsd) WriteTo(w io.Writer) (count int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func sum(a []float64) float64 { _ = "STUB: not implemented"; return 0 }

func last(a []float64) float64 { _ = "STUB: not implemented"; return 0 }

func sampling(r float64) string { _ = "STUB: not implemented"; return "" }

type observeFunc func(name string, lvs lv.LabelValues, value float64)

type Counter struct {
	name string
	obs  observeFunc
}

func (c *Counter) With(...string) metrics.Counter {
	_ = "STUB: not implemented"
	return *new(metrics.Counter)
}

func (c *Counter) Add(delta float64) { _ = "STUB: not implemented"; return }

type Gauge struct {
	name string
	obs  observeFunc
	add  observeFunc
}

func (g *Gauge) With(...string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (g *Gauge) Set(value float64) { _ = "STUB: not implemented"; return }

func (g *Gauge) Add(delta float64) { _ = "STUB: not implemented"; return }

type Timing struct {
	name string
	obs  observeFunc
}

func (t *Timing) With(...string) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (t *Timing) Observe(value float64) { _ = "STUB: not implemented"; return }
