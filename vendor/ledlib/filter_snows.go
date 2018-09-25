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
	snow.timer = NewTimer(100)
	snow.x = rand.Intn(LedWidth)
	snow.z = rand.Intn(LedDepth)
	snow.y = 0
	snow.gravity = (rand.Float64() / 5) + 0.3

	return &snow
}

func (o *ObjectSnow) DidDetach() {
}

func (o *ObjectSnow) Draw(cube util.CubeImage) {
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
	filter.timer = NewTimer(1000)
	filter.filterObjects = NewFilterObjects(canvas)

	return &filter
}

func (f *FilterSnows) PreShow() {
	f.filterObjects.PreShow()
	if f.timer.IsPast() {
		f.filterObjects.Append(NewObjectSnow())
	}
}

func (f *FilterSnows) Show(c util.CubeImage) {
	f.filterObjects.Show(c)
}
