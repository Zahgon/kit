package jwt

import (
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/golang-jwt/jwt/v4"
)

type contextKey string

const (
	JWTContextKey contextKey = "JWTToken"

	JWTTokenContextKey = JWTContextKey

	JWTClaimsContextKey contextKey = "JWTClaims"
)

var (
	ErrTokenContextMissing = errors.New("token up for parsing was not passed through the context")

	ErrTokenInvalid = errors.New("JWT was invalid")

	ErrTokenExpired = errors.New("JWT is expired")

	ErrTokenMalformed = errors.New("JWT is malformed")

	ErrTokenNotActive = errors.New("token is not valid yet")

	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
)

func NewSigner(kid string, key []byte, method jwt.SigningMethod, claims jwt.Claims) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}

type ClaimsFactory func() jwt.Claims

func MapClaimsFactory() jwt.Claims { _ = "STUB: not implemented"; return *new(jwt.Claims) }

func StandardClaimsFactory() jwt.Claims { _ = "STUB: not implemented"; return *new(jwt.Claims) }

func NewParser(keyFunc jwt.Keyfunc, method jwt.SigningMethod, newClaims ClaimsFactory) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}
