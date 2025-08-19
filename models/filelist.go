package models

import (
    "encoding/json"
    "errors"
)

// FilesList Model
type FileList struct {
    // Total number of files rows that matched your query.
    Total int `json:"total"`
    // List of files.
    Files []File `json:"files"`

    // Used by Decode() method
    data []byte
}

func (model FileList) New(data []byte) *FileList {
    model.data = data
    return &model
}

func (model *FileList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}