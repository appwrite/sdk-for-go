package models

import (
    "encoding/json"
    "errors"
)

// AlgoMD5 Model
type AlgoMd5 struct {
    // Algo type.
    Type string `json:"xtype"`

    // Used by Decode() method
    data []byte
}

func (model AlgoMd5) New(data []byte) *AlgoMd5 {
    model.data = data
    return &model
}

func (model *AlgoMd5) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
