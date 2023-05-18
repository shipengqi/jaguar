package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shipengqi/errors"
	"github.com/shipengqi/log"

	metav1 "github.com/jaguar/apiskeleton/pkg/api/meta/v1"
	"github.com/jaguar/apiskeleton/pkg/code"
	"github.com/jaguar/apiskeleton/pkg/response"
)

// List return the users in the storage.
func (c *Controller) List(ctx *gin.Context) {
	log.Info("list user function called.")

	var r metav1.ListOptions
	if err := ctx.ShouldBindQuery(&r); err != nil {
		response.Fail(ctx, errors.WithCode(err, code.ErrBind))
		return
	}

	users, err := c.svc.Users().List(ctx, r)
	if err != nil {
		response.Fail(ctx, err)
		return
	}

	response.OkWithData(ctx, users)
}
