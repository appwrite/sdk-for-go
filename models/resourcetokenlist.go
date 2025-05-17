package models

import (
    "encoding/json"
    "errors"
)

// ResourceTokensList Model
type ResourceTokenList struct {
    // Total number of tokens documents that matched your query.
    Total int `json:"total"`
    // List of tokens.
    Tokens []ResourceToken `json:"tokens"`

    // Used by Decode() method
    data []byte
}

func (model ResourceTokenList) New(data []byte) *ResourceTokenList {
    model.data = data
    return &model
}

func (model *ResourceTokenList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}