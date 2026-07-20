package httprp

import (
	"context"
	"net/http"
	"net/url"
)

type RequestFunc func(context.Context, *http.Request) context.Context

type Server struct {
	proxy        http.Handler
	before       []RequestFunc
	errorEncoder func(w http.ResponseWriter, err error)
}

func NewServer(
	baseURL *url.URL,
	options ...ServerOption,
) *Server {
	_ = "STUB: not implemented"
	return nil
}

type ServerOption func(*Server)

func ServerBefore(before ...RequestFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
