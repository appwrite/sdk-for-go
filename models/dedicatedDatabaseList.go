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
