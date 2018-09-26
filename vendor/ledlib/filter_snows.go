package ledlib

import (
	"ledlib/util"
	"math/rand"
	"time"
)

type ObjectSnow struct {
	timer   *Timer
	x       int
	z       int
	y       float64
	gravity float64
}

func NewObjectSnow() LedManagedObject {
	snow := ObjectSnow{}
	rand.Seed(time.Now().UnixNano())
	snow.timer = NewTimer(100 * time.Millisecond)
	snow.x = rand.Intn(LedWidth)
	snow.z = rand.Intn(LedDepth)
	snow.y = 0
	snow.gravity = (rand.Float64() / 5) + 0.3

	return &snow
}

func (o *ObjectSnow) Draw(cube util.Image3D) {
	if o.timer.IsPast() {
		o.y = o.y + o.gravity
	}

	if o.IsExpired() {
		return
	}
	cube.SetAt(o.x, util.RoundToInt(o.y), o.z, util.NewFromRGB(0xff, 0xff, 0xff))
}
func (o *ObjectSnow) IsExpired() bool {
	if o.y > LedHeight {
		return true
	}
	return false
}

//////////////////////////

type FilterSnows struct {
	filterObjects *FilterObjects
	timer         *Timer
}

func NewFilterSnows(canvas LedCanvas) LedCanvas {
	filter := FilterSnows{}
	filter.timer = NewTimer(1 * time.Second)
	filter.filterObjects = NewFilterObjects(canvas)

	return &filter
}

func (f *FilterSnows) Show(c util.Image3D, param LedCanvasParam) {
	if f.timer.IsPast() {
		f.filterObjects.Append(NewObjectSnow())
	}
	f.filterObjects.Show(c, param)
}
