package models

import (
	"encoding/json"
	"errors"
)

// ContinentsList Model
type ContinentList struct {
	// Total number of continents that matched your query.
	Total int `json:"total"`
	// List of continents.
	Continents []Continent `json:"continents"`

	// Used by Decode() method
	data []byte
}

func (model ContinentList) New(data []byte) *ContinentList {
	model.data = data
	return &model
}

func (model *ContinentList) UnmarshalJSON(data []byte) error {
	type alias ContinentList
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

func (model *ContinentList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
