package ledlib

import (
	"testing"
)

func TestisInRange(t *testing.T) {
	if !isInRange(0, 0, 0) {
		t.Error()
	}
	if !isInRange(15, 31, 7) {
		t.Error()
	}
}

func TestisNoInRange(t *testing.T) {
	if isInRange(-1, 0, 0) {
		t.Error()
	}
	if isInRange(0, -1, 0) {
		t.Error()
	}
	if isInRange(0, 0, -1) {
		t.Error()
	}
	if isInRange(-1, -1, -1) {
		t.Error()
	}

	if isInRange(16, 0, 0) {
		t.Error()
	}
	if isInRange(0, 32, 0) {
		t.Error()
	}
	if isInRange(0, 0, 8) {
		t.Error()
	}
	if isInRange(16, 32, 8) {
		t.Error()
	}
}
