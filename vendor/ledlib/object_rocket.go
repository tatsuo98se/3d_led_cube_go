package ledlib

import "ledlib/util"

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
	for i, _ := range paths {
		paths[i] = util.GetFullPath(paths[i])
	}
	return NewObjectBitmap(paths)
}
