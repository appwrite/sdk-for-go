package models

import (
    "encoding/json"
    "errors"
)

// PolicyPasswordPersonalData Model
type PolicyPasswordPersonalData struct {
    // Policy ID.
    Id string `json:"$id"`
    // Whether password personal data policy is enabled.
    Enabled bool `json:"enabled"`

    // Used by Decode() method
    data []byte
}

func (model PolicyPasswordPersonalData) New(data []byte) *PolicyPasswordPersonalData {
    model.data = data
    return &model
}

func (model *PolicyPasswordPersonalData) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}