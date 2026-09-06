package log

import (
	"io"

	"github.com/go-kit/log"
)

type SwapLogger = log.SwapLogger

func NewSyncWriter(w io.Writer) io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func NewSyncLogger(logger Logger) Logger { _ = "STUB: not implemented"; return *new(Logger) }
