package teststat

import (
	"io"

	"github.com/go-kit/kit/metrics/generic"
)

func SumLines(w io.WriterTo, regex string) func() float64 { _ = "STUB: not implemented"; return nil }

func LastLine(w io.WriterTo, regex string) func() []float64 { _ = "STUB: not implemented"; return nil }

func Quantiles(w io.WriterTo, regex string, buckets int) func() (float64, float64, float64, float64) {
	_ = "STUB: not implemented"
	return nil
}

func stats(w io.WriterTo, regex string, h *generic.Histogram) (sum, final float64) {
	_ = "STUB: not implemented"
	return 0, 0
}
