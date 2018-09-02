package util

import "image/color"

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
