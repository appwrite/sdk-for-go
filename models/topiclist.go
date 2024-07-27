package models

import (
    "encoding/json"
    "errors"
)

// TopicList Model
type TopicList struct {
    // Total number of topics documents that matched your query.
    Total int `json:"total"`
    // List of topics.
    Topics []Topic `json:"topics"`

    // Used by Decode() method
    data []byte
}

func (model TopicList) New(data []byte) *TopicList {
    model.data = data
    return &model
}

func (model *TopicList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}