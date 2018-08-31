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
	SetAt(x, y, z int, c color.Color)
	Show()
}

/**
DummyLedCanvas for test
*/
type DummyLedCanvas struct {
}

func (canvas *DummyLedCanvas) SetAt(x, y, z int, c color.Color) {
}

func (canvas *DummyLedCanvas) Show() {
}

/**
LedCanvas for real ledcube
*/
type LedCanvas struct {
	canvas map[string]color.Color
}

func NewLedCanvas() ILedCanvas {
	return &LedCanvas{make(map[string]color.Color)}
}

func (canvas *LedCanvas) SetAt(x, y, z int, c color.Color) {
	canvas.canvas[fmt.Sprintf("%d:%d:%d", x, y, z)] = c

}

func (canvas *LedCanvas) Show() {
	for k, v := range canvas.canvas {
		s := strings.Split(k, ":")
		pt, err := Atois(s)
		if err == nil {
			C.SetLed(C.int(pt[0]), C.int(pt[1]), C.int(pt[2]), C.int(ColorToUint32(v)))
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

func ColorToUint32(c color.Color) uint32 {

	var r, g, b uint32
	rr, gg, bb, _ := c.RGBA()
	r = rr / 0x100
	g = gg / 0x100
	b = bb / 0x100
	return (r << 16) | (g << 8) | b
}
