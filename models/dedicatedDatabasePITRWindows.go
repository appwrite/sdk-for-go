package models

import (
	"encoding/json"
	"errors"
)

// PITRWindows Model
type DedicatedDatabasePITRWindows struct {
	// Earliest available recovery point.
	Earliest string `json:"earliest"`
	// Latest available recovery point.
	Latest string `json:"latest"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabasePITRWindows) New(data []byte) *DedicatedDatabasePITRWindows {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DedicatedDatabasePITRWindows) UnmarshalJSON(data []byte) error {
	type alias DedicatedDatabasePITRWindows
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Earliest json.RawMessage `json:"earliest"`
		Latest   json.RawMessage `json:"latest"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Earliest == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Earliest); err != nil {
			return err
		}
	}
	if present.Latest == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Latest); err != nil {
			return err
		}
	}

	*model = DedicatedDatabasePITRWindows(decoded)
	return nil
}

func (model *DedicatedDatabasePITRWindows) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
