package models

import (
	"encoding/json"
	"errors"
)

// EmailTemplatesList Model
type EmailTemplateList struct {
	// Total number of templates that matched your query.
	Total int `json:"total"`
	// List of templates.
	Templates []EmailTemplate `json:"templates"`

	// Used by Decode() method
	data []byte
}

func (model EmailTemplateList) New(data []byte) *EmailTemplateList {
	model.data = data
	return &model
}

func (model *EmailTemplateList) UnmarshalJSON(data []byte) error {
	type alias EmailTemplateList
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

func (model *EmailTemplateList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
