package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Gitlab Model
type OAuth2Gitlab struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // GitLab OAuth2 application ID.
    ApplicationId string `json:"applicationId"`
    // GitLab OAuth2 secret.
    Secret string `json:"secret"`
    // GitLab OAuth2 endpoint URL. Defaults to https://gitlab.com for self-hosted
    // instances.
    Endpoint string `json:"endpoint"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Gitlab) New(data []byte) *OAuth2Gitlab {
    model.data = data
    return &model
}

func (model *OAuth2Gitlab) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}