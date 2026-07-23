package models

import (
    "encoding/json"
    "errors"
)

// AppInstallationsList Model
type AppInstallationList struct {
    // Total number of installations that matched your query.
    Total int `json:"total"`
    // List of installations.
    Installations []AppInstallation `json:"installations"`

    // Used by Decode() method
    data []byte
}

func (model AppInstallationList) New(data []byte) *AppInstallationList {
    model.data = data
    return &model
}

func (model *AppInstallationList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}