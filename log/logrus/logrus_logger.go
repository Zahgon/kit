package logrus

import (
	"errors"

	"github.com/go-kit/log"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	field logrus.FieldLogger
	level logrus.Level
}

type Option func(*Logger)

var errMissingValue = errors.New("(MISSING)")

func NewLogger(logger logrus.FieldLogger, options ...Option) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

func WithLevel(level logrus.Level) Option { _ = "STUB: not implemented"; return *new(Option) }

func (l Logger) Log(keyvals ...interface{}) error { _ = "STUB: not implemented"; return nil }
