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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DatabaseStatusVolume) UnmarshalJSON(data []byte) error {
	type alias DatabaseStatusVolume
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Path        json.RawMessage `json:"path"`
		UsedPercent json.RawMessage `json:"usedPercent"`
		Available   json.RawMessage `json:"available"`
		Mounted     json.RawMessage `json:"mounted"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Path == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Path); err != nil {
			return err
		}
	}
	if present.UsedPercent == nil {
		if err := json.Unmarshal([]byte("\"0%\""), &decoded.UsedPercent); err != nil {
			return err
		}
	}
	if present.Available == nil {
		if err := json.Unmarshal([]byte("\"0GB\""), &decoded.Available); err != nil {
			return err
		}
	}
	if present.Mounted == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Mounted); err != nil {
			return err
		}
	}

	*model = DatabaseStatusVolume(decoded)
	return nil
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
