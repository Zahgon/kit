package level

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func Error(logger log.Logger) log.Logger { _ = "STUB: not implemented"; return *new(log.Logger) }

func Warn(logger log.Logger) log.Logger { _ = "STUB: not implemented"; return *new(log.Logger) }

func Info(logger log.Logger) log.Logger { _ = "STUB: not implemented"; return *new(log.Logger) }

func Debug(logger log.Logger) log.Logger { _ = "STUB: not implemented"; return *new(log.Logger) }

func NewFilter(next log.Logger, options ...Option) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

type Option = level.Option

func AllowAll() Option { _ = "STUB: not implemented"; return *new(Option) }

func AllowDebug() Option { _ = "STUB: not implemented"; return *new(Option) }

func AllowInfo() Option { _ = "STUB: not implemented"; return *new(Option) }

func AllowWarn() Option { _ = "STUB: not implemented"; return *new(Option) }

func AllowError() Option { _ = "STUB: not implemented"; return *new(Option) }

func AllowNone() Option { _ = "STUB: not implemented"; return *new(Option) }

func ErrNotAllowed(err error) Option { _ = "STUB: not implemented"; return *new(Option) }

func SquelchNoLevel(squelch bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func ErrNoLevel(err error) Option { _ = "STUB: not implemented"; return *new(Option) }

func NewInjector(next log.Logger, lvl Value) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

type Value = level.Value

func Key() interface{} { _ = "STUB: not implemented"; return nil }

func ErrorValue() Value { _ = "STUB: not implemented"; return *new(Value) }

func WarnValue() Value { _ = "STUB: not implemented"; return *new(Value) }

func InfoValue() Value { _ = "STUB: not implemented"; return *new(Value) }

func DebugValue() Value { _ = "STUB: not implemented"; return *new(Value) }
