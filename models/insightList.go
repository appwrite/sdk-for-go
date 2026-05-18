package models

import (
    "encoding/json"
    "errors"
)

// InsightsList Model
type InsightList struct {
    // Total number of insights that matched your query.
    Total int `json:"total"`
    // List of insights.
    Insights []Insight `json:"insights"`

    // Used by Decode() method
    data []byte
}

func (model InsightList) New(data []byte) *InsightList {
    model.data = data
    return &model
}

func (model *InsightList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}