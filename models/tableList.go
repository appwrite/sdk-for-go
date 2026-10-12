package models

import (
	"encoding/json"
	"errors"
)

// TablesList Model
type TableList struct {
	// Total number of tables that matched your query.
	Total int `json:"total"`
	// List of tables.
	Tables []Table `json:"tables"`

	// Used by Decode() method
	data []byte
}

func (model TableList) New(data []byte) *TableList {
	model.data = data
	return &model
}

func (model *TableList) UnmarshalJSON(data []byte) error {
	type alias TableList
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

func (model *TableList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
