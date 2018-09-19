package ledlib

import (
	"ledlib/util"
)

const MaxAdd = 300

type LedRollingFilter struct {
	canvas LedCanvas
	add    int
	timer  *Timer
	cube   util.CubeImage
}

func NewLedRollingFilter(canvas LedCanvas) LedCanvas {
	return &LedRollingFilter{canvas, 0, NewTimer(100), NewLedCubeImage()}
}

func (f *LedRollingFilter) PreShow() {
	f.cube.Clear()
	f.canvas.PreShow()
	if f.timer.IsPast() {
		f.add = (f.add + 1)
	}
}

func (f *LedRollingFilter) Show(c util.CubeImage) {
	c.ConcurrentForEach(func(x, y, z int, c util.Color32) {
		f.cube.SetAt(x, (y+f.add)%LedHeight, z, c)
	})
	f.canvas.Show(f.cube)
}
