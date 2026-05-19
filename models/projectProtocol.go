package models

import (
    "encoding/json"
    "errors"
)

// ProjectProtocol Model
type ProjectProtocol struct {
    // Protocol ID.
    Id string `json:"$id"`
    // Protocol status.
    Enabled bool `json:"enabled"`

    // Used by Decode() method
    data []byte
}

func (model ProjectProtocol) New(data []byte) *ProjectProtocol {
    model.data = data
    return &model
}

func (model *ProjectProtocol) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}