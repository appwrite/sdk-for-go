package models

import (
    "encoding/json"
    "errors"
)

// HealthQueue Model
type HealthQueue struct {
    // Amount of actions in the queue.
    Size int `json:"size"`

    // Used by Decode() method
    data []byte
}

func (model HealthQueue) New(data []byte) *HealthQueue {
    model.data = data
    return &model
}

func (model *HealthQueue) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
