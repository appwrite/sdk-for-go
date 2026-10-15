package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Project Model
type Oauth2Project struct {
	// Project ID.
	Id string `json:"$id"`
	// Region ID the project is deployed in.
	Region string `json:"region"`
	// API endpoint of the region the project is deployed in. Empty when the
	// region has no public hostname configured.
	Endpoint string `json:"endpoint"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2Project) New(data []byte) *Oauth2Project {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Oauth2Project) UnmarshalJSON(data []byte) error {
	type alias Oauth2Project
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id       json.RawMessage `json:"$id"`
		Region   json.RawMessage `json:"region"`
		Endpoint json.RawMessage `json:"endpoint"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.Region == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Region); err != nil {
			return err
		}
	}
	if present.Endpoint == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Endpoint); err != nil {
			return err
		}
	}

	*model = Oauth2Project(decoded)
	return nil
}

func (model *Oauth2Project) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
