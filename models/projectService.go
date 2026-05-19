package models

import (
    "encoding/json"
    "errors"
)

// ProjectService Model
type ProjectService struct {
    // Service ID.
    Id string `json:"$id"`
    // Service status.
    Enabled bool `json:"enabled"`

    // Used by Decode() method
    data []byte
}

func (model ProjectService) New(data []byte) *ProjectService {
    model.data = data
    return &model
}

func (model *ProjectService) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}