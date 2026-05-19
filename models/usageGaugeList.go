package models

import (
    "encoding/json"
    "errors"
)

// UsageGaugesList Model
type UsageGaugeList struct {
    // Total number of gauges that matched your query.
    Total int `json:"total"`
    // List of gauges.
    Gauges []UsageGauge `json:"gauges"`

    // Used by Decode() method
    data []byte
}

func (model UsageGaugeList) New(data []byte) *UsageGaugeList {
    model.data = data
    return &model
}

func (model *UsageGaugeList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}