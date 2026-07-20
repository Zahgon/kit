package metrics

import "time"

type Timer struct {
	h Histogram
	t time.Time
	u time.Duration
}

func NewTimer(h Histogram) *Timer { _ = "STUB: not implemented"; return nil }

func (t *Timer) ObserveDuration() { _ = "STUB: not implemented"; return }

func (t *Timer) Unit(u time.Duration) { _ = "STUB: not implemented"; return }
