package log

import (
	"io"

	"github.com/go-kit/log"
)

type StdlibWriter = log.StdlibWriter

type StdlibAdapter = log.StdlibAdapter

type StdlibAdapterOption = log.StdlibAdapterOption

func TimestampKey(key string) StdlibAdapterOption {
	_ = "STUB: not implemented"
	return *new(StdlibAdapterOption)
}

func FileKey(key string) StdlibAdapterOption {
	_ = "STUB: not implemented"
	return *new(StdlibAdapterOption)
}

func MessageKey(key string) StdlibAdapterOption {
	_ = "STUB: not implemented"
	return *new(StdlibAdapterOption)
}

func Prefix(prefix string, joinPrefixToMsg bool) StdlibAdapterOption {
	_ = "STUB: not implemented"
	return *new(StdlibAdapterOption)
}

func NewStdlibAdapter(logger Logger, options ...StdlibAdapterOption) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}
