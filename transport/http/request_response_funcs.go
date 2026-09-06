package http

import (
	"context"
	"net/http"
)

type RequestFunc func(context.Context, *http.Request) context.Context

type ServerResponseFunc func(context.Context, http.ResponseWriter) context.Context

type ClientResponseFunc func(context.Context, *http.Response) context.Context

func SetContentType(contentType string) ServerResponseFunc {
	_ = "STUB: not implemented"
	return *new(ServerResponseFunc)
}

func SetResponseHeader(key, val string) ServerResponseFunc {
	_ = "STUB: not implemented"
	return *new(ServerResponseFunc)
}

func SetRequestHeader(key, val string) RequestFunc {
	_ = "STUB: not implemented"
	return *new(RequestFunc)
}

func PopulateRequestContext(ctx context.Context, r *http.Request) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type contextKey int

const (
	ContextKeyRequestMethod contextKey = iota

	ContextKeyRequestURI

	ContextKeyRequestPath

	ContextKeyRequestProto

	ContextKeyRequestHost

	ContextKeyRequestRemoteAddr

	ContextKeyRequestXForwardedFor

	ContextKeyRequestXForwardedProto

	ContextKeyRequestAuthorization

	ContextKeyRequestReferer

	ContextKeyRequestUserAgent

	ContextKeyRequestXRequestID

	ContextKeyRequestAccept

	ContextKeyResponseHeaders

	ContextKeyResponseSize
)
