package models

import (
    "encoding/json"
    "errors"
)

// BucketsList Model
type BucketList struct {
    // Total number of buckets rows that matched your query.
    Total int `json:"total"`
    // List of buckets.
    Buckets []Bucket `json:"buckets"`

    // Used by Decode() method
    data []byte
}

func (model BucketList) New(data []byte) *BucketList {
    model.data = data
    return &model
}

func (model *BucketList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}