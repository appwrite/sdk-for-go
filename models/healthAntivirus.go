package models

import (
    "encoding/json"
    "errors"
)

// HealthAntivirus Model
type HealthAntivirus struct {
    // Antivirus version.
    Version string `json:"version"`
    // Antivirus status. Possible values are: `disabled`, `offline`, `online`
    Status string `json:"status"`

    // Used by Decode() method
    data []byte
}

func (model HealthAntivirus) New(data []byte) *HealthAntivirus {
    model.data = data
    return &model
}

func (model *HealthAntivirus) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}