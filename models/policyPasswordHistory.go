package models

import (
	"encoding/json"
	"errors"
)

// PolicyPasswordHistory Model
type PolicyPasswordHistory struct {
	// Policy ID.
	Id string `json:"$id"`
	// Password history length. A value of 0 means the policy is disabled.
	Total int `json:"total"`

	// Used by Decode() method
	data []byte
}

func (model PolicyPasswordHistory) New(data []byte) *PolicyPasswordHistory {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *PolicyPasswordHistory) UnmarshalJSON(data []byte) error {
	type alias PolicyPasswordHistory
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id    json.RawMessage `json:"$id"`
		Total json.RawMessage `json:"total"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.Total == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Total); err != nil {
			return err
		}
	}

	*model = PolicyPasswordHistory(decoded)
	return nil
}

func (model *PolicyPasswordHistory) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
