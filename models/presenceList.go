package models

import (
    "encoding/json"
    "errors"
)

// PresencesList Model
type PresenceList struct {
    // Total number of presences that matched your query.
    Total int `json:"total"`
    // List of presences.
    Presences []Presence `json:"presences"`

    // Used by Decode() method
    data []byte
}

func (model PresenceList) New(data []byte) *PresenceList {
    model.data = data
    return &model
}

func (model *PresenceList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}