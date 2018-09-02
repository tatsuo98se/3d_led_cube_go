package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundToInt(t *testing.T) {

	assert.Equal(t, 1, RoundToInt(0.5))
	assert.Equal(t, 0, RoundToInt(0.4))
	assert.Equal(t, 1, RoundToInt(1.1))
}
