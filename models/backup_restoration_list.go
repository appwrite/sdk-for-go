package models

import (
    "encoding/json"
    "errors"
)

// BackupRestorationList Model
type BackupRestorationList struct {
    // Total number of restorations that matched your query.
    Total int `json:"total"`
    // List of restorations.
    Restorations []BackupRestoration `json:"restorations"`

    // Used by Decode() method
    data []byte
}

func (model BackupRestorationList) New(data []byte) *BackupRestorationList {
    model.data = data
    return &model
}

func (model *BackupRestorationList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}