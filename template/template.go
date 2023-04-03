package template

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:dist
var content embed.FS

// All returns the content of the all directory.
func Dist() http.Handler {
	// Get the subdirectory of the content directory.
	distSub, err := fs.Sub(content, "dist")
	// Check if an error occurred.
	if err != nil {
		log.Fatal(err)
	}
	// Return the file server for the subdirectory.
	return http.FileServer(http.FS(distSub))
}
