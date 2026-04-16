package models

import (
    "encoding/json"
    "errors"
)

// Webhook Model
type Webhook struct {
    // Webhook ID.
    Id string `json:"$id"`
    // Webhook creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Webhook update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Webhook name.
    Name string `json:"name"`
    // Webhook URL endpoint.
    Url string `json:"url"`
    // Webhook trigger events.
    Events []string `json:"events"`
    // Indicates if SSL / TLS certificate verification is enabled.
    Tls bool `json:"tls"`
    // HTTP basic authentication username.
    AuthUsername string `json:"authUsername"`
    // HTTP basic authentication password.
    AuthPassword string `json:"authPassword"`
    // Signature key which can be used to validate incoming webhook payloads. Only
    // returned on creation and secret rotation.
    Secret string `json:"secret"`
    // Indicates if this webhook is enabled.
    Enabled bool `json:"enabled"`
    // Webhook error logs from the most recent failure.
    Logs string `json:"logs"`
    // Number of consecutive failed webhook attempts.
    Attempts int `json:"attempts"`

    // Used by Decode() method
    data []byte
}

func (model Webhook) New(data []byte) *Webhook {
    model.data = data
    return &model
}

func (model *Webhook) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}