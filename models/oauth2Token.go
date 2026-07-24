package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Token Model
type Oauth2Token struct {
    // OAuth2 access token.
    AccessToken string `json:"access_token"`
    // OAuth2 token type.
    TokenType string `json:"token_type"`
    // Access token lifetime in seconds.
    ExpiresIn int `json:"expires_in"`
    // OAuth2 refresh token.
    RefreshToken string `json:"refresh_token"`
    // Space-separated scopes granted to the access token.
    Scope string `json:"scope"`
    // Granted RFC 9396 authorization details as a JSON string.
    AuthorizationDetails string `json:"authorization_details"`
    // OpenID Connect ID token. Returned when the `openid` scope is granted.
    IdToken string `json:"id_token"`

    // Used by Decode() method
    data []byte
}

func (model Oauth2Token) New(data []byte) *Oauth2Token {
    model.data = data
    return &model
}

func (model *Oauth2Token) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}