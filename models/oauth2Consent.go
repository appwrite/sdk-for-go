package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Consent Model
type Oauth2Consent struct {
    // Consent ID.
    Id string `json:"$id"`
    // Consent creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Consent update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // ID of the user the consent belongs to.
    UserId string `json:"userId"`
    // ID of the registered app the consent was given to. Empty for URL-form
    // (CIMD) clients.
    AppId string `json:"appId"`
    // Client ID metadata document URL of the client the consent was given to.
    // Empty for registered apps.
    CimdUrl string `json:"cimdUrl"`
    // OAuth2 scopes the user consented to.
    Scopes []string `json:"scopes"`
    // RFC 8707 resource indicators the user consented to.
    Resources []string `json:"resources"`
    // Authorization details the user consented to, as a JSON string. Each entry
    // has a `type` plus project-defined fields.
    AuthorizationDetails string `json:"authorizationDetails"`
    // Consent expiration time in ISO 8601 format. Empty when the consent has no
    // token-bound expiry yet.
    Expire string `json:"expire"`

    // Used by Decode() method
    data []byte
}

func (model Oauth2Consent) New(data []byte) *Oauth2Consent {
    model.data = data
    return &model
}

func (model *Oauth2Consent) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}