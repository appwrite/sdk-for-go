package models

import (
    "encoding/json"
    "errors"
)

// PhonesList Model
type PhoneList struct {
    // Total number of phones that matched your query.
    Total int `json:"total"`
    // List of phones.
    Phones []Phone `json:"phones"`

    // Used by Decode() method
    data []byte
}

func (model PhoneList) New(data []byte) *PhoneList {
    model.data = data
    return &model
}

func (model *PhoneList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}