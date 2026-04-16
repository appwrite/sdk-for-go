package models

import (
    "encoding/json"
    "errors"
)

// BackupPolicyList Model
type BackupPolicyList struct {
    // Total number of policies that matched your query.
    Total int `json:"total"`
    // List of policies.
    Policies []BackupPolicy `json:"policies"`

    // Used by Decode() method
    data []byte
}

func (model BackupPolicyList) New(data []byte) *BackupPolicyList {
    model.data = data
    return &model
}

func (model *BackupPolicyList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}