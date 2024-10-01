package models

import (
    "encoding/json"
    "errors"
)

// IndexesList Model
type IndexList struct {
    // Total number of indexes documents that matched your query.
    Total int `json:"total"`
    // List of indexes.
    Indexes []Index `json:"indexes"`

    // Used by Decode() method
    data []byte
}

func (model IndexList) New(data []byte) *IndexList {
    model.data = data
    return &model
}

func (model *IndexList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
