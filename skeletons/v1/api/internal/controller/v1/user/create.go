package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shipengqi/log"

	v1 "github.com/jaguar/apiskeleton/pkg/api/apiserver/v1"
	metav1 "github.com/jaguar/apiskeleton/pkg/api/meta/v1"
	"github.com/jaguar/apiskeleton/pkg/response"
	"github.com/jaguar/apiskeleton/pkg/secret"
)

// Create add new user to the storage.
func (c *Controller) Create(ctx *gin.Context) {
	log.Info("user create function called.")

	var u v1.User

	if err := ctx.ShouldBindJSON(&u); err != nil {
		response.Fail(ctx, err)
		return
	}

	u.Password, _ = secret.Encrypt(u.Password)
	u.Status = 1
	u.LoginAt = time.Now()

	// Insert the user to the storage.
	if err := c.svc.Users().Create(ctx, &u, metav1.CreateOptions{}); err != nil {
		response.Fail(ctx, err)
		return
	}
	response.OkWithData(ctx, u)
}
