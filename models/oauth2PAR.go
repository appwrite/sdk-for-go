package models

import (
	"encoding/json"
	"errors"
)

// OAuth2PAR Model
type Oauth2PAR struct {
	// Authorization request handle to pass to the authorize endpoint.
	RequestUri string `json:"request_uri"`
	// Lifetime of the authorization request handle in seconds.
	ExpiresIn int `json:"expires_in"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2PAR) New(data []byte) *Oauth2PAR {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Oauth2PAR) UnmarshalJSON(data []byte) error {
	type alias Oauth2PAR
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		RequestUri json.RawMessage `json:"request_uri"`
		ExpiresIn  json.RawMessage `json:"expires_in"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.RequestUri == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.RequestUri); err != nil {
			return err
		}
	}
	if present.ExpiresIn == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ExpiresIn); err != nil {
			return err
		}
	}

	*model = Oauth2PAR(decoded)
	return nil
}

func (model *Oauth2PAR) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
