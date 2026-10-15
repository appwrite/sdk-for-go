package models

import (
	"encoding/json"
	"errors"
)

// PolicyPasswordRotation Model
type PolicyPasswordRotation struct {
	// Policy ID.
	Id string `json:"$id"`
	// Whether expired passwords require a reset before email/password sign-in.
	Enabled bool `json:"enabled"`
	// Maximum password age in seconds.
	MaxAge int `json:"maxAge"`

	// Used by Decode() method
	data []byte
}

func (model PolicyPasswordRotation) New(data []byte) *PolicyPasswordRotation {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *PolicyPasswordRotation) UnmarshalJSON(data []byte) error {
	type alias PolicyPasswordRotation
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id      json.RawMessage `json:"$id"`
		Enabled json.RawMessage `json:"enabled"`
		MaxAge  json.RawMessage `json:"maxAge"`
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
	if present.MaxAge == nil {
		if err := json.Unmarshal([]byte("31536000"), &decoded.MaxAge); err != nil {
			return err
		}
	}

	*model = PolicyPasswordRotation(decoded)
	return nil
}

func (model *PolicyPasswordRotation) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
