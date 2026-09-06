package term

import (
	"io"

	"github.com/go-kit/log"
)

func NewLogger(w io.Writer, newLogger func(io.Writer) log.Logger, color func(keyvals ...interface{}) FgBgColor) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

func IsTerminal(w io.Writer) bool { _ = "STUB: not implemented"; return false }
