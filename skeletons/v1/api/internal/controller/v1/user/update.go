package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shipengqi/errors"
	"github.com/shipengqi/log"

	v1 "github.com/jaguar/apiskeleton/pkg/api/apiserver/v1"
	metav1 "github.com/jaguar/apiskeleton/pkg/api/meta/v1"
	"github.com/jaguar/apiskeleton/pkg/code"
	"github.com/jaguar/apiskeleton/pkg/response"
)

// Update updates a user info by the user identifier.
func (c *Controller) Update(ctx *gin.Context) {
	log.Info("update user function called.")

	var r v1.User

	if err := ctx.ShouldBindJSON(&r); err != nil {
		response.Fail(ctx, errors.WithCode(err, code.ErrBind))
		return
	}

	user, err := c.svc.Users().Get(ctx, ctx.Param("name"), metav1.GetOptions{})
	if err != nil {
		response.Fail(ctx, err)
		return
	}

	user.Nickname = r.Nickname
	user.Email = r.Email
	user.Phone = r.Phone
	user.Extend = r.Extend

	// Save changed fields.
	if err = c.svc.Users().Update(ctx, user, metav1.UpdateOptions{}); err != nil {
		response.Fail(ctx, err)
		return
	}

	response.OK(ctx)
}
