package models

import (
    "encoding/json"
    "errors"
)

// AlgoScryptModified Model
type AlgoScryptModified struct {
    // Algo type.
    Type string `json:"xtype"`
    // Salt used to compute hash.
    Salt string `json:"salt"`
    // Separator used to compute hash.
    SaltSeparator string `json:"saltSeparator"`
    // Key used to compute hash.
    SignerKey string `json:"signerKey"`

    // Used by Decode() method
    data []byte
}

func (model AlgoScryptModified) New(data []byte) *AlgoScryptModified {
    model.data = data
    return &model
}

func (model *AlgoScryptModified) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}