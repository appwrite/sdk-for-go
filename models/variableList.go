package models

import (
	"encoding/json"
	"errors"
)

// VariablesList Model
type VariableList struct {
	// Total number of variables that matched your query.
	Total int `json:"total"`
	// List of variables.
	Variables []Variable `json:"variables"`

	// Used by Decode() method
	data []byte
}

func (model VariableList) New(data []byte) *VariableList {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *VariableList) UnmarshalJSON(data []byte) error {
	type alias VariableList
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

	*model = VariableList(decoded)
	return nil
}

func (model *VariableList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
