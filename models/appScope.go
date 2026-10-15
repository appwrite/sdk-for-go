package models

import (
	"encoding/json"
	"errors"
)

// AppScope Model
type AppScope struct {
	// Scope value as requested by apps.
	Value string `json:"value"`
	// Human-readable description of what the scope grants.
	Description string `json:"description"`
	// What the scope grants access to. One of `account`, `project`, or
	// `organization`. Only `project` and `organization` scopes are installable.
	Type string `json:"type"`
	// Scope category, used to group scopes on consent and installation screens.
	Category string `json:"category"`
	// Whether the scope is deprecated. Deprecated scopes can still be requested
	// but should not be offered for new grants.
	Deprecated bool `json:"deprecated"`

	// Used by Decode() method
	data []byte
}

func (model AppScope) New(data []byte) *AppScope {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AppScope) UnmarshalJSON(data []byte) error {
	type alias AppScope
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Value       json.RawMessage `json:"value"`
		Description json.RawMessage `json:"description"`
		Type        json.RawMessage `json:"type"`
		Category    json.RawMessage `json:"category"`
		Deprecated  json.RawMessage `json:"deprecated"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Value == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Value); err != nil {
			return err
		}
	}
	if present.Description == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Description); err != nil {
			return err
		}
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Type); err != nil {
			return err
		}
	}
	if present.Category == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Category); err != nil {
			return err
		}
	}
	if present.Deprecated == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Deprecated); err != nil {
			return err
		}
	}

	*model = AppScope(decoded)
	return nil
}

func (model *AppScope) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
