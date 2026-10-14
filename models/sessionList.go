package models

import (
	"encoding/json"
	"errors"
)

// SessionsList Model
type SessionList struct {
	// Total number of sessions that matched your query.
	Total int `json:"total"`
	// List of sessions.
	Sessions []Session `json:"sessions"`

	// Used by Decode() method
	data []byte
}

func (model SessionList) New(data []byte) *SessionList {
	model.data = data
	return &model
}

func (model *SessionList) UnmarshalJSON(data []byte) error {
	type alias SessionList
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

func (model *SessionList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
