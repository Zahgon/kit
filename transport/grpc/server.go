package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport"
	"github.com/go-kit/log"
)

type Handler interface {
	ServeGRPC(ctx context.Context, request interface{}) (context.Context, interface{}, error)
}

type Server struct {
	e            endpoint.Endpoint
	dec          DecodeRequestFunc
	enc          EncodeResponseFunc
	before       []ServerRequestFunc
	after        []ServerResponseFunc
	finalizer    []ServerFinalizerFunc
	errorHandler transport.ErrorHandler
}

func NewServer(
	e endpoint.Endpoint,
	dec DecodeRequestFunc,
	enc EncodeResponseFunc,
	options ...ServerOption,
) *Server {
	_ = "STUB: not implemented"
	return nil
}

type ServerOption func(*Server)

func ServerBefore(before ...ServerRequestFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ServerAfter(after ...ServerResponseFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ServerErrorLogger(logger log.Logger) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ServerErrorHandler(errorHandler transport.ErrorHandler) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ServerFinalizer(f ...ServerFinalizerFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func (s Server) ServeGRPC(ctx context.Context, req interface{}) (retctx context.Context, resp interface{}, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

type ServerFinalizerFunc func(ctx context.Context, err error)

func Interceptor(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
