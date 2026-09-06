package teststat

import (
	"github.com/go-kit/kit/metrics"
)

func TestCounter(counter metrics.Counter, value func() float64) error {
	_ = "STUB: not implemented"
	return nil
}

func FillCounter(counter metrics.Counter) float64 { _ = "STUB: not implemented"; return 0 }

func TestGauge(gauge metrics.Gauge, value func() []float64) error {
	_ = "STUB: not implemented"
	return nil
}

func TestHistogram(histogram metrics.Histogram, quantiles func() (p50, p90, p95, p99 float64), tolerance float64) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	Count = 12345

	Mean = 500

	Stdev = 25
)

func ExpectedObservationsLessThan(bucket int64) int64 { _ = "STUB: not implemented"; return 0 }

func cmp(want, have, tol float64) bool { _ = "STUB: not implemented"; return false }
