package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shipengqi/errors"
	"github.com/shipengqi/golib/cryptoutil/secret"
	"github.com/shipengqi/log"

	metav1 "github.com/jaguar/apiskeleton/pkg/api/meta/v1"
	"github.com/jaguar/apiskeleton/pkg/code"
	"github.com/jaguar/apiskeleton/pkg/response"
)

// ChangePasswordRequest defines the ChangePasswordRequest data format.
type ChangePasswordRequest struct {
	// Old password.
	// Required: true
	OldPassword string `json:"old_password" binding:"required"`

	// New password.
	// Required: true
	NewPassword string `json:"new_password" binding:"required"`
}

// ChangePassword change the user's password by the user identifier.
func (c *Controller) ChangePassword(ctx *gin.Context) {
	log.Info("update user function called.")

	var r ChangePasswordRequest

	if err := ctx.ShouldBindJSON(&r); err != nil {
		response.Fail(ctx, errors.WithCode(err, code.ErrBind))
		return
	}

	user, err := c.svc.Users().Get(ctx, ctx.Param("name"), metav1.GetOptions{})
	if err != nil {
		response.Fail(ctx, err)
		return
	}

	if err = secret.Compare(user.Password, r.OldPassword); err != nil {
		response.Fail(ctx, errors.WithCode(err, code.ErrPasswordIncorrect))
		return
	}

	// Todo update password
	response.OK(ctx)
}
