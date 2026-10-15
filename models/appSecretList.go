package models

import (
	"encoding/json"
	"errors"
)

// AppSecretsList Model
type AppSecretList struct {
	// Total number of secrets that matched your query.
	Total int `json:"total"`
	// List of secrets.
	Secrets []AppSecret `json:"secrets"`

	// Used by Decode() method
	data []byte
}

func (model AppSecretList) New(data []byte) *AppSecretList {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AppSecretList) UnmarshalJSON(data []byte) error {
	type alias AppSecretList
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

	*model = AppSecretList(decoded)
	return nil
}

func (model *AppSecretList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
