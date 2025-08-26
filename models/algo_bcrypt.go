package models

import (
    "encoding/json"
    "errors"
)

// AlgoBcrypt Model
type AlgoBcrypt struct {
    // Algo type.
    Type string `json:"type"`

    // Used by Decode() method
    data []byte
}

func (model AlgoBcrypt) New(data []byte) *AlgoBcrypt {
    model.data = data
    return &model
}

func (model *AlgoBcrypt) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}