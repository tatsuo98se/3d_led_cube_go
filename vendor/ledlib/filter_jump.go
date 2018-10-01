package ledlib

import (
	"ledlib/util"
	"time"
)

const Gravity = 0.5
const UpdateFreq = 0.08

type FilterJump struct {
	canvas       LedCanvas
	timer        *Timer
	cube         util.Image3D
	currentPower float64
	initialPower float64
}

func NewFilterJump(canvas LedCanvas) LedCanvas {
	f := FilterJump{}
	f.canvas = canvas
	f.timer = NewTimer(50 * time.Millisecond)
	f.cube = NewLedImage3D()
	f.currentPower = getInitialPower()
	f.initialPower = getInitialPower()
	return &f
}

func getInitialPower() float64 {
	return 3
}
func getPower(power float64) float64 {
	return power * power
}

func (f *FilterJump) Show(c util.Image3D, param LedCanvasParam) {
	f.cube.Clear()
	if f.timer.IsPast() {
		f.currentPower -= Gravity
		if f.currentPower < -f.initialPower {
			f.initialPower = getInitialPower()
			f.currentPower = getInitialPower()
		}
	}
	c.ConcurrentForEach(func(x, y, z int, c util.Color32) {
		f.cube.SetAt(x,
			y+util.RoundToInt(getPower(f.currentPower)-getPower(f.initialPower)),
			z,
			c)
	})
	f.canvas.Show(f.cube, param)
}
