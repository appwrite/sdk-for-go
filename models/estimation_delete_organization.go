package models

import (
    "encoding/json"
    "errors"
)

// EstimationDeleteOrganization Model
type EstimationDeleteOrganization struct {
    // List of unpaid invoices
    UnpaidInvoices []Invoice `json:"unpaidInvoices"`

    // Used by Decode() method
    data []byte
}

func (model EstimationDeleteOrganization) New(data []byte) *EstimationDeleteOrganization {
    model.data = data
    return &model
}

func (model *EstimationDeleteOrganization) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}