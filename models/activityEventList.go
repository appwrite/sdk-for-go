package models

import (
	"encoding/json"
	"errors"
)

// ActivityEventList Model
type ActivityEventList struct {
	// Total number of events that matched your query.
	Total int `json:"total"`
	// List of events.
	Events []ActivityEvent `json:"events"`

	// Used by Decode() method
	data []byte
}

func (model ActivityEventList) New(data []byte) *ActivityEventList {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *ActivityEventList) UnmarshalJSON(data []byte) error {
	type alias ActivityEventList
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

	*model = ActivityEventList(decoded)
	return nil
}

func (model *ActivityEventList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
