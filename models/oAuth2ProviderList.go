package models

import (
	"encoding/json"
	"errors"
)

// OAuth2ProvidersList Model
type OAuth2ProviderList struct {
	// Total number of OAuth2 providers in the given project.
	Total int `json:"total"`
	// List of OAuth2 providers.
	Providers []interface{} `json:"providers"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2ProviderList) New(data []byte) *OAuth2ProviderList {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *OAuth2ProviderList) UnmarshalJSON(data []byte) error {
	type alias OAuth2ProviderList
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

	*model = OAuth2ProviderList(decoded)
	return nil
}

func (model *OAuth2ProviderList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
