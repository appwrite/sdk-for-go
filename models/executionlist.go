package models

import (
    "encoding/json"
    "errors"
)

// ExecutionsList Model
type ExecutionList struct {
    // Total number of executions documents that matched your query.
    Total int `json:"total"`
    // List of executions.
    Executions []Execution `json:"executions"`

    // Used by Decode() method
    data []byte
}

func (model ExecutionList) New(data []byte) *ExecutionList {
    model.data = data
    return &model
}

func (model *ExecutionList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}