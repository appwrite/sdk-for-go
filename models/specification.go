package models

import (
    "encoding/json"
    "errors"
)

// Specification Model
type Specification struct {
    // Memory size in MB.
    Memory int `json:"memory"`
    // Number of CPUs.
    Cpus float64 `json:"cpus"`
    // Is size enabled.
    Enabled bool `json:"enabled"`
    // Size slug.
    Slug string `json:"slug"`

    // Used by Decode() method
    data []byte
}

func (model Specification) New(data []byte) *Specification {
    model.data = data
    return &model
}

func (model *Specification) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
