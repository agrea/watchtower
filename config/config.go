package config

import (
	"os"
	"path"
)

var (
	// StaticPath is the name of the directory where static files are stored.
	StaticPath = path.Join(os.Getenv("GOPATH"), "src", "github.com", "agrea", "watchtower", "static")

	// TemplatePath is the name of the directory where templates are stored.
	TemplatePath = path.Join(os.Getenv("GOPATH"), "src", "github.com", "agrea", "watchtower", "templates")
)
