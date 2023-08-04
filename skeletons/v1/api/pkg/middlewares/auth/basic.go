package auth

import (
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jaguar/apiskeleton/pkg/middlewares"
	"github.com/jaguar/apiskeleton/pkg/response"
	"github.com/shipengqi/errors"
)

var _ middlewares.AuthStrategy = &BasicStrategy{}

// BasicStrategy defines Basic authentication strategy.
type BasicStrategy struct {
	compare func(username string, password string) bool
}

// NewBasicStrategy create basic strategy with compare function.
func NewBasicStrategy(compare func(username string, password string) bool) BasicStrategy {
	return BasicStrategy{
		compare: compare,
	}
}

// AuthFunc defines basic strategy as the gin authentication middleware.
func (b BasicStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			response.Send(c, nil, errors.New("Authorization header format is wrong."))
			c.Abort()
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !b.compare(pair[0], pair[1]) {
			response.Send(c, nil, errors.New("Authorization header format is wrong."))
			c.Abort()

			return
		}

		c.Set(middlewares.UsernameKey, pair[0])

		c.Next()
	}
}
