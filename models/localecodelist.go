package models

import (
    "encoding/json"
    "errors"
)

// LocaleCodesList Model
type LocaleCodeList struct {
    // Total number of localeCodes documents that matched your query.
    Total int `json:"total"`
    // List of localeCodes.
    LocaleCodes []LocaleCode `json:"localeCodes"`

    // Used by Decode() method
    data []byte
}

func (model LocaleCodeList) New(data []byte) *LocaleCodeList {
    model.data = data
    return &model
}

func (model *LocaleCodeList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}