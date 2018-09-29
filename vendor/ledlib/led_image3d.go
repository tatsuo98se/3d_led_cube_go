package ledlib

import "ledlib/util"

// 内部のバッファサイズ
// マイナス座標の書き込みや、範囲外の書き込みも多少受け付けるため、大きめのバッファサイズにする
const LedInternalWidh = LedWidth * 3
const LedInternalHeight = LedHeight * 3
const LedInternalDepth = LedDepth * 3

const ledCubeOffsetX = LedWidth
const ledCubeOffsetY = LedHeight
const ledCubeOffsetZ = LedDepth

func NewLedData3D() util.Data3D {
	return util.NewData3D(
		LedInternalWidh, LedInternalHeight, LedInternalDepth,
		ledCubeOffsetX, ledCubeOffsetY, ledCubeOffsetZ)
}

func NewLedImage3D() util.Image3D {
	return util.NewImage3D(
		LedInternalWidh, LedInternalHeight, LedInternalDepth,
		ledCubeOffsetX, ledCubeOffsetY, ledCubeOffsetZ)
}
