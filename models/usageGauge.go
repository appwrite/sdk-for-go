package models

import (
    "encoding/json"
    "errors"
)

// UsageGauge Model
type UsageGauge struct {
    // The metric key.
    Metric string `json:"metric"`
    // The current snapshot value.
    Value int `json:"value"`
    // The snapshot timestamp.
    Time string `json:"time"`

    // Used by Decode() method
    data []byte
}

func (model UsageGauge) New(data []byte) *UsageGauge {
    model.data = data
    return &model
}

func (model *UsageGauge) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}