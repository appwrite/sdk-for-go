package models

import (
	"encoding/json"
	"errors"
)

// DocumentsList Model
type DocumentList struct {
	// Total number of documents that matched your query.
	Total int `json:"total"`
	// List of documents.
	Documents []Document `json:"documents"`

	// Used by Decode() method
	data []byte
}

func (model DocumentList) New(data []byte) *DocumentList {
	model.data = data
	return &model
}

func (model *DocumentList) UnmarshalJSON(data []byte) error {
	type alias DocumentList
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

func (model *DocumentList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
