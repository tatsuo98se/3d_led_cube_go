package ledlib

import (
	"ledlib/util"
	"time"
)

type FilterGrass struct {
	filterObjects *FilterObjects
}

func NewFilterGrass(canvas LedCanvas) LedCanvas {
	filter := FilterGrass{}
	filter.filterObjects = NewFilterObjects(canvas)

	duration := 100 * time.Millisecond
	filter.filterObjects.Append(NewObjectScrolledBitmap(
		util.GetFullPath("./asset/image/grass/grass1.png"), 0, duration))
	filter.filterObjects.Append(NewObjectScrolledBitmap(
		util.GetFullPath("./asset/image/grass/grass2.png"), 1, duration))
	filter.filterObjects.Append(NewObjectScrolledBitmap(
		util.GetFullPath("./asset/image/grass/grass1.png"), 2, duration))
	filter.filterObjects.Append(NewObjectScrolledBitmap(
		util.GetFullPath("./asset/image/grass/grass2.png"), 3, duration))
	filter.filterObjects.Append(NewObjectScrolledBitmap(
		util.GetFullPath("./asset/image/grass/grass1.png"), 4, duration))
	filter.filterObjects.Append(NewObjectScrolledBitmap(
		util.GetFullPath("./asset/image/grass/grass3.png"), 5, duration))

	return &filter
}

func (f *FilterGrass) Show(c util.Image3D, param LedCanvasParam) {
	f.filterObjects.Show(c, param)
}
