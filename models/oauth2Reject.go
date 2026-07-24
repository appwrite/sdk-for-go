package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Reject Model
type Oauth2Reject struct {
    // URL the end user should be redirected to after the grant is rejected,
    // carrying an `access_denied` error.
    RedirectUrl string `json:"redirectUrl"`

    // Used by Decode() method
    data []byte
}

func (model Oauth2Reject) New(data []byte) *Oauth2Reject {
    model.data = data
    return &model
}

func (model *Oauth2Reject) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}