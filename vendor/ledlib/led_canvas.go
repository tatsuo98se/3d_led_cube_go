package ledlib

import (
	"ledlib/util"
	"strconv"
)

type ILedCanvas interface {
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

type LedCanvas struct {
}

func NewLedCanvas() ILedCanvas {
	return &LedCanvas{}
}

func (canvas *LedCanvas) PreShow() {

}

func (canvas *LedCanvas) Show(c util.CubeImage) {
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
