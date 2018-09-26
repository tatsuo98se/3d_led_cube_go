package ledlib

import (
	"ledlib/util"
	"strconv"
)

type LedCanvasParam interface {
	AppendsEffect(effect string)
	HasEffect(effect string) bool
}

type LedCanvasParamImpl struct {
	effects []string
}

func NewLedCanvasParam() LedCanvasParam {
	return &LedCanvasParamImpl{make([]string, 0)}
}
func (l *LedCanvasParamImpl) AppendsEffect(effect string) {
	l.effects = append(l.effects, effect)
}
func (l *LedCanvasParamImpl) HasEffect(effect string) bool {
	for _, e := range l.effects {
		if effect == e {
			return true
		}
	}
	return false
}

type LedCanvas interface {
	Show(c util.Image3D, param LedCanvasParam)
}

/**
DummyLedCanvas for test
*/
type DummyLedCanvas struct {
}

func (canvas *DummyLedCanvas) Show(c util.Image3D, param LedCanvasParam) {
}

type LedCanvasImpl struct {
}

func NewLedCanvas() *LedCanvasImpl {
	return &LedCanvasImpl{}
}

func (canvas *LedCanvasImpl) Show(c util.Image3D, param LedCanvasParam) {
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
