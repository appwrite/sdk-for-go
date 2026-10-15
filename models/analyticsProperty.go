package models

import (
	"encoding/json"
	"errors"
)

// AnalyticsProperty Model
type AnalyticsProperty struct {
	// Analytics property ID.
	Id string `json:"$id"`
	// Property creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Property update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Human-readable name of the tracked website or application.
	Name string `json:"name"`
	// Primary domain being tracked (e.g. example.com). May be empty for native
	// apps.
	Domain string `json:"domain"`
	// Whether tracking is currently active.
	Enabled bool `json:"enabled"`
	// List of origins allowed to send tracking events. Use ["*"] to allow all.
	AllowedOrigins []string `json:"allowedOrigins"`
	// Most recent event date in ISO 8601 format. This attribute is only updated
	// again after 24 hours.
	AccessedAt string `json:"accessedAt"`
	// First event date in ISO 8601 format. Empty until the property receives its
	// first event.
	FirstAccessedAt string `json:"firstAccessedAt"`

	// Used by Decode() method
	data []byte
}

func (model AnalyticsProperty) New(data []byte) *AnalyticsProperty {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AnalyticsProperty) UnmarshalJSON(data []byte) error {
	type alias AnalyticsProperty
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id              json.RawMessage `json:"$id"`
		CreatedAt       json.RawMessage `json:"$createdAt"`
		UpdatedAt       json.RawMessage `json:"$updatedAt"`
		Name            json.RawMessage `json:"name"`
		Domain          json.RawMessage `json:"domain"`
		Enabled         json.RawMessage `json:"enabled"`
		AllowedOrigins  json.RawMessage `json:"allowedOrigins"`
		AccessedAt      json.RawMessage `json:"accessedAt"`
		FirstAccessedAt json.RawMessage `json:"firstAccessedAt"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.CreatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedAt); err != nil {
			return err
		}
	}
	if present.UpdatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UpdatedAt); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Domain == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Domain); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.AllowedOrigins == nil {
		if err := json.Unmarshal([]byte("[\"*\"]"), &decoded.AllowedOrigins); err != nil {
			return err
		}
	}
	if present.AccessedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AccessedAt); err != nil {
			return err
		}
	}
	if present.FirstAccessedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.FirstAccessedAt); err != nil {
			return err
		}
	}

	*model = AnalyticsProperty(decoded)
	return nil
}

func (model *AnalyticsProperty) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
