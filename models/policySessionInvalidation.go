package models

import (
	"encoding/json"
	"errors"
)

// PolicySessionInvalidation Model
type PolicySessionInvalidation struct {
	// Policy ID.
	Id string `json:"$id"`
	// Whether session invalidation policy is enabled.
	Enabled bool `json:"enabled"`

	// Used by Decode() method
	data []byte
}

func (model PolicySessionInvalidation) New(data []byte) *PolicySessionInvalidation {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *PolicySessionInvalidation) UnmarshalJSON(data []byte) error {
	type alias PolicySessionInvalidation
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id      json.RawMessage `json:"$id"`
		Enabled json.RawMessage `json:"enabled"`
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
		if err := json.Unmarshal([]byte("true"), &decoded.Enabled); err != nil {
			return err
		}
	}

	*model = PolicySessionInvalidation(decoded)
	return nil
}

func (model *PolicySessionInvalidation) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
