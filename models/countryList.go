package models

import (
	"encoding/json"
	"errors"
)

// CountriesList Model
type CountryList struct {
	// Total number of countries that matched your query.
	Total int `json:"total"`
	// List of countries.
	Countries []Country `json:"countries"`

	// Used by Decode() method
	data []byte
}

func (model CountryList) New(data []byte) *CountryList {
	model.data = data
	return &model
}

func (model *CountryList) UnmarshalJSON(data []byte) error {
	type alias CountryList
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

func (model *CountryList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
