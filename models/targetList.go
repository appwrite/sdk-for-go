package models

import (
	"encoding/json"
	"errors"
)

// TargetList Model
type TargetList struct {
	// Total number of targets that matched your query.
	Total int `json:"total"`
	// List of targets.
	Targets []Target `json:"targets"`

	// Used by Decode() method
	data []byte
}

func (model TargetList) New(data []byte) *TargetList {
	model.data = data
	return &model
}

func (model *TargetList) UnmarshalJSON(data []byte) error {
	type alias TargetList
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

func (model *TargetList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
