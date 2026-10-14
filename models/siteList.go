package models

import (
	"encoding/json"
	"errors"
)

// SitesList Model
type SiteList struct {
	// Total number of sites that matched your query.
	Total int `json:"total"`
	// List of sites.
	Sites []Site `json:"sites"`

	// Used by Decode() method
	data []byte
}

func (model SiteList) New(data []byte) *SiteList {
	model.data = data
	return &model
}

func (model *SiteList) UnmarshalJSON(data []byte) error {
	type alias SiteList
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"total\":0}"), &fields); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	merged, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	return json.Unmarshal(merged, (*alias)(model))
}

func (model *SiteList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
