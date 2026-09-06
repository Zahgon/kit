//go:build !windows && !plan9 && !nacl
// +build !windows,!plan9,!nacl

package syslog

import (
	"io"

	"github.com/go-kit/log"
	"github.com/go-kit/log/syslog"
)

type SyslogWriter = syslog.SyslogWriter

func NewSyslogLogger(w SyslogWriter, newLogger func(io.Writer) log.Logger, options ...Option) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

type Option = syslog.Option

type PrioritySelector = syslog.PrioritySelector

func PrioritySelectorOption(selector PrioritySelector) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
