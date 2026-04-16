package models

import (
    "encoding/json"
    "errors"
)

// AuthProvider Model
type AuthProvider struct {
    // Auth Provider.
    Key string `json:"key"`
    // Auth Provider name.
    Name string `json:"name"`
    // OAuth 2.0 application ID.
    AppId string `json:"appId"`
    // OAuth 2.0 application secret. Might be JSON string if provider requires
    // extra configuration.
    Secret string `json:"secret"`
    // Auth Provider is active and can be used to create session.
    Enabled bool `json:"enabled"`

    // Used by Decode() method
    data []byte
}

func (model AuthProvider) New(data []byte) *AuthProvider {
    model.data = data
    return &model
}

func (model *AuthProvider) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}