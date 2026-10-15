package models

import (
	"encoding/json"
	"errors"
)

// PolicySessionDuration Model
type PolicySessionDuration struct {
	// Policy ID.
	Id string `json:"$id"`
	// Session duration in seconds.
	Duration int `json:"duration"`

	// Used by Decode() method
	data []byte
}

func (model PolicySessionDuration) New(data []byte) *PolicySessionDuration {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *PolicySessionDuration) UnmarshalJSON(data []byte) error {
	type alias PolicySessionDuration
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id       json.RawMessage `json:"$id"`
		Duration json.RawMessage `json:"duration"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.Duration == nil {
		if err := json.Unmarshal([]byte("31536000"), &decoded.Duration); err != nil {
			return err
		}
	}

	*model = PolicySessionDuration(decoded)
	return nil
}

func (model *PolicySessionDuration) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
