package util

type EnumImage3DCallback func(x, y, z int, c Color32)
type Image3D interface {
	SetAt(x, y, z int, c Color32)
	GetAt(x, y, z int) Color32
	ForEach(callback EnumImage3DCallback)
	ConcurrentForEach(callback EnumImage3DCallback)
	ConcurrentForEachAll(callback EnumImage3DCallback)
	Clear()
	Fill(c Color32)
}

type Image3DImpl struct {
	data Data3D
}

func NewImage3D(x, y, z, offsetX, offsetY, offsetZ int) Image3D {

	cube := Image3DImpl{}
	cube.data = NewData3D(x, y, z, offsetX, offsetY, offsetZ)
	return &cube
}

func (l *Image3DImpl) isInRange(x, y, z int) bool {
	return l.data.IsInRange(x, y, z)
}

func (l *Image3DImpl) SetAt(x, y, z int, c Color32) {
	l.data.SetAt(x, y, z, c)
}

func (l *Image3DImpl) GetAt(x, y, z int) Color32 {
	if l.isInRange(x, y, z) {
		if data := l.data.GetAt(x, y, z); data != nil {
			return data.(Color32)
		} else {
			return nil
		}
	}
	return NewFromRGB(0, 0, 0)
}

func (l *Image3DImpl) Clear() {
	l.data.Clear()
}
func (l *Image3DImpl) Fill(c Color32) {
	l.data.Fill(c)
}

func (l *Image3DImpl) ForEach(callback EnumImage3DCallback) {
	l.data.ForEach(func(x, y, z int, data interface{}) {
		c := data.(Color32)
		if c != nil && !c.IsOff() {
			callback(x, y, z, c)
		}
	})
}

func (l *Image3DImpl) ConcurrentForEach(callback EnumImage3DCallback) {
	l.data.ConcurrentForEach(func(x, y, z int, data interface{}) {
		c := data.(Color32)
		if c != nil && !c.IsOff() {
			callback(x, y, z, c)
		}
	})
}
func (l *Image3DImpl) ConcurrentForEachAll(callback EnumImage3DCallback) {
	l.data.ConcurrentForEachAll(func(x, y, z int, data interface{}) {
		c, _ := data.(Color32)
		callback(x, y, z, c)
	})
}
