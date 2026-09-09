package user

import (
	"github.com/gin-gonic/gin"

	metav1 "{{ .App.ModuleName }}/pkg/api/meta/v1"
	"{{ .App.ModuleName }}/pkg/response"
	xlog "{{ .App.ModuleName }}/pkg/xlog"
)

// Get return a user by the user identifier.
func (c *Controller) Get(ctx *gin.Context) {
	xlog.Info("get user function called.")

	user, err := c.svc.Users().Get(ctx, ctx.Param("name"), metav1.GetOptions{})
	if err != nil {
		response.Fail(ctx, err)
		return
	}

	response.OKWithData(ctx, user)
}
