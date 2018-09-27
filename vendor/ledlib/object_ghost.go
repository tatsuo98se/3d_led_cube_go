package ledlib

import "ledlib/util"

func NewObjectGhost() LedObject {
	paths := []string{
		"./asset/image/ghost/ghost1.png",
		"./asset/image/ghost/ghost2.png",
		"./asset/image/ghost/ghost3.png",
		"./asset/image/ghost/ghost4.png",
		"./asset/image/ghost/ghost4.png",
		"./asset/image/ghost/ghost3.png",
		"./asset/image/ghost/ghost2.png",
		"./asset/image/ghost/ghost5.png",
	}
	for i, _ := range paths {
		paths[i] = util.GetFullPath(paths[i])
	}
	return NewObjectBitmap(paths)
}
