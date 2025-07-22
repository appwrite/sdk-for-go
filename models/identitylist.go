package models

import (
    "encoding/json"
    "errors"
)

// IdentitiesList Model
type IdentityList struct {
    // Total number of identities rows that matched your query.
    Total int `json:"total"`
    // List of identities.
    Identities []Identity `json:"identities"`

    // Used by Decode() method
    data []byte
}

func (model IdentityList) New(data []byte) *IdentityList {
    model.data = data
    return &model
}

func (model *IdentityList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}