package jwt

import (
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/transport/http"
)

const (
	bearer       string = "bearer"
	bearerFormat string = "Bearer %s"
)

func HTTPToContext() http.RequestFunc { _ = "STUB: not implemented"; return *new(http.RequestFunc) }

func ContextToHTTP() http.RequestFunc { _ = "STUB: not implemented"; return *new(http.RequestFunc) }

func GRPCToContext() grpc.ServerRequestFunc {
	_ = "STUB: not implemented"
	return *new(grpc.ServerRequestFunc)
}

func ContextToGRPC() grpc.ClientRequestFunc {
	_ = "STUB: not implemented"
	return *new(grpc.ClientRequestFunc)
}

func extractTokenFromAuthHeader(val string) (token string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

func generateAuthHeaderFromToken(token string) string { _ = "STUB: not implemented"; return "" }
