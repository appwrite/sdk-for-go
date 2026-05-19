package models

import (
    "encoding/json"
    "errors"
)

// ReportsList Model
type ReportList struct {
    // Total number of reports that matched your query.
    Total int `json:"total"`
    // List of reports.
    Reports []Report `json:"reports"`

    // Used by Decode() method
    data []byte
}

func (model ReportList) New(data []byte) *ReportList {
    model.data = data
    return &model
}

func (model *ReportList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}