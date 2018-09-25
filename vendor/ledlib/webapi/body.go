package webapi

import (
	"encoding/json"
)

type Configration struct {
	Enable bool `json:"enable"`
}

func UnmarshalConfigration(body []byte) (*Configration, error) {
	config := Configration{true}
	if err := json.Unmarshal(body, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
