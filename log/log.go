package log

import (
	"github.com/go-kit/log"
)

type Logger = log.Logger

var ErrMissingValue = log.ErrMissingValue

func With(logger Logger, keyvals ...interface{}) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

func WithPrefix(logger Logger, keyvals ...interface{}) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

func WithSuffix(logger Logger, keyvals ...interface{}) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

type LoggerFunc = log.LoggerFunc
