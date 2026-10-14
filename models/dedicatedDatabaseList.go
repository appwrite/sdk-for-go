package models

import (
	"encoding/json"
	"errors"
)

// DedicatedDatabasesList Model
type DedicatedDatabaseList struct {
	// Total number of databases that matched your query.
	Total int `json:"total"`
	// List of databases.
	Databases []DedicatedDatabase `json:"databases"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabaseList) New(data []byte) *DedicatedDatabaseList {
	model.data = data
	return &model
}

func (model *DedicatedDatabaseList) UnmarshalJSON(data []byte) error {
	type alias DedicatedDatabaseList
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

func (model *DedicatedDatabaseList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
