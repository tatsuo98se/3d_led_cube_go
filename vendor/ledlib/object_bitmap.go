package ledlib

import (
	"image/png"
	"ledlib/util"
	"log"
	"os"
)

type ObjectBitmap struct {
	cube util.CubeImage
}

func NewObjectBitmap(paths []string) LedObject {
	bmp := ObjectBitmap{NewLedCubeImage()}
	bmp.load(paths)
	return &bmp
}

func (b *ObjectBitmap) load(paths []string) {

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
		width, height := m.Bounds().Dx(), m.Bounds().Dy()
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				if m == nil {
					continue
				}
				b.cube.SetAt(x, y, z, util.NewFromColorColor(m.At(x, y)))
			}
		}
	}
}

func (b *ObjectBitmap) Draw(canvas LedCanvas) {
	canvas.Show(b.cube)
}

func (b *ObjectBitmap) DidDetach() {

}
