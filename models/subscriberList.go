package models

import (
	"encoding/json"
	"errors"
)

// SubscriberList Model
type SubscriberList struct {
	// Total number of subscribers that matched your query.
	Total int `json:"total"`
	// List of subscribers.
	Subscribers []Subscriber `json:"subscribers"`

	// Used by Decode() method
	data []byte
}

func (model SubscriberList) New(data []byte) *SubscriberList {
	model.data = data
	return &model
}

func (model *SubscriberList) UnmarshalJSON(data []byte) error {
	type alias SubscriberList
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

func (model *SubscriberList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
