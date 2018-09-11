package ledlib

import (
	"image/color"
	"ledlib/util"
	"testing"
)

func TestColorToUint32(t *testing.T) {

	data := &color.RGBA{0xff, 0xff, 0xff, 0xff}
	result := util.NewFromColorColor(data).Uint32()
	if result != 0xffffff {
		t.Log(data)
		t.Fatalf("failed test result:%d", result)
	}
}
