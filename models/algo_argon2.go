package models

import (
    "encoding/json"
    "errors"
)

// AlgoArgon2 Model
type AlgoArgon2 struct {
    // Algo type.
    Type string `json:"type"`
    // Memory used to compute hash.
    MemoryCost int `json:"memoryCost"`
    // Amount of time consumed to compute hash
    TimeCost int `json:"timeCost"`
    // Number of threads used to compute hash.
    Threads int `json:"threads"`

    // Used by Decode() method
    data []byte
}

func (model AlgoArgon2) New(data []byte) *AlgoArgon2 {
    model.data = data
    return &model
}

func (model *AlgoArgon2) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}