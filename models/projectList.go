package models

import (
    "encoding/json"
    "errors"
)

// ProjectsList Model
type ProjectList struct {
    // Total number of projects that matched your query.
    Total int `json:"total"`
    // List of projects.
    Projects []Project `json:"projects"`

    // Used by Decode() method
    data []byte
}

func (model ProjectList) New(data []byte) *ProjectList {
    model.data = data
    return &model
}

func (model *ProjectList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}