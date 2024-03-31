package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shipengqi/errors"
	"github.com/shipengqi/log"
)

const (
	SUCCESS = 0
)

// ErrResponse defines the return messages when an error occurred.
// Reference is optional, if there is an error occurred, maybe it is useful to solve the error.
// swagger:model
type ErrResponse struct {
	// Code defines the business error code.
	Code int `json:"code"`

	// Message contains the detail of this message.
	// This message is suitable to be exposed to external
	Message string `json:"message"`

	// Reference returns the reference document which maybe useful to solve this error.
	Reference string `json:"reference,omitempty"`

	Data interface{} `json:"data,omitempty"`
}

// Send write an error or the response data into http response body.
// It uses the errors.ParseCoder to parse any error into errors.Coder
// errors.Coder contains error code, user-safe error message and http status code.
func Send(c *gin.Context, data interface{}, err error) {
	if err != nil {
		log.Errorf("%#+v", err)
		coder := errors.ParseCoder(err)
		c.JSON(coder.HTTPStatus(), ErrResponse{
			Code:      coder.Code(),
			Message:   coder.String(),
			Reference: coder.Reference(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func OK(c *gin.Context) {
	Send(c, map[string]interface{}{}, nil)
}

func OKWithData(c *gin.Context, data interface{}) {
	Send(c, data, nil)
}

func Fail(c *gin.Context, err error) {
	Send(c, map[string]interface{}{}, err)
}
