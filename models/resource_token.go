package models

import (
    "encoding/json"
    "errors"
)

// ResourceToken Model
type ResourceToken struct {
    // Token ID.
    Id string `json:"$id"`
    // Token creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Resource ID.
    ResourceId string `json:"resourceId"`
    // Resource type.
    ResourceType string `json:"resourceType"`
    // Token expiration date in ISO 8601 format.
    Expire string `json:"expire"`
    // JWT encoded string.
    Secret string `json:"secret"`
    // Most recent access date in ISO 8601 format. This attribute is only updated
    // again after 24 hours.
    AccessedAt string `json:"accessedAt"`

    // Used by Decode() method
    data []byte
}

func (model ResourceToken) New(data []byte) *ResourceToken {
    model.data = data
    return &model
}

func (model *ResourceToken) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}