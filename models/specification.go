package models

import (
	"encoding/json"
	"errors"
)

// Specification Model
type Specification struct {
	// Memory size in MB.
	Memory int `json:"memory"`
	// Number of CPUs.
	Cpus float64 `json:"cpus"`
	// Is size enabled.
	Enabled bool `json:"enabled"`
	// Size slug.
	Slug string `json:"slug"`

	// Used by Decode() method
	data []byte
}

func (model Specification) New(data []byte) *Specification {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Specification) UnmarshalJSON(data []byte) error {
	type alias Specification
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Memory  json.RawMessage `json:"memory"`
		Cpus    json.RawMessage `json:"cpus"`
		Enabled json.RawMessage `json:"enabled"`
		Slug    json.RawMessage `json:"slug"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Memory == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Memory); err != nil {
			return err
		}
	}
	if present.Cpus == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Cpus); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.Slug == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Slug); err != nil {
			return err
		}
	}

	*model = Specification(decoded)
	return nil
}

func (model *Specification) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
