package models

import (
    "encoding/json"
    "errors"
)

// OAuth2ConsentTokensList Model
type Oauth2ConsentTokenList struct {
    // Total number of tokens that matched your query.
    Total int `json:"total"`
    // List of tokens.
    Tokens []Oauth2ConsentToken `json:"tokens"`

    // Used by Decode() method
    data []byte
}

func (model Oauth2ConsentTokenList) New(data []byte) *Oauth2ConsentTokenList {
    model.data = data
    return &model
}

func (model *Oauth2ConsentTokenList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}