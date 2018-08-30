package ledlib

import (
	"testing"
)

func TestNewBitmapObj(t *testing.T) {

	canvas := &DummyLedCanvas{}
	target := NewRocketBitmapObj(canvas)
	if target == nil {
		t.Fail()
	}
}
