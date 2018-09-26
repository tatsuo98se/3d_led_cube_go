package util

import (
	"image/png"
	"log"
	"os"
)

type Image2D interface {
	GetAt(x, y int) Color32
	GetWidth() int
	GetHeight() int
}

type Image2DImpl struct {
	image [][]Color32
	X, Y  int
}

func NewImage2D(path string) Image2D {
	image := Image2DImpl{}
	image.load(path)
	return &image
}

func (i *Image2DImpl) GetAt(x, y int) Color32 {
	if i.isInRange(x, y) {
		return i.image[x][y]
	} else {
		return NewFromRGB(0, 0, 0)
	}
}

func (i *Image2DImpl) GetWidth() int {
	return i.X
}

func (i *Image2DImpl) GetHeight() int {
	return i.Y
}

func (i *Image2DImpl) isInRange(x, y int) bool {
	switch {
	case 0 > x:
		fallthrough
	case x >= i.X:
		fallthrough
	case 0 > y:
		fallthrough
	case y >= i.Y:
		return false
	}
	return true
}
func (i *Image2DImpl) SetAt(x, y int, c Color32) {
	if i.isInRange(x, y) {
		i.image[x][y] = c
	}
}
func (i *Image2DImpl) load(path string) {

	reader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	m, err := png.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	i.X, i.Y = m.Bounds().Dx(), m.Bounds().Dy()
	i.image = MakeImage(i.X, i.Y)
	for x := 0; x < i.X; x++ {
		for y := 0; y < i.Y; y++ {
			if m == nil {
				continue
			}
			i.SetAt(x, y, NewFromColorColor(m.At(x, y)))
		}
	}
}

func MakeImage(x, y int) [][]Color32 {
	image := make([][]Color32, x)

	for xx := range image {
		image[xx] = make([]Color32, y)
	}
	return image
}
