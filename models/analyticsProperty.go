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
	// IANA timezone used to define the daily boundary for stats.
	Timezone string `json:"timezone"`
	// Whether tracking is currently active.
	Enabled bool `json:"enabled"`
	// Whether stats for this property are publicly viewable.
	Public bool `json:"public"`
	// List of origins allowed to send tracking events. Use ["*"] to allow all.
	AllowedOrigins []string `json:"allowedOrigins"`
	// Unique identifier for the tracking script snippet.
	SnippetId string `json:"snippetId"`

	// Used by Decode() method
	data []byte
}

func (model AnalyticsProperty) New(data []byte) *AnalyticsProperty {
	model.data = data
	return &model
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
