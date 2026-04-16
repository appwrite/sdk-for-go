package models

import (
    "encoding/json"
    "errors"
)

// PlatformAndroid Model
type PlatformAndroid struct {
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
    // Android application ID.
    ApplicationId string `json:"applicationId"`

    // Used by Decode() method
    data []byte
}

func (model PlatformAndroid) New(data []byte) *PlatformAndroid {
    model.data = data
    return &model
}

func (model *PlatformAndroid) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}