package models

import (
	"encoding/json"
	"errors"
)

// PolicyUserLimit Model
type PolicyUserLimit struct {
	// Policy ID.
	Id string `json:"$id"`
	// Maximum number of users allowed in the project. A value of 0 means the
	// policy is disabled.
	Total int `json:"total"`

	// Used by Decode() method
	data []byte
}

func (model PolicyUserLimit) New(data []byte) *PolicyUserLimit {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *PolicyUserLimit) UnmarshalJSON(data []byte) error {
	type alias PolicyUserLimit
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

	*model = PolicyUserLimit(decoded)
	return nil
}

func (model *PolicyUserLimit) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
