package ledlib

import "ledlib/util"

func ShowObject(canvas LedCanvas, obj LedObject, param LedCanvasParam) {
	canvas.Show(obj.GetImage3D(), param)
}

type LedObject interface {
	GetImage3D() util.Image3D
}
