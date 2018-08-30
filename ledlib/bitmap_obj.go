package ledlib

import (
	"image"
	"image/png"
	"log"
	"os"
)

type BitmapObj struct {
	canvas ILedCanvas
	img    []image.Image
}

func NewBitmapObj(canvas ILedCanvas, paths []string) *BitmapObj {
	bmp := BitmapObj{canvas, make([]image.Image, LedDepth, LedDepth)}
	bmp.load(paths)
	return &bmp
}

func (obj *BitmapObj) load(paths []string) {

	for i, path := range paths {
		if path == "" {
			continue
		}
		reader, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer reader.Close()
		m, err := png.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}
		obj.img[i] = m
	}
}

func (obj *BitmapObj) Draw() {
	for x := 0; x < LedWidth; x++ {
		for y := 0; y < LedHeight; y++ {
			for z := 0; z < LedDepth; z++ {
				if obj.img[z] == nil {
					continue
				}
				obj.canvas.SetAt(x, y, z, obj.img[z].At(x, y))
			}
		}
	}
}
