package models

import (
    "encoding/json"
    "errors"
)

// SpecificationsList Model
type SpecificationList struct {
    // Total number of specifications documents that matched your query.
    Total int `json:"total"`
    // List of specifications.
    Specifications []Specification `json:"specifications"`

    // Used by Decode() method
    data []byte
}

func (model SpecificationList) New(data []byte) *SpecificationList {
    model.data = data
    return &model
}

func (model *SpecificationList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
