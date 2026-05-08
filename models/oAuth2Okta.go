package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Okta Model
type OAuth2Okta struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // Okta OAuth2 client ID.
    ClientId string `json:"clientId"`
    // Okta OAuth2 client secret.
    ClientSecret string `json:"clientSecret"`
    // Okta OAuth2 domain.
    Domain string `json:"domain"`
    // Okta OAuth2 authorization server ID.
    AuthorizationServerId string `json:"authorizationServerId"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Okta) New(data []byte) *OAuth2Okta {
    model.data = data
    return &model
}

func (model *OAuth2Okta) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}