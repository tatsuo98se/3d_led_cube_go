package webapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalCOnfigration1(t *testing.T) {
	config, err := UnmarshalConfigration([]byte(`{"enable":true}`))
	assert.Nil(t, err)
	assert.True(t, config.Enable)
}

func TestUnmarshalCOnfigration2(t *testing.T) {
	config, err := UnmarshalConfigration([]byte(`{"enable":false}`))
	assert.Nil(t, err)
	assert.False(t, config.Enable)
}

func TestUnmarshalCOnfigrationInvalidBody(t *testing.T) {
	config, err := UnmarshalConfigration([]byte(`{"enabl":"abd"}`))
	assert.Nil(t, err)
	assert.True(t, config.Enable)
}
