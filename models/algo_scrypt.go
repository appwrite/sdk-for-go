package models

import (
    "encoding/json"
    "errors"
)

// AlgoScrypt Model
type AlgoScrypt struct {
    // Algo type.
    Type string `json:"type"`
    // CPU complexity of computed hash.
    CostCpu int `json:"costCpu"`
    // Memory complexity of computed hash.
    CostMemory int `json:"costMemory"`
    // Parallelization of computed hash.
    CostParallel int `json:"costParallel"`
    // Length used to compute hash.
    Length int `json:"length"`

    // Used by Decode() method
    data []byte
}

func (model AlgoScrypt) New(data []byte) *AlgoScrypt {
    model.data = data
    return &model
}

func (model *AlgoScrypt) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}