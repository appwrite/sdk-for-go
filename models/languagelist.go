package models

import (
    "encoding/json"
    "errors"
)

// LanguagesList Model
type LanguageList struct {
    // Total number of languages documents that matched your query.
    Total int `json:"total"`
    // List of languages.
    Languages []Language `json:"languages"`

    // Used by Decode() method
    data []byte
}

func (model LanguageList) New(data []byte) *LanguageList {
    model.data = data
    return &model
}

func (model *LanguageList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}