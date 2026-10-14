package models

import (
	"encoding/json"
	"errors"
)

// ProjectsList Model
type ProjectList struct {
	// Total number of projects that matched your query.
	Total int `json:"total"`
	// List of projects.
	Projects []Project `json:"projects"`

	// Used by Decode() method
	data []byte
}

func (model ProjectList) New(data []byte) *ProjectList {
	model.data = data
	return &model
}

func (model *ProjectList) UnmarshalJSON(data []byte) error {
	type alias ProjectList
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

func (model *ProjectList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
