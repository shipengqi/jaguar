package internal

import (
	"github.com/gin-gonic/gin"

	"{{ .App.ModuleName }}/internal/controller/v1/user"
	"{{ .App.ModuleName }}/internal/store"
	"{{ .App.ModuleName }}/pkg/code"
	"{{ .App.ModuleName }}/pkg/response"
	"{{ .App.ModuleName }}/pkg/xerr"
)

func initRouter(g *gin.Engine) {
	installMiddlewares(g)
	installControllers(g)
	mountFrontend(g)
}

func installMiddlewares(g *gin.Engine) {}

func installControllers(g *gin.Engine) {
	g.NoRoute(func(c *gin.Context) {
		response.Fail(c, xerr.Codef(code.ErrPageNotFound, "Page not found."))
	})

	// v1 handlers, requiring authentication
	storeIns := store.Client()
	v1 := g.Group("/api/v1")
	{
		userv1 := v1.Group("/users")
		{
			userc := user.New(storeIns)

			userv1.POST("", userc.Create)
			userv1.DELETE("", userc.DeleteList)
			userv1.DELETE(":name", userc.Delete)
			userv1.PUT(":name/change-password", userc.ChangePassword)
			userv1.PUT(":name", userc.Update)
			userv1.GET("", userc.List)
			userv1.GET(":name", userc.Get)
		}
	}
}
