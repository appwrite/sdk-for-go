package models

import (
    "encoding/json"
    "errors"
)

// HealthStatus Model
type HealthStatus struct {
    // Name of the service.
    Name string `json:"name"`
    // Duration in milliseconds how long the health check took.
    Ping int `json:"ping"`
    // Service status. Possible values can are: `pass`, `fail`
    Status string `json:"status"`

    // Used by Decode() method
    data []byte
}

func (model HealthStatus) New(data []byte) *HealthStatus {
    model.data = data
    return &model
}

func (model *HealthStatus) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
