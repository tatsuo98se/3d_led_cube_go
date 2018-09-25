package ledlib

import (
	"testing"
)

func TestNewObjectBitmap(t *testing.T) {

	target := NewObjectRocket()
	if target == nil {
		t.Fail()
	}
}
