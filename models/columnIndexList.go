package models

import (
	"encoding/json"
	"errors"
)

// ColumnIndexesList Model
type ColumnIndexList struct {
	// Total number of indexes that matched your query.
	Total int `json:"total"`
	// List of indexes.
	Indexes []ColumnIndex `json:"indexes"`

	// Used by Decode() method
	data []byte
}

func (model ColumnIndexList) New(data []byte) *ColumnIndexList {
	model.data = data
	return &model
}

func (model *ColumnIndexList) UnmarshalJSON(data []byte) error {
	type alias ColumnIndexList
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

func (model *ColumnIndexList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
