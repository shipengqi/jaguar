package server

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MountSPA mounts the embedded static frontend and serves it as a SPA.
// All unmatched routes fall back to index.html to support client-side routing.
func MountSPA(r *gin.Engine, staticFS fs.FS) {
	fileServer := http.FileServer(http.FS(staticFS))

	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if _, err := staticFS.Open(path); err != nil {
			// SPA fallback
			c.Request.URL.Path = "/"
		}
		fileServer.ServeHTTP(c.Writer, c.Request)
	})
}
