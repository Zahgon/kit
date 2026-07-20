package teststat

import (
	"github.com/go-kit/kit/metrics"
)

func PopulateNormalHistogram(h metrics.Histogram, seed int) { _ = "STUB: not implemented"; return }

func normalQuantiles() (p50, p90, p95, p99 float64) { _ = "STUB: not implemented"; return 0, 0, 0, 0 }

func nvq(quantile int) float64 { _ = "STUB: not implemented"; return 0 }

func erfinv(y float64) float64 { _ = "STUB: not implemented"; return 0 }
