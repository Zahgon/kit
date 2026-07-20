package casbin

import (
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type contextKey string

const (
	CasbinModelContextKey contextKey = "CasbinModel"

	CasbinPolicyContextKey contextKey = "CasbinPolicy"

	CasbinEnforcerContextKey contextKey = "CasbinEnforcer"
)

var (
	ErrModelContextMissing = errors.New("CasbinModel is required in context")

	ErrPolicyContextMissing = errors.New("CasbinPolicy is required in context")

	ErrUnauthorized = errors.New("Unauthorized Access")
)

func NewEnforcer(
	subject string, object interface{}, action string,
) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}
