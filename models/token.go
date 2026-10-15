package models

import (
	"encoding/json"
	"errors"
)

// Token Model
type Token struct {
	// Token ID.
	Id string `json:"$id"`
	// Token creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// User ID.
	UserId string `json:"userId"`
	// Token secret key. This will return an empty string unless the response is
	// returned using an API key or as part of a webhook payload.
	Secret string `json:"secret"`
	// Token expiration date in ISO 8601 format.
	Expire string `json:"expire"`
	// Security phrase of a token. Empty if security phrase was not requested when
	// creating a token. It includes randomly generated phrase which is also sent
	// in the external resource such as email.
	Phrase string `json:"phrase"`

	// Used by Decode() method
	data []byte
}

func (model Token) New(data []byte) *Token {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Token) UnmarshalJSON(data []byte) error {
	type alias Token
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id        json.RawMessage `json:"$id"`
		CreatedAt json.RawMessage `json:"$createdAt"`
		UserId    json.RawMessage `json:"userId"`
		Secret    json.RawMessage `json:"secret"`
		Expire    json.RawMessage `json:"expire"`
		Phrase    json.RawMessage `json:"phrase"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.CreatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedAt); err != nil {
			return err
		}
	}
	if present.UserId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserId); err != nil {
			return err
		}
	}
	if present.Secret == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Secret); err != nil {
			return err
		}
	}
	if present.Expire == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Expire); err != nil {
			return err
		}
	}
	if present.Phrase == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Phrase); err != nil {
			return err
		}
	}

	*model = Token(decoded)
	return nil
}

func (model *Token) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
