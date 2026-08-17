package models

import (
	"encoding/json"
	"errors"
)

// DatabaseMigrationsList Model
type DatabaseMigrationList struct {
	// Total number of migrations that matched your query.
	Total int `json:"total"`
	// List of migrations.
	Migrations []DatabaseMigration `json:"migrations"`

	// Used by Decode() method
	data []byte
}

func (model DatabaseMigrationList) New(data []byte) *DatabaseMigrationList {
	model.data = data
	return &model
}

func (model *DatabaseMigrationList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
