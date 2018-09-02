package ledlib

/*
#cgo LDFLAGS: -lledlib
#include "./../../lib/led.h"
*/
import "C"
import (
	"ledlib/util"
	"strconv"
)

const LedHeight = 32
const LedDepth = 8
const LedWidth = 16

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
	util.EnumXYZ(LedWidth, LedHeight, LedDepth, func(x, y, z int) {
		px := c.GetAt(x, y, z)
		if px != nil && !px.IsOff() {
			C.SetLed(C.int(x), C.int(y), C.int(z), C.int(px.Uint32()))
		} else {
			C.SetLed(C.int(x), C.int(y), C.int(z), C.int(0))
		}
	})
	C.Show()
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
