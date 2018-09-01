package ledlib

import "sync"

const ledCubeWidh = LedWidth * 3
const ledCubeHeight = LedHeight * 3
const ledCubeDepth = LedDepth * 3

const ledOffsetX = LedWidth
const ledOffsetY = LedHeight
const ledOffsetZ = LedDepth

type EnumLedCubeCallback func(x, y, z int, c Color32)
type LedCube interface {
	SetAt(x, y, z int, c Color32)
	GetAt(x, y, z int) Color32
	ForEach(callback EnumLedCubeCallback)
	ConcurrentForEach(callback EnumLedCubeCallback)
	Clear()
}

type LedCubeImpl struct {
	image [][][]Color32
}

func NewLedCube() LedCube {
	cube := LedCubeImpl{make([][][]Color32, ledCubeWidh)}
	for x := range cube.image {
		cube.image[x] = make([][]Color32, ledCubeHeight)
		for y := range cube.image[x] {
			cube.image[x][y] = make([]Color32, ledCubeDepth)
		}
	}
	return &cube
}

func isInRange(x, y, z int) bool {
	switch {
	case 0 > x+ledOffsetX:
		fallthrough
	case x+ledOffsetX >= ledCubeWidh:
		fallthrough
	case 0 > y+ledOffsetY:
		fallthrough
	case y+ledOffsetY >= ledCubeHeight:
		fallthrough
	case 0 > z+ledOffsetZ:
		fallthrough
	case z+ledOffsetZ >= ledCubeDepth:
		return false
	}
	return true
}

func (l *LedCubeImpl) SetAt(x, y, z int, c Color32) {
	if isInRange(x, y, z) {
		l.image[x+ledOffsetX][y+ledOffsetY][z+ledOffsetZ] = c
	}
}

func (l *LedCubeImpl) GetAt(x, y, z int) Color32 {
	if isInRange(x, y, z) {
		return l.image[x+ledOffsetX][y+ledOffsetY][z+ledOffsetZ]
	} else {
		return NewFromRGB(0, 0, 0)
	}
}

func (l *LedCubeImpl) Clear() {
	ConcurrentEnumXYZ(ledCubeWidh, ledCubeHeight, ledCubeDepth, func(x, y, z int) {
		l.SetAt(x, y, z, nil)
	})
}

func (l *LedCubeImpl) ForEach(callback EnumLedCubeCallback) {
	EnumXYZ(ledCubeWidh, ledCubeHeight, ledCubeDepth, func(x, y, z int) {
		c := l.GetAt(x, y, z)
		if c != nil && !c.IsOff() {
			callback(x, y, z, c)
		}
	})
}
func (l *LedCubeImpl) ConcurrentForEach(callback EnumLedCubeCallback) {
	ConcurrentEnumXYZ(ledCubeWidh, ledCubeHeight, ledCubeDepth, func(x, y, z int) {
		c := l.GetAt(x, y, z)
		if c != nil && !c.IsOff() {
			callback(x, y, z, c)
		}
	})
}

type EnumXYZCallback func(x, y, z int)

func EnumXYZ(x, y, z int, callback EnumXYZCallback) {
	for xx := 0; xx < x; xx++ {
		for yy := 0; yy < y; yy++ {
			for zz := 0; zz < z; zz++ {
				callback(xx, yy, zz)
			}
		}
	}
}
func ConcurrentEnumXYZ(x, y, z int, callback EnumXYZCallback) {
	var wg sync.WaitGroup
	wg.Add(2)
	xloop := func(xstart, xend int) {
		defer wg.Done()
		for xx := xstart; xx < xend; xx++ {
			for yy := 0; yy < y; yy++ {
				for zz := 0; zz < z; zz++ {
					callback(xx, yy, zz)
				}
			}
		}
	}

	xhelf := x / 2
	go xloop(0, xhelf)
	go xloop(xhelf, x)
	wg.Wait()
}
