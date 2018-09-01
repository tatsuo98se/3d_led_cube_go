package ledlib

/*
#cgo LDFLAGS: -lledlib
#include "./../lib/led.h"
*/
import "C"
import (
	"image/color"
	"strconv"
)

const LedHeight = 32
const LedDepth = 8
const LedWidth = 16

type ILedCanvas interface {
	PreShow()
	Show(c LedCube)
}

/**
DummyLedCanvas for test
*/
type DummyLedCanvas struct {
}

func (canvas *DummyLedCanvas) PreShow() {

}

func (canvas *DummyLedCanvas) Show(c LedCube) {
}

/**
LedCanvas for real ledcube
*/
type Color32 interface {
	Uint32() uint32
	IsOff() bool
}

type RGB struct {
	r   uint8
	g   uint8
	b   uint8
	rgb uint32
}

func (rgb *RGB) Uint32() uint32 {
	return rgb.rgb
}
func (rgb *RGB) IsOff() bool {
	return rgb.rgb == 0
}

func NewFromRGB(r, g, b uint8) Color32 {
	return &RGB{r, g, b, ToUint32(r, g, b)}
}

func NewFromColorColor(c color.Color) Color32 {
	var r, g, b uint8
	rr, gg, bb, _ := c.RGBA()
	r = uint8(rr / 0x100)
	g = uint8(gg / 0x100)
	b = uint8(bb / 0x100)
	return &RGB{r, g, b, ToUint32(r, g, b)}
}

func ToUint32(r, g, b uint8) uint32 {
	return (uint32(r) << 16) | (uint32(g) << 8) | uint32(b)
}

type LedCanvas struct {
}

func NewLedCanvas() ILedCanvas {
	return &LedCanvas{}
}

func (canvas *LedCanvas) PreShow() {

}

func (canvas *LedCanvas) Show(c LedCube) {
	EnumXYZ(LedWidth, LedHeight, LedDepth, func(x, y, z int) {
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
