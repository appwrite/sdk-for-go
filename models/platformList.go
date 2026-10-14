package models

import (
	"encoding/json"
	"errors"
)

// PlatformsList Model
type PlatformList struct {
	// Total number of platforms in the given project.
	Total int `json:"total"`
	// List of platforms.
	Platforms []interface{} `json:"platforms"`

	// Used by Decode() method
	data []byte
}

func (model PlatformList) New(data []byte) *PlatformList {
	model.data = data
	return &model
}

func (model *PlatformList) UnmarshalJSON(data []byte) error {
	type alias PlatformList
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

func (model *PlatformList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
