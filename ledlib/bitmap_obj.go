package ledlib

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
)

type BitmapObj struct {
	canvas ILedCanvas
	imgx   [][][]color.Color
}

func NewBitmapObj(canvas ILedCanvas, paths []string) *BitmapObj {
	bmp := BitmapObj{canvas,
		make([][][]color.Color, LedWidth)}

	for x := range bmp.imgx {
		bmp.imgx[x] = make([][]color.Color, LedHeight)
		for y := range bmp.imgx[x] {
			bmp.imgx[x][y] = make([]color.Color, LedDepth)
		}
	}

	bmp.load(paths)

	return &bmp
}

func (obj *BitmapObj) load(paths []string) {

	for z, path := range paths {
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
		for x := 0; x < LedWidth; x++ {
			for y := 0; y < LedHeight; y++ {
				if m == nil {
					continue
				}
				fmt.Println(x, y, z)
				obj.imgx[x][y][z] = m.At(x, y)
			}
		}
	}
}

func (obj *BitmapObj) Draw() {
	for x := 0; x < LedWidth; x++ {
		for y := 0; y < LedHeight; y++ {
			for z := 0; z < LedDepth; z++ {
				if obj.imgx[x][y][z] == nil {
					continue
				}
				obj.canvas.SetAt(x, y, z, obj.imgx[x][y][z])
			}
		}
	}
}
