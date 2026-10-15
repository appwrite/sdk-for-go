package models

import (
	"encoding/json"
	"errors"
)

// PolicyPasswordStrength Model
type PolicyPasswordStrength struct {
	// Policy ID.
	Id string `json:"$id"`
	// Minimum password length required for user passwords.
	Min int `json:"min"`
	// Whether passwords must include at least one uppercase letter.
	Uppercase bool `json:"uppercase"`
	// Whether passwords must include at least one lowercase letter.
	Lowercase bool `json:"lowercase"`
	// Whether passwords must include at least one number.
	Number bool `json:"number"`
	// Whether passwords must include at least one symbol.
	Symbols bool `json:"symbols"`

	// Used by Decode() method
	data []byte
}

func (model PolicyPasswordStrength) New(data []byte) *PolicyPasswordStrength {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *PolicyPasswordStrength) UnmarshalJSON(data []byte) error {
	type alias PolicyPasswordStrength
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id        json.RawMessage `json:"$id"`
		Min       json.RawMessage `json:"min"`
		Uppercase json.RawMessage `json:"uppercase"`
		Lowercase json.RawMessage `json:"lowercase"`
		Number    json.RawMessage `json:"number"`
		Symbols   json.RawMessage `json:"symbols"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.Min == nil {
		if err := json.Unmarshal([]byte("8"), &decoded.Min); err != nil {
			return err
		}
	}
	if present.Uppercase == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Uppercase); err != nil {
			return err
		}
	}
	if present.Lowercase == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Lowercase); err != nil {
			return err
		}
	}
	if present.Number == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Number); err != nil {
			return err
		}
	}
	if present.Symbols == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Symbols); err != nil {
			return err
		}
	}

	*model = PolicyPasswordStrength(decoded)
	return nil
}

func (model *PolicyPasswordStrength) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
