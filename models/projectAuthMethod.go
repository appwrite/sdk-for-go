package models

import (
	"encoding/json"
	"errors"
)

// ProjectAuthMethod Model
type ProjectAuthMethod struct {
	// Auth method ID.
	Id string `json:"$id"`
	// Auth method status.
	Enabled bool `json:"enabled"`

	// Used by Decode() method
	data []byte
}

func (model ProjectAuthMethod) New(data []byte) *ProjectAuthMethod {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *ProjectAuthMethod) UnmarshalJSON(data []byte) error {
	type alias ProjectAuthMethod
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Enabled json.RawMessage `json:"enabled"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Enabled); err != nil {
			return err
		}
	}

	*model = ProjectAuthMethod(decoded)
	return nil
}

func (model *ProjectAuthMethod) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
