package models

import (
	"encoding/json"
	"errors"
)

// AnalyticsPropertiesList Model
type AnalyticsPropertyList struct {
	// Total number of properties that matched your query.
	Total int `json:"total"`
	// List of properties.
	Properties []AnalyticsProperty `json:"properties"`

	// Used by Decode() method
	data []byte
}

func (model AnalyticsPropertyList) New(data []byte) *AnalyticsPropertyList {
	model.data = data
	return &model
}

func (model *AnalyticsPropertyList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
