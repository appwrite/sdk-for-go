package models

import (
    "encoding/json"
    "errors"
)

// UsageResource Model
type UsageResources struct {
    // Invoice name
    Name string `json:"name"`
    // Invoice value
    Value int `json:"value"`
    // Invoice amount
    Amount float64 `json:"amount"`
    // Invoice rate
    Rate float64 `json:"rate"`
    // Invoice description
    Desc string `json:"desc"`
    // Resource ID
    ResourceId string `json:"resourceId"`

    // Used by Decode() method
    data []byte
}

func (model UsageResources) New(data []byte) *UsageResources {
    model.data = data
    return &model
}

func (model *UsageResources) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}