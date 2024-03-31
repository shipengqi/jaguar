package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shipengqi/log"

	metav1 "github.com/jaguar/apiskeleton/pkg/api/meta/v1"
	"github.com/jaguar/apiskeleton/pkg/response"
)

// Get return a user by the user identifier.
func (c *Controller) Get(ctx *gin.Context) {
	log.Info("get user function called.")

	user, err := c.svc.Users().Get(ctx, ctx.Param("name"), metav1.GetOptions{})
	if err != nil {
		response.Fail(ctx, err)
		return
	}

	response.OKWithData(ctx, user)
}
