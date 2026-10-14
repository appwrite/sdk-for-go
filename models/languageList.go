package models

import (
	"encoding/json"
	"errors"
)

// LanguagesList Model
type LanguageList struct {
	// Total number of languages that matched your query.
	Total int `json:"total"`
	// List of languages.
	Languages []Language `json:"languages"`

	// Used by Decode() method
	data []byte
}

func (model LanguageList) New(data []byte) *LanguageList {
	model.data = data
	return &model
}

func (model *LanguageList) UnmarshalJSON(data []byte) error {
	type alias LanguageList
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

func (model *LanguageList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
