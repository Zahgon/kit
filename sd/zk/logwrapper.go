package zk

import (
	"github.com/go-zookeeper/zk"

	"github.com/go-kit/log"
)

type wrapLogger struct {
	log.Logger
}

func (logger wrapLogger) Printf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func withLogger(logger log.Logger) func(c *zk.Conn) { _ = "STUB: not implemented"; return nil }
