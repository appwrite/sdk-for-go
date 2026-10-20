package models

import (
	"encoding/json"
	"errors"
)

// PolicyPasswordPersonalData Model
type PolicyPasswordPersonalData struct {
	// Policy ID.
	Id string `json:"$id"`
	// Whether password personal data policy is enabled.
	Enabled bool `json:"enabled"`
	// Custom terms blocked by case-insensitive substring matching when enabled.
	DenyList []string `json:"denyList"`

	// Used by Decode() method
	data []byte
}

func (model PolicyPasswordPersonalData) New(data []byte) *PolicyPasswordPersonalData {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *PolicyPasswordPersonalData) UnmarshalJSON(data []byte) error {
	type alias PolicyPasswordPersonalData
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id       json.RawMessage `json:"$id"`
		Enabled  json.RawMessage `json:"enabled"`
		DenyList json.RawMessage `json:"denyList"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.DenyList == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.DenyList); err != nil {
			return err
		}
	}

	*model = PolicyPasswordPersonalData(decoded)
	return nil
}

func (model *PolicyPasswordPersonalData) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
