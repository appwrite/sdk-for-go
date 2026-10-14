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

func (model *PasskeyChallenge) UnmarshalJSON(data []byte) error {
	type alias PasskeyChallenge
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"$createdAt\":\"\",\"passkeyId\":\"\",\"expire\":\"\",\"publicKey\":[]}"), &fields); err != nil {
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
