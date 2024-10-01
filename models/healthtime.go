package models

import (
    "encoding/json"
    "errors"
)

// HealthTime Model
type HealthTime struct {
    // Current unix timestamp on trustful remote server.
    RemoteTime int `json:"remoteTime"`
    // Current unix timestamp of local server where Appwrite runs.
    LocalTime int `json:"localTime"`
    // Difference of unix remote and local timestamps in milliseconds.
    Diff int `json:"diff"`

    // Used by Decode() method
    data []byte
}

func (model HealthTime) New(data []byte) *HealthTime {
    model.data = data
    return &model
}

func (model *HealthTime) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
