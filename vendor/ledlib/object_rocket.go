package ledlib

import (
	"go/build"
	"path/filepath"
)

func NewObjectRocket() LedObject {
	paths := []string{
		"./asset/image/rocket/rocket1.png",
		"./asset/image/rocket/rocket2.png",
		"./asset/image/rocket/rocket3.png",
		"./asset/image/rocket/rocket4.png",
		"./asset/image/rocket/rocket4.png",
		"./asset/image/rocket/rocket3.png",
		"./asset/image/rocket/rocket2.png",
		"./asset/image/rocket/rocket1.png",
	}
	programPath := "src/github.com/tatsuo98se/3d_led_cube_go/"
	for i, _ := range paths {
		paths[i] = filepath.Join(build.Default.GOPATH, programPath, paths[i])
	}
	return NewObjectBitmap(paths)
}
