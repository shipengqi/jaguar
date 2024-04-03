package auth

import (
	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"{{ .App.ModuleName }}/pkg/middlewares"
)

// JWTStrategy defines jwt bearer authentication strategy.
type JWTStrategy struct {
	*ginjwt.GinJWTMiddleware
}

var _ middlewares.AuthStrategy = &JWTStrategy{}

// NewJWTStrategy create jwt bearer strategy with GinJWTMiddleware.
func NewJWTStrategy(jm *ginjwt.GinJWTMiddleware) JWTStrategy {
	return JWTStrategy{jm}
}

// AuthFunc defines jwt bearer strategy as the gin authentication middleware.
func (j JWTStrategy) AuthFunc() gin.HandlerFunc {
	return j.MiddlewareFunc()
}
