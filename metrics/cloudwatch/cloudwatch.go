package cloudwatch

import (
	"context"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/internal/lv"
	"github.com/go-kit/log"
)

const (
	maxConcurrentRequests = 20
	maxValuesInABatch     = 150
)

type CloudWatch struct {
	mtx                   sync.RWMutex
	sem                   chan struct{}
	namespace             string
	svc                   cloudwatchiface.CloudWatchAPI
	counters              *lv.Space
	gauges                *lv.Space
	histograms            *lv.Space
	percentiles           []float64
	logger                log.Logger
	numConcurrentRequests int
}

type Option func(*CloudWatch)

func WithLogger(logger log.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPercentiles(percentiles ...float64) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithConcurrentRequests(n int) Option { _ = "STUB: not implemented"; return *new(Option) }

func New(namespace string, svc cloudwatchiface.CloudWatchAPI, options ...Option) *CloudWatch {
	_ = "STUB: not implemented"
	return nil
}

func (cw *CloudWatch) NewCounter(name string) metrics.Counter {
	_ = "STUB: not implemented"
	return *new(metrics.Counter)
}

func (cw *CloudWatch) NewGauge(name string) metrics.Gauge {
	_ = "STUB: not implemented"
	return *new(metrics.Gauge)
}

func (cw *CloudWatch) NewHistogram(name string) metrics.Histogram {
	_ = "STUB: not implemented"
	return *new(metrics.Histogram)
}

func (cw *CloudWatch) WriteLoop(ctx context.Context, c <-chan time.Time) {
	_ = "STUB: not implemented"
	return
}

func (cw *CloudWatch) Send() error { _ = "STUB: not implemented"; return nil }

func sum(a []float64) float64 { _ = "STUB: not implemented"; return 0 }

func min(a, b int) int { _ = "STUB: not implemented"; return 0 }

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

func makeDimensions(labelValues ...string) []*cloudwatch.Dimension {
	_ = "STUB: not implemented"
	return nil
}
