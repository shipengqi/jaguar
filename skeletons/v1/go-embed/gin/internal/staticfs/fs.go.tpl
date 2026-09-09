package staticfs

import (
	"embed"
	"io/fs"
)

//go:embed all:web
var embeddedFS embed.FS

// FS returns a sub-filesystem rooted at the "web" directory.
func FS() fs.FS {
	sub, err := fs.Sub(embeddedFS, "web")
	if err != nil {
		panic(err)
	}
	return sub
}
