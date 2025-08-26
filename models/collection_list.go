package models

import (
    "encoding/json"
    "errors"
)

// CollectionsList Model
type CollectionList struct {
    // Total number of collections that matched your query.
    Total int `json:"total"`
    // List of collections.
    Collections []Collection `json:"collections"`

    // Used by Decode() method
    data []byte
}

func (model CollectionList) New(data []byte) *CollectionList {
    model.data = data
    return &model
}

func (model *CollectionList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}