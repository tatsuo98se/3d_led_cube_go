package util

import (
	"sync"
)

const usingCore = 2

type EnumImage3DCallback func(x, y, z int, c Color32)
type Image3D interface {
	SetAt(x, y, z int, c Color32)
	GetAt(x, y, z int) Color32
	ForEach(callback EnumImage3DCallback)
	ConcurrentForEach(callback EnumImage3DCallback)
	Clear()
	Fill(c Color32)
}

type Image3DImpl struct {
	X, Y, Z                   int
	offsetX, offsetY, offsetZ int
	image                     [][][]Color32
}

func NewImage3D(x, y, z, offsetX, offsetY, offsetZ int) Image3D {
	cube := Image3DImpl{
		x, y, z,
		offsetX, offsetY, offsetZ,
		make([][][]Color32, x)}

	for xx := range cube.image {
		cube.image[xx] = make([][]Color32, cube.Y)
		for yy := range cube.image[xx] {
			cube.image[xx][yy] = make([]Color32, cube.Z)
		}
	}
	return &cube
}

func (l *Image3DImpl) isInRange(x, y, z int) bool {
	switch {
	case 0 > x+l.offsetX:
		fallthrough
	case x+l.offsetX >= l.X:
		fallthrough
	case 0 > y+l.offsetY:
		fallthrough
	case y+l.offsetY >= l.Y:
		fallthrough
	case 0 > z+l.offsetZ:
		fallthrough
	case z+l.offsetZ >= l.Z:
		return false
	}
	return true
}

func (l *Image3DImpl) SetAt(x, y, z int, c Color32) {
	if l.isInRange(x, y, z) {
		l.image[x+l.offsetX][y+l.offsetY][z+l.offsetZ] = c
	}
}

func (l *Image3DImpl) GetAt(x, y, z int) Color32 {
	if l.isInRange(x, y, z) {
		return l.image[x+l.offsetX][y+l.offsetY][z+l.offsetZ]
	} else {
		return NewFromRGB(0, 0, 0)
	}
}

func (l *Image3DImpl) Clear() {
	ConcurrentEnumXYZ(l.X, l.Y, l.Z, func(x, y, z int) {
		l.SetAt(x, y, z, nil)
	})
}
func (l *Image3DImpl) Fill(c Color32) {
	ConcurrentEnumXYZ(l.X, l.Y, l.Z, func(x, y, z int) {
		l.SetAt(x, y, z, c)
	})
}

func (l *Image3DImpl) ForEach(callback EnumImage3DCallback) {
	EnumXYZ(l.X, l.Y, l.Z, func(x, y, z int) {
		c := l.GetAt(x, y, z)
		if c != nil && !c.IsOff() {
			callback(x, y, z, c)
		}
	})
}
func (l *Image3DImpl) ConcurrentForEach(callback EnumImage3DCallback) {
	ConcurrentEnumXYZ(l.X, l.Y, l.Z, func(x, y, z int) {
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
	wg.Add(usingCore)
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

	work := x / usingCore
	for c := 0; c < usingCore; c++ {
		if c == usingCore-1 {
			go xloop(c*work, x)
		} else {
			go xloop(c*work, (c+1)*work)
		}
	}
	wg.Wait()
}
