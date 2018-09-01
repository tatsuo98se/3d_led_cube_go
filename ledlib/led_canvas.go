package ledlib

/*
#cgo LDFLAGS: -lledlib
#include "./../lib/led.h"
*/
import "C"
import (
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

const LedHeight = 32
const LedDepth = 8
const LedWidth = 16

type ILedCanvas interface {
	PreDraw()
	SetAt(x, y, z int, c Color32)
	Show()
}

/**
DummyLedCanvas for test
*/
type DummyLedCanvas struct {
}

func (canvas *DummyLedCanvas) PreDraw() {

}

func (canvas *DummyLedCanvas) SetAt(x, y, z int, c Color32) {
}

func (canvas *DummyLedCanvas) Show() {
}

/**
LedCanvas for real ledcube
*/
type Color32 interface {
	Uint32() uint32
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
	canvas map[string]Color32
}

func NewLedCanvas() ILedCanvas {
	return &LedCanvas{make(map[string]Color32)}
}

func (canvas *LedCanvas) PreDraw() {

}

func (canvas *LedCanvas) SetAt(x, y, z int, c Color32) {
	canvas.canvas[fmt.Sprintf("%d:%d:%d", x, y, z)] = c

}

func (canvas *LedCanvas) Show() {
	for k, v := range canvas.canvas {
		s := strings.Split(k, ":")
		pt, err := Atois(s)
		if err == nil {
			C.SetLed(C.int(pt[0]), C.int(pt[1]), C.int(pt[2]), C.int(v.Uint32()))
		}
	}
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
