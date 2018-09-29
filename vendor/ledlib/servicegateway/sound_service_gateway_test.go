package servicegateway

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMarshalSoundOrder(t *testing.T) {
	InitSoundSeriveGateway("tcp://localhost:5751")
	player := GetSoundSeriveGateway()
	defer player.Terminate()

	target := &soundOrder{}
	target.Function = "stop"

	json, err := json.Marshal(target)
	assert.Nil(t, err)
	assert.NotNil(t, json)
	t.Log(string(json))
}

func TestPlaySound(t *testing.T) {
	InitSoundSeriveGateway("tcp://localhost:5751")
	player := GetSoundSeriveGateway()
	defer player.Terminate()

	player.Play("asset/audio/se_jump.wav", false, false)
	time.Sleep(3 * time.Second)
}
