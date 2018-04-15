package kubesanityutils

import (
	"os"
	"path"
)

//GetDir gets diretory by name
func GetDir(dirToGet string) string {
	var projectPath = "/src/github.com/zanetworker/go-kubesanity/"

	switch dirToGet {
	case "root":
		return path.Join(os.Getenv("GOPATH") + projectPath)
	}

	return ""
}
