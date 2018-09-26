package ledlib

import (
	"ledlib/util"
)

type ObjectFill struct {
	cube util.Image3D
}

func NewObjectFill(c util.Color32) LedObject {
	obj := ObjectFill{NewLedImage3D()}
	obj.cube.Fill(c)
	return &obj
}

func (b *ObjectFill) GetImage3D() util.Image3D {
	return b.cube
}
