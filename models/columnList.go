package models

import (
	"encoding/json"
	"errors"
)

// ColumnsList Model
type ColumnList struct {
	// Total number of columns in the given table.
	Total int `json:"total"`
	// List of columns.
	Columns []interface{} `json:"columns"`

	// Used by Decode() method
	data []byte
}

func (model ColumnList) New(data []byte) *ColumnList {
	model.data = data
	return &model
}

func (model *ColumnList) UnmarshalJSON(data []byte) error {
	type alias ColumnList
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

func (model *ColumnList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
