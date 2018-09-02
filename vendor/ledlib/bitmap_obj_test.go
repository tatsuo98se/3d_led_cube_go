package ledlib

import (
	"testing"
)

func TestNewBitmapObj(t *testing.T) {

	target := NewRocketBitmapObj()
	if target == nil {
		t.Fail()
	}
}
