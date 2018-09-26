package ledlib

import (
	"ledlib/util"
	"time"
)

type ObjectScrolledBitmap struct {
	timer  *Timer
	offset int
	z      int
	image  util.Image2D
}

func NewObjectScrolledBitmap(path string, z int, updateRate time.Duration) LedManagedObject {
	obj := ObjectScrolledBitmap{}
	obj.timer = NewTimer(updateRate)
	obj.image = util.NewImage2D(path)
	obj.z = z
	return &obj
}

func (o *ObjectScrolledBitmap) Draw(cube util.Image3D) {
	o.offset = int(o.timer.GetPastCount())
	for x := 0; x < LedWidth; x++ {
		for y := 0; y < LedHeight; y++ {
			c := o.image.GetAt((x+o.offset)%o.image.GetWidth(), y)
			if !c.IsOff() {
				cube.SetAt(x, y, o.z, c)
			}
		}
	}
}
func (o *ObjectScrolledBitmap) IsExpired() bool {
	return false
}
