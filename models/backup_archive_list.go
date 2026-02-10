package models

import (
    "encoding/json"
    "errors"
)

// BackupArchiveList Model
type BackupArchiveList struct {
    // Total number of archives that matched your query.
    Total int `json:"total"`
    // List of archives.
    Archives []BackupArchive `json:"archives"`

    // Used by Decode() method
    data []byte
}

func (model BackupArchiveList) New(data []byte) *BackupArchiveList {
    model.data = data
    return &model
}

func (model *BackupArchiveList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}