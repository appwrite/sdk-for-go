package models

import (
    "encoding/json"
    "errors"
)

// UsageEvent Model
type UsageEvent struct {
    // The metric key.
    Metric string `json:"metric"`
    // The metric value.
    Value int `json:"value"`
    // The event timestamp.
    Time string `json:"time"`
    // The API endpoint path.
    Path string `json:"path"`
    // The HTTP method.
    Method string `json:"method"`
    // HTTP status code. Stored as string to preserve unset state (empty string =
    // not available).
    Status string `json:"status"`
    // The resource type.
    ResourceType string `json:"resourceType"`
    // The resource ID.
    ResourceId string `json:"resourceId"`
    // Country code in [ISO 3166-1](http://en.wikipedia.org/wiki/ISO_3166-1)
    // two-character format.
    CountryCode string `json:"countryCode"`
    // The user agent string.
    UserAgent string `json:"userAgent"`

    // Used by Decode() method
    data []byte
}

func (model UsageEvent) New(data []byte) *UsageEvent {
    model.data = data
    return &model
}

func (model *UsageEvent) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}