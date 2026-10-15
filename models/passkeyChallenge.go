package models

import (
	"encoding/json"
	"errors"
)

// PasskeyChallenge Model
type PasskeyChallenge struct {
	// Challenge ID. When signing in, pass it back with the credential to receive
	// a token.
	Id string `json:"$id"`
	// Challenge creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// ID of the passkey being registered. Empty for sign-in challenges.
	PasskeyId string `json:"passkeyId"`
	// Challenge expiration date in ISO 8601 format.
	Expire string `json:"expire"`
	// WebAuthn options in JSON form. Pass them to
	// `PublicKeyCredential.parseCreationOptionsFromJSON()` when registering, or
	// `PublicKeyCredential.parseRequestOptionsFromJSON()` when signing in.
	PublicKey interface{} `json:"publicKey"`

	// Used by Decode() method
	data []byte
}

func (model PasskeyChallenge) New(data []byte) *PasskeyChallenge {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *PasskeyChallenge) UnmarshalJSON(data []byte) error {
	type alias PasskeyChallenge
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id        json.RawMessage `json:"$id"`
		CreatedAt json.RawMessage `json:"$createdAt"`
		PasskeyId json.RawMessage `json:"passkeyId"`
		Expire    json.RawMessage `json:"expire"`
		PublicKey json.RawMessage `json:"publicKey"`
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
	if present.PasskeyId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.PasskeyId); err != nil {
			return err
		}
	}
	if present.Expire == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Expire); err != nil {
			return err
		}
	}
	if present.PublicKey == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.PublicKey); err != nil {
			return err
		}
	}

	*model = PasskeyChallenge(decoded)
	return nil
}

func (model *PasskeyChallenge) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
