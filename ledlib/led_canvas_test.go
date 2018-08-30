package ledlib

import (
	"image/color"
	"testing"
)

func TestColorToUint32(t *testing.T) {

	data := &color.RGBA{0xff, 0xff, 0xff, 0xff}
	result := ColorToUint32(data)
	if result != 0xffffffff {
		t.Log(data)
		t.Fatalf("failed test result:%d", result)
	}
}
