package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Oidc Model
type OAuth2Oidc struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // OpenID Connect OAuth2 client ID.
    ClientId string `json:"clientId"`
    // OpenID Connect OAuth2 client secret.
    ClientSecret string `json:"clientSecret"`
    // OpenID Connect well-known configuration URL. When set, authorization,
    // token, and user info endpoints can be discovered automatically.
    WellKnownURL string `json:"wellKnownURL"`
    // OpenID Connect authorization endpoint URL.
    AuthorizationURL string `json:"authorizationURL"`
    // OpenID Connect token endpoint URL.
    TokenURL string `json:"tokenURL"`
    // OpenID Connect user info endpoint URL.
    UserInfoURL string `json:"userInfoURL"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Oidc) New(data []byte) *OAuth2Oidc {
    model.data = data
    return &model
}

func (model *OAuth2Oidc) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}