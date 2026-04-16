package models

import (
    "encoding/json"
    "errors"
)

// PlatformWindows Model
type PlatformWindows struct {
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
    // Windows package identifier name.
    PackageIdentifierName string `json:"packageIdentifierName"`

    // Used by Decode() method
    data []byte
}

func (model PlatformWindows) New(data []byte) *PlatformWindows {
    model.data = data
    return &model
}

func (model *PlatformWindows) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}