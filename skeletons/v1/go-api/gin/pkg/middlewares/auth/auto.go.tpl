package auth

import (
	"strings"

	"github.com/gin-gonic/gin"

	"{{ .App.ModuleName }}/pkg/middlewares"
	"{{ .App.ModuleName }}/pkg/response"
	xerr "{{ .App.ModuleName }}/pkg/xerr"
)

const authHeaderCount = 2

// AutoStrategy defines authentication strategy which can automatically choose between Basic and Bearer
// according `Authorization` header.
type AutoStrategy struct {
	basic middlewares.AuthStrategy
	jwt   middlewares.AuthStrategy
}

var _ middlewares.AuthStrategy = &AutoStrategy{}

// NewAutoStrategy create auto strategy with basic strategy and jwt strategy.
func NewAutoStrategy(basic, jwt middlewares.AuthStrategy) AutoStrategy {
	return AutoStrategy{
		basic: basic,
		jwt:   jwt,
	}
}

// AuthFunc defines auto strategy as the gin authentication middleware.
func (a AutoStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		operator := middlewares.AuthOperator{}
		authHeader := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(authHeader) != authHeaderCount {
			response.Fail(c, xerr.New("Authorization header format is wrong."))
			c.Abort()
			return
		}

		switch authHeader[0] {
		case "Basic":
			operator.SetStrategy(a.basic)
		case "Bearer":
			operator.SetStrategy(a.jwt)
		default:
			response.Fail(c, xerr.New("unrecognized Authorization header."))
			c.Abort()
			return
		}

		operator.AuthFunc()(c)

		c.Next()
	}
}
