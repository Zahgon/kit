package log

import (
	"time"

	"github.com/go-kit/log"
)

type Valuer = log.Valuer

func Timestamp(t func() time.Time) Valuer { _ = "STUB: not implemented"; return *new(Valuer) }

func TimestampFormat(t func() time.Time, layout string) Valuer {
	_ = "STUB: not implemented"
	return *new(Valuer)
}

func Caller(depth int) Valuer { _ = "STUB: not implemented"; return *new(Valuer) }

var (
	DefaultTimestamp = log.DefaultTimestamp

	DefaultTimestampUTC = log.DefaultTimestampUTC

	DefaultCaller = log.DefaultCaller
)
