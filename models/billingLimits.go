package models

import (
    "encoding/json"
    "errors"
)

// BillingLimits Model
type BillingLimits struct {
    // Bandwidth limit
    Bandwidth int `json:"bandwidth"`
    // Storage limit
    Storage int `json:"storage"`
    // Users limit
    Users int `json:"users"`
    // Executions limit
    Executions int `json:"executions"`
    // GBHours limit
    GBHours int `json:"GBHours"`
    // Image transformations limit
    ImageTransformations int `json:"imageTransformations"`
    // Auth phone limit
    AuthPhone int `json:"authPhone"`
    // Budget limit percentage
    BudgetLimit int `json:"budgetLimit"`

    // Used by Decode() method
    data []byte
}

func (model BillingLimits) New(data []byte) *BillingLimits {
    model.data = data
    return &model
}

func (model *BillingLimits) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}