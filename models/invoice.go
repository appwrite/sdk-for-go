package models

import (
    "encoding/json"
    "errors"
)

// Invoice Model
type Invoice struct {
    // Invoice ID.
    Id string `json:"$id"`
    // Invoice creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Invoice update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Invoice permissions. [Learn more about permissions](/docs/permissions).
    Permissions []string `json:"$permissions"`
    // Project ID
    TeamId string `json:"teamId"`
    // Aggregation ID
    AggregationId string `json:"aggregationId"`
    // Billing plan selected. Can be one of `tier-0`, `tier-1` or `tier-2`.
    Plan string `json:"plan"`
    // Usage breakdown per resource
    Usage []UsageResources `json:"usage"`
    // Invoice Amount
    Amount float64 `json:"amount"`
    // Tax percentage
    Tax float64 `json:"tax"`
    // Tax amount
    TaxAmount float64 `json:"taxAmount"`
    // VAT percentage
    Vat float64 `json:"vat"`
    // VAT amount
    VatAmount float64 `json:"vatAmount"`
    // Gross amount after vat, tax, and discounts applied.
    GrossAmount float64 `json:"grossAmount"`
    // Credits used.
    CreditsUsed float64 `json:"creditsUsed"`
    // Currency the invoice is in
    Currency string `json:"currency"`
    // Client secret for processing failed payments in front-end
    ClientSecret string `json:"clientSecret"`
    // Invoice status
    Status string `json:"status"`
    // Last payment error associated with the invoice
    LastError string `json:"lastError"`
    // Invoice due date.
    DueAt string `json:"dueAt"`
    // Beginning date of the invoice
    From string `json:"from"`
    // End date of the invoice
    To string `json:"to"`

    // Used by Decode() method
    data []byte
}

func (model Invoice) New(data []byte) *Invoice {
    model.data = data
    return &model
}

func (model *Invoice) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}