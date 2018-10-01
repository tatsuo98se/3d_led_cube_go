package ledlib

import (
	"ledlib/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DummyObject struct {
	Expired bool
}

func (o *DummyObject) IsExpired() bool {
	return o.Expired
}
func (o *DummyObject) Draw(cube util.Image3D) {

}

func TestNewFilterObjects(t *testing.T) {
	target := NewFilterObjects(&DummyLedCanvas{})
	obj1 := DummyObject{false}
	rocket := NewObjectRocket()

	target.Append(&obj1)
	ShowObject(target, rocket, nil)
	assert.Equal(t, 1, target.Len())

	obj1.Expired = true
	ShowObject(target, rocket, nil)
	assert.Equal(t, 0, target.Len())
}
