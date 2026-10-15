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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *ColumnIndexList) UnmarshalJSON(data []byte) error {
	type alias ColumnIndexList
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Total json.RawMessage `json:"total"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Total == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Total); err != nil {
			return err
		}
	}

	*model = ColumnIndexList(decoded)
	return nil
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
