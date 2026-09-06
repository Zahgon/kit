package cloudwatch2

import (
	"context"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/internal/lv"
	"github.com/go-kit/log"
)

const (
	maxConcurrentRequests = 20
)

type CloudWatchAPI interface {
	PutMetricData(ctx context.Context, params *cloudwatch.PutMetricDataInput, optFns ...func(*cloudwatch.Options)) (*cloudwatch.PutMetricDataOutput, error)
}

type CloudWatch struct {
	mtx                   sync.RWMutex
	sem                   chan struct{}
	namespace             string
	svc                   CloudWatchAPI
	counters              *lv.Space
	logger                log.Logger
	numConcurrentRequests int
}

type Option func(*CloudWatch)

func WithLogger(logger log.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithConcurrentRequests(n int) Option { _ = "STUB: not implemented"; return *new(Option) }

func New(namespace string, svc CloudWatchAPI, options ...Option) *CloudWatch {
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

var zero = float64(0.0)

var zeros = types.StatisticSet{
	Maximum:     &zero,
	Minimum:     &zero,
	Sum:         &zero,
	SampleCount: &zero,
}

func stats(a []float64) *types.StatisticSet { _ = "STUB: not implemented"; return nil }

func makeDimensions(labelValues ...string) []types.Dimension { _ = "STUB: not implemented"; return nil }

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
