package jsonrpc

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

type requestIDKeyType struct{}

var requestIDKey requestIDKeyType

type Server struct {
	ecm          EndpointCodecMap
	before       []httptransport.RequestFunc
	beforeCodec  []RequestFunc
	after        []httptransport.ServerResponseFunc
	errorEncoder httptransport.ErrorEncoder
	finalizer    httptransport.ServerFinalizerFunc
	logger       log.Logger
}

func NewServer(
	ecm EndpointCodecMap,
	options ...ServerOption,
) *Server {
	_ = "STUB: not implemented"
	return nil
}

type ServerOption func(*Server)

func ServerBefore(before ...httptransport.RequestFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ServerBeforeCodec(beforeCodec ...RequestFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ServerAfter(after ...httptransport.ServerResponseFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ServerErrorEncoder(ee httptransport.ErrorEncoder) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ServerErrorLogger(logger log.Logger) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ServerFinalizer(f httptransport.ServerFinalizerFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func DefaultErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

type ErrorCoder interface {
	ErrorCode() int
}

type interceptingWriter struct {
	http.ResponseWriter
	code int
}

func (w *interceptingWriter) WriteHeader(code int) { _ = "STUB: not implemented"; return }
