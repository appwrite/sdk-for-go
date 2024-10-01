package models

import (
    "encoding/json"
    "errors"
)

// LogsList Model
type LogList struct {
    // Total number of logs documents that matched your query.
    Total int `json:"total"`
    // List of logs.
    Logs []Log `json:"logs"`

    // Used by Decode() method
    data []byte
}

func (model LogList) New(data []byte) *LogList {
    model.data = data
    return &model
}

func (model *LogList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
