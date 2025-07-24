package models

import (
    "encoding/json"
    "errors"
)

// SessionsList Model
type SessionList struct {
    // Total number of sessions documents that matched your query.
    Total int `json:"total"`
    // List of sessions.
    Sessions []Session `json:"sessions"`

    // Used by Decode() method
    data []byte
}

func (model SessionList) New(data []byte) *SessionList {
    model.data = data
    return &model
}

func (model *SessionList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}