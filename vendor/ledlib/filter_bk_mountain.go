package ledlib

import (
	"ledlib/util"
	"time"
)

type FilterBkMountain struct {
	filterObjects     *FilterObjects
	filterObjectsSnow *FilterObjects
}

func NewFilterBkMountain(canvas LedCanvas) LedCanvas {
	filter := FilterBkMountain{}
	filter.filterObjects = NewFilterObjects(canvas)
	filter.filterObjects.Append(NewObjectScrolledBitmap(
		util.GetFullPath("./asset/image/mountain/mountain1.png"), 6, 300*time.Millisecond))
	filter.filterObjects.Append(NewObjectScrolledBitmap(
		util.GetFullPath("./asset/image/mountain/mountain2.png"), 7, 300*time.Millisecond))
	filter.filterObjectsSnow = NewFilterObjects(canvas)
	filter.filterObjectsSnow.Append(NewObjectScrolledBitmap(
		util.GetFullPath("./asset/image/mountain/mountain1-s.png"), 6, 300*time.Millisecond))
	filter.filterObjectsSnow.Append(NewObjectScrolledBitmap(
		util.GetFullPath("./asset/image/mountain/mountain2-s.png"), 7, 300*time.Millisecond))

	return &filter
}

func (f *FilterBkMountain) Show(c util.Image3D, param LedCanvasParam) {
	if param.HasEffect("filter-snows") {
		f.filterObjectsSnow.Show(c, param)
	} else {
		f.filterObjects.Show(c, param)
	}

}
