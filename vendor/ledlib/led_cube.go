package ledlib

import "ledlib/util"

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
