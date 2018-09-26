package ledlib

import (
	"ledlib/util"
	"time"
)

const MaxAdd = 300

type FilterRolling struct {
	canvas LedCanvas
	add    int
	timer  *Timer
	cube   util.Image3D
}

func NewFilterRolling(canvas LedCanvas) LedCanvas {
	return &FilterRolling{canvas, 0, NewTimer(100 * time.Millisecond), NewLedImage3D()}
}

func (f *FilterRolling) Show(c util.Image3D, param LedCanvasParam) {
	f.cube.Clear()
	if f.timer.IsPast() {
		f.add = (f.add + 1)
	}
	c.ConcurrentForEach(func(x, y, z int, c util.Color32) {
		f.cube.SetAt(x, (y+f.add)%LedHeight, z, c)
	})
	f.canvas.Show(f.cube, param)
}
