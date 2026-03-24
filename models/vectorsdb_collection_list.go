package models

import (
    "encoding/json"
    "errors"
)

// VectorsDBCollectionsList Model
type VectorsdbCollectionList struct {
    // Total number of collections that matched your query.
    Total int `json:"total"`
    // List of collections.
    Collections []VectorsdbCollection `json:"collections"`

    // Used by Decode() method
    data []byte
}

func (model VectorsdbCollectionList) New(data []byte) *VectorsdbCollectionList {
    model.data = data
    return &model
}

func (model *VectorsdbCollectionList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}