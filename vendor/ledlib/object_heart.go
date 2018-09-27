package ledlib

import "ledlib/util"

func NewObjectHeart() LedObject {
	paths := []string{
		"./asset/image/heart/heart1.png",
		"./asset/image/heart/heart2.png",
		"./asset/image/heart/heart3.png",
		"./asset/image/heart/heart4.png",
		"./asset/image/heart/heart4.png",
		"./asset/image/heart/heart3.png",
		"./asset/image/heart/heart2.png",
		"./asset/image/heart/heart1.png",
	}
	for i, _ := range paths {
		paths[i] = util.GetFullPath(paths[i])
	}
	return NewObjectBitmap(paths)
}
