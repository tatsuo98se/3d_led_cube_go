package ledlib

import "path/filepath"

func NewRocketBitmapObj(canvas ILedCanvas) *BitmapObj {
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
		var err error
		paths[i], err = filepath.Abs(paths[i])
		if err != nil {
			return nil
		}
	}
	return NewBitmapObj(canvas, paths)
}
