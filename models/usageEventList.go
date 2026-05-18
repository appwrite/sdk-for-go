package models

import (
    "encoding/json"
    "errors"
)

// UsageEventsList Model
type UsageEventList struct {
    // Total number of events that matched your query.
    Total int `json:"total"`
    // List of events.
    Events []UsageEvent `json:"events"`

    // Used by Decode() method
    data []byte
}

func (model UsageEventList) New(data []byte) *UsageEventList {
    model.data = data
    return &model
}

func (model *UsageEventList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}