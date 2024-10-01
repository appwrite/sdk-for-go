package models

import (
    "encoding/json"
    "errors"
)

// Team Model
type Team struct {
    // Team ID.
    Id string `json:"$id"`
    // Team creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Team update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Team name.
    Name string `json:"name"`
    // Total number of team members.
    Total int `json:"total"`
    // Team preferences as a key-value object
    Prefs Preferences `json:"prefs"`

    // Used by Decode() method
    data []byte
}

func (model Team) New(data []byte) *Team {
    model.data = data
    return &model
}

func (model *Team) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
