package models

import (
    "encoding/json"
    "errors"
)

// AppSecretPlaintext Model
type AppSecretPlaintext struct {
    // Secret ID.
    Id string `json:"$id"`
    // Secret creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Secret update time in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Application ID this secret belongs to.
    AppId string `json:"appId"`
    // Application client secret. Returned only when the secret is created;
    // subsequent reads always return an empty value.
    Secret string `json:"secret"`
    // Last few characters of the client secret, used to help identify it.
    Hint string `json:"hint"`
    // ID of the user who created the secret.
    CreatedById string `json:"createdById"`
    // Name of the user who created the secret.
    CreatedByName string `json:"createdByName"`
    // Time the secret was last used for authentication in ISO 8601 format. Null
    // if never used.
    LastAccessedAt string `json:"lastAccessedAt"`

    // Used by Decode() method
    data []byte
}

func (model AppSecretPlaintext) New(data []byte) *AppSecretPlaintext {
    model.data = data
    return &model
}

func (model *AppSecretPlaintext) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}