package ledlib

import "ledlib/util"

// 内部のバッファサイズ
// マイナス座標の書き込みや、範囲外の書き込みも多少受け付けるため、大きめのバッファサイズにする
const ledCubeWidh = LedWidth * 3
const ledCubeHeight = LedHeight * 3
const ledCubeDepth = LedDepth * 3

const ledCubeOffsetX = LedWidth
const ledCubeOffsetY = LedHeight
const ledCubeOffsetZ = LedDepth

func NewLedCubeImage() util.CubeImage {
	return util.NewCubeImage(
		ledCubeWidh, ledCubeHeight, ledCubeDepth,
		ledCubeOffsetX, ledCubeOffsetY, ledCubeOffsetZ)
}
