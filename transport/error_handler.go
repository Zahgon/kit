package transport

import (
	"context"

	"github.com/go-kit/log"
)

type ErrorHandler interface {
	Handle(ctx context.Context, err error)
}

type LogErrorHandler struct {
	logger log.Logger
}

func NewLogErrorHandler(logger log.Logger) *LogErrorHandler { _ = "STUB: not implemented"; return nil }

func (h *LogErrorHandler) Handle(ctx context.Context, err error) { _ = "STUB: not implemented"; return }

type ErrorHandlerFunc func(ctx context.Context, err error)

func (f ErrorHandlerFunc) Handle(ctx context.Context, err error) { _ = "STUB: not implemented"; return }
