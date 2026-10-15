package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Organization Model
type Oauth2Organization struct {
	// Organization ID.
	Id string `json:"$id"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2Organization) New(data []byte) *Oauth2Organization {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Oauth2Organization) UnmarshalJSON(data []byte) error {
	type alias Oauth2Organization
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id json.RawMessage `json:"$id"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}

	*model = Oauth2Organization(decoded)
	return nil
}

func (model *Oauth2Organization) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
