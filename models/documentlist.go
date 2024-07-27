package models

import (
    "encoding/json"
    "errors"
)

// DocumentsList Model
type DocumentList struct {
    // Total number of documents documents that matched your query.
    Total int `json:"total"`
    // List of documents.
    Documents []Document `json:"documents"`

    // Used by Decode() method
    data []byte
}

func (model DocumentList) New(data []byte) *DocumentList {
    model.data = data
    return &model
}

func (model *DocumentList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}