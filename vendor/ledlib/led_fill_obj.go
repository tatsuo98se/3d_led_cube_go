package ledlib

import (
	"ledlib/util"
)

type LedFillObj struct {
	cube util.CubeImage
}

func NewFillObj(c util.Color32) LedObject {
	obj := LedFillObj{NewLedCubeImage()}
	obj.cube.Fill(c)
	return &obj
}

func (b *LedFillObj) Draw(canvas LedCanvas) {
	canvas.Show(b.cube)
}

func (b *LedFillObj) DidDetach() {

}
