package ledlib

import "fmt"

const MaxAdd = 300

type LedRollingFilter struct {
	canvas ILedCanvas
	add    int
	timer  *Timer
}

func NewLedRollingFilter(canvas ILedCanvas) ILedCanvas {
	return &LedRollingFilter{canvas, 0, NewTimer(100)}
}

func (f *LedRollingFilter) PreDraw() {
	f.canvas.PreDraw()
	if f.timer.IsPast() {
		f.add = (f.add + 1)
		fmt.Printf("add:%d\n", f.add)
	}
}

func (f *LedRollingFilter) SetAt(x, y, z int, c Color32) {

	f.canvas.SetAt(x, (y+f.add)%LedHeight, z, c)
}

func (f *LedRollingFilter) Show() {
	f.canvas.Show()
}
