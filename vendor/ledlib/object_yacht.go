package ledlib

import "ledlib/util"

func NewObjectYacht() LedObject {
	paths := []string{
		"./asset/image/yacht/yacht1.png",
		"./asset/image/yacht/yacht2.png",
		"./asset/image/yacht/yacht3.png",
		"./asset/image/yacht/yacht4.png",
		"./asset/image/yacht/yacht4.png",
		"./asset/image/yacht/yacht3.png",
		"./asset/image/yacht/yacht2.png",
		"./asset/image/yacht/yacht1.png",
	}
	for i, _ := range paths {
		paths[i] = util.GetFullPath(paths[i])
	}
	return NewObjectBitmap(paths)
}
