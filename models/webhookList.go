package models

import (
	"encoding/json"
	"errors"
)

// WebhooksList Model
type WebhookList struct {
	// Total number of webhooks that matched your query.
	Total int `json:"total"`
	// List of webhooks.
	Webhooks []Webhook `json:"webhooks"`

	// Used by Decode() method
	data []byte
}

func (model WebhookList) New(data []byte) *WebhookList {
	model.data = data
	return &model
}

func (model *WebhookList) UnmarshalJSON(data []byte) error {
	type alias WebhookList
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

func (model *WebhookList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
