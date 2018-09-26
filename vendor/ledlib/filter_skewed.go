package ledlib

import (
	"ledlib/util"
	"math"
	"time"
)

const T = 20

type FilterSkewed struct {
	canvas LedCanvas
	cube   util.Image3D
	timer  *Timer
	yt     float64
	zt     float64
	ys     float64
	zs     float64
	yc     float64
	zc     float64
}

func NewFilterSkewed(canvas LedCanvas) LedCanvas {
	return &FilterSkewed{canvas, NewLedImage3D(), NewTimer(50 * time.Millisecond),
		0, 0, 0, 0, 0, 0}
}

func (f *FilterSkewed) Show(c util.Image3D, param LedCanvasParam) {

	if f.timer.IsPast() {
		f.cube.Clear()
		f.yt += 0.02 * 0.15
		f.zt += 0.02 * 0.15
		f.ys = math.Sin(f.yt * 3.14 * T)
		f.yc = math.Cos(f.yt * 3.14 * T)
		f.zs = math.Sin(f.zt * 3.14 * T)
		f.zc = math.Cos(f.zt * 3.14 * T)
	}

	dx := float64(LedWidth / 2.0)
	dy := float64(LedHeight / 4.0 * 3)
	dz := float64(LedDepth / 2.0)

	c.ConcurrentForEach(func(x, y, z int, c util.Color32) {
		xx, yy, zz := ((float64(x)-dx)*f.yc+(float64(z)-dz)*f.ys)+dx, float64(y), (-(float64(x)-dx)*f.ys+(float64(z)-dz)*f.yc)+dz

		xx, yy = ((xx-dx)*f.zc+(yy-dy)*f.zs)+dx, (-(xx-dx)*f.zs+(yy-dy)*f.zc)+dy
		f.cube.SetAt(int(math.Round(xx)), int(math.Round(yy)), int(math.Round(zz)), c)
	})
	f.canvas.Show(f.cube, param)
}
