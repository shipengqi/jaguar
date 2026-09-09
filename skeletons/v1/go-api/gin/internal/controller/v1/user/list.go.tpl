package user

import (
	"github.com/gin-gonic/gin"

	metav1 "{{ .App.ModuleName }}/pkg/api/meta/v1"
	"{{ .App.ModuleName }}/pkg/code"
	"{{ .App.ModuleName }}/pkg/response"
	xerr "{{ .App.ModuleName }}/pkg/xerr"
	xlog "{{ .App.ModuleName }}/pkg/xlog"
)

// List return the users in the storage.
func (c *Controller) List(ctx *gin.Context) {
	xlog.Info("list user function called.")

	var r metav1.ListOptions
	if err := ctx.ShouldBindQuery(&r); err != nil {
		response.Fail(ctx, xerr.WithCode(err, code.ErrBind))
		return
	}

	users, err := c.svc.Users().List(ctx, r)
	if err != nil {
		response.Fail(ctx, err)
		return
	}

	response.OKWithData(ctx, users)
}
