package util

import (
	"go/build"
	"path/filepath"
)

var ProgramPath string = "src/github.com/tatsuo98se/3d_led_cube_go/"

func GetFullPath(relativePath string) string {

	return filepath.Join(build.Default.GOPATH, ProgramPath, relativePath)
}
