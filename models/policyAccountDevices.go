package models

import (
	"encoding/json"
	"errors"
)

// PolicyAccountDevices Model
type PolicyAccountDevices struct {
	// Policy ID.
	Id string `json:"$id"`
	// Whether account devices are enabled. When disabled, new account devices can
	// no longer be created.
	Enabled bool `json:"enabled"`

	// Used by Decode() method
	data []byte
}

func (model PolicyAccountDevices) New(data []byte) *PolicyAccountDevices {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *PolicyAccountDevices) UnmarshalJSON(data []byte) error {
	type alias PolicyAccountDevices
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

	*model = PolicyAccountDevices(decoded)
	return nil
}

func (model *PolicyAccountDevices) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
