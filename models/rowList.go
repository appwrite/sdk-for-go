package models

import (
    "encoding/json"
    "errors"
)

// RowsList Model
type RowList struct {
    // Total number of rows that matched your query.
    Total int `json:"total"`
    // List of rows.
    Rows []Row `json:"rows"`

    // Used by Decode() method
    data []byte
}

func (model RowList) New(data []byte) *RowList {
    model.data = data
    return &model
}

func (model *RowList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}