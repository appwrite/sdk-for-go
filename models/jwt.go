package models

import (
	"encoding/json"
	"errors"
)

// JWT Model
type Jwt struct {
	// JWT encoded string.
	Jwt string `json:"jwt"`

	// Used by Decode() method
	data []byte
}

func (model Jwt) New(data []byte) *Jwt {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Jwt) UnmarshalJSON(data []byte) error {
	type alias Jwt
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Jwt json.RawMessage `json:"jwt"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Jwt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Jwt); err != nil {
			return err
		}
	}

	*model = Jwt(decoded)
	return nil
}

func (model *Jwt) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
