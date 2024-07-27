package models

import (
    "encoding/json"
    "errors"
)

// AlgoPHPass Model
type AlgoPhpass struct {
    // Algo type.
    Type string `json:"xtype"`

    // Used by Decode() method
    data []byte
}

func (model AlgoPhpass) New(data []byte) *AlgoPhpass {
    model.data = data
    return &model
}

func (model *AlgoPhpass) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}