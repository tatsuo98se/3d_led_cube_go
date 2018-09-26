package ledlib

import (
	"testing"
)

func TestNewObjectCubeBitmap(t *testing.T) {

	target := NewObjectRocket()
	if target == nil {
		t.Fail()
	}
}
