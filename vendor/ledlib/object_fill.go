package ledlib

import (
	"ledlib/util"
)

type ObjectFill struct {
	cube util.CubeImage
}

func NewObjectFill(c util.Color32) LedObject {
	obj := ObjectFill{NewLedCubeImage()}
	obj.cube.Fill(c)
	return &obj
}

func (b *ObjectFill) Draw(canvas LedCanvas) {
	canvas.Show(b.cube)
}

func (b *ObjectFill) DidDetach() {

}
