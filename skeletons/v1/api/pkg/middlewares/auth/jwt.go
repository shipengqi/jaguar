package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/jaguar/apiskeleton/pkg/identity"
	"github.com/jaguar/apiskeleton/pkg/middlewares"
)

// JWTStrategy defines jwt bearer authentication strategy.
type JWTStrategy struct {
	*identity.Middleware
}

var _ middlewares.AuthStrategy = &JWTStrategy{}

// NewJWTStrategy create jwt bearer strategy with GinJWTMiddleware.
func NewJWTStrategy(im *identity.Middleware) JWTStrategy {
	return JWTStrategy{im}
}

// AuthFunc defines jwt bearer strategy as the gin authentication middleware.
func (j JWTStrategy) AuthFunc() gin.HandlerFunc {
	return j.MiddlewareFunc()
}
