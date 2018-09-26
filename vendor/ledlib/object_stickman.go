package ledlib

import "ledlib/util"

func NewObjectStickman() LedObject {
	paths := []string{
		"",
		"./asset/image/stickman/stickman2.png",
		"./asset/image/stickman/stickman3.png",
		"./asset/image/stickman/stickman4.png",
		"./asset/image/stickman/stickman4.png",
		"./asset/image/stickman/stickman5.png",
		"./asset/image/stickman/stickman2.png",
	}
	for i, path := range paths {
		if path == "" {
			continue
		}
		paths[i] = util.GetFullPath(path)
	}
	return NewObjectBitmap(paths)
}
