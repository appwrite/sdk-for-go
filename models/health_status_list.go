package models

import (
    "encoding/json"
    "errors"
)

// StatusList Model
type HealthStatusList struct {
    // Total number of statuses that matched your query.
    Total int `json:"total"`
    // List of statuses.
    Statuses []HealthStatus `json:"statuses"`

    // Used by Decode() method
    data []byte
}

func (model HealthStatusList) New(data []byte) *HealthStatusList {
    model.data = data
    return &model
}

func (model *HealthStatusList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}