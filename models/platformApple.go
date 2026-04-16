package models

import (
    "encoding/json"
    "errors"
)

// PlatformApple Model
type PlatformApple struct {
    // Platform ID.
    Id string `json:"$id"`
    // Platform creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Platform update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Platform name.
    Name string `json:"name"`
    // Platform type. Possible values are: windows, apple, android, linux, web.
    Type string `json:"type"`
    // Apple bundle identifier.
    BundleIdentifier string `json:"bundleIdentifier"`

    // Used by Decode() method
    data []byte
}

func (model PlatformApple) New(data []byte) *PlatformApple {
    model.data = data
    return &model
}

func (model *PlatformApple) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}