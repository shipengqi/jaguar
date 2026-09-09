package internal

import (
	embedserver "{{ .App.ModuleName }}/internal/server"
	"{{ .App.ModuleName }}/internal/staticfs"

	"github.com/gin-gonic/gin"
)

func mountFrontend(g *gin.Engine) {
	embedserver.MountSPA(g, staticfs.FS())
}
