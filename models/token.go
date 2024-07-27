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