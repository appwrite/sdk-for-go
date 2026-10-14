package models

import (
	"encoding/json"
	"errors"
)

// Volume Model
type DatabaseStatusVolume struct {
	// Mount path of the volume.
	Path string `json:"path"`
	// Percentage of storage used.
	UsedPercent string `json:"usedPercent"`
	// Available storage space.
	Available string `json:"available"`
	// Whether the volume is mounted.
	Mounted bool `json:"mounted"`

	// Used by Decode() method
	data []byte
}

func (model DatabaseStatusVolume) New(data []byte) *DatabaseStatusVolume {
	model.data = data
	return &model
}

func (model *DatabaseStatusVolume) UnmarshalJSON(data []byte) error {
	type alias DatabaseStatusVolume
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"path\":\"\",\"usedPercent\":\"0%\",\"available\":\"0GB\",\"mounted\":false}"), &fields); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	merged, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	return json.Unmarshal(merged, (*alias)(model))
}

func (model *DatabaseStatusVolume) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
