package ledlib

import (
	"ledlib/util"
	"strconv"
)

type LedCanvas interface {
	PreShow()
	Show(c util.CubeImage)
}

/**
DummyLedCanvas for test
*/
type DummyLedCanvas struct {
}

func (canvas *DummyLedCanvas) PreShow() {

}

func (canvas *DummyLedCanvas) Show(c util.CubeImage) {
}

type LedCanvasImpl struct {
}

func NewLedCanvas() *LedCanvasImpl {
	return &LedCanvasImpl{}
}

func (canvas *LedCanvasImpl) PreShow() {

}

func (canvas *LedCanvasImpl) Show(c util.CubeImage) {
	util.ConcurrentEnumXYZ(LedWidth, LedHeight, LedDepth, func(x, y, z int) {
		px := c.GetAt(x, y, z)
		if px != nil && !px.IsOff() {
			GetLed().SetLed(x, y, z, px.Uint32())
		} else {
			GetLed().SetLed(x, y, z, 0)
		}
	})
	GetLed().Show()
}

func Atois(is []string) ([]int, error) {
	r := make([]int, len(is))
	for i, data := range is {
		d, err := strconv.Atoi(data)
		if err != nil {
			return nil, err
		}
		r[i] = d
	}
	return r, nil
}
