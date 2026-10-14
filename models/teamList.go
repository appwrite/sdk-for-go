package models

import (
	"encoding/json"
	"errors"
)

// TeamsList Model
type TeamList struct {
	// Total number of teams that matched your query.
	Total int `json:"total"`
	// List of teams.
	Teams []Team `json:"teams"`

	// Used by Decode() method
	data []byte
}

func (model TeamList) New(data []byte) *TeamList {
	model.data = data
	return &model
}

func (model *TeamList) UnmarshalJSON(data []byte) error {
	type alias TeamList
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

func (model *TeamList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
