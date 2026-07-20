package zap

import (
	"github.com/go-kit/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapSugarLogger func(msg string, keysAndValues ...interface{})

func (l zapSugarLogger) Log(kv ...interface{}) error { _ = "STUB: not implemented"; return nil }

func NewZapSugarLogger(logger *zap.Logger, level zapcore.Level) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}
