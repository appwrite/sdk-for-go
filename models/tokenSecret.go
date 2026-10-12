package models

import (
	"encoding/json"
	"errors"
)

// TokenSecret Model
type TokenSecret struct {
	// Token ID.
	Id string `json:"$id"`
	// Token creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// User ID.
	UserId string `json:"userId"`
	// Token secret key. Exchange it for a session with the create session
	// endpoint.
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

func (model TokenSecret) New(data []byte) *TokenSecret {
	model.data = data
	return &model
}

func (model *TokenSecret) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
