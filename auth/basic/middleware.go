package basic

import (
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type AuthError struct {
	Realm string
}

func (AuthError) StatusCode() int { _ = "STUB: not implemented"; return 0 }

func (AuthError) Error() string { _ = "STUB: not implemented"; return "" }

func (e AuthError) Headers() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func parseBasicAuth(auth string) (username, password []byte, ok bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

func toHashSlice(s []byte) []byte { _ = "STUB: not implemented"; return nil }

func AuthMiddleware(requiredUser, requiredPassword, realm string) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}
