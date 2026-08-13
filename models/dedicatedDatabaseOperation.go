package models

import (
    "encoding/json"
    "errors"
)

// Operation Model
type DedicatedDatabaseOperation struct {
    // Operation ID.
    Id string `json:"$id"`
    // Operation creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Database ID the operation ran against.
    DatabaseId string `json:"databaseId"`
    // Operation type, such as provision, update, restore, pausing, resuming,
    // failover, backup-create or cross-region-enable.
    Type string `json:"type"`
    // Operation status. Possible values: running (in progress), completed
    // (finished successfully), failed (ended in an error).
    Status string `json:"status"`
    // Number of times this operation has been attempted.
    Attempts int `json:"attempts"`
    // Time the operation was requested, in ISO 8601 format.
    RequestedAt string `json:"requestedAt"`
    // Time the operation started, in ISO 8601 format.
    StartedAt string `json:"startedAt"`
    // Time the operation reached a terminal state, in ISO 8601 format.
    CompletedAt string `json:"completedAt"`
    // Machine-readable failure code. `Interrupted` marks an attempt that ended
    // before its outcome could be confirmed.
    ErrorCode string `json:"errorCode"`
    // Failure message if the operation failed.
    ErrorMessage string `json:"errorMessage"`

    // Used by Decode() method
    data []byte
}

func (model DedicatedDatabaseOperation) New(data []byte) *DedicatedDatabaseOperation {
    model.data = data
    return &model
}

func (model *DedicatedDatabaseOperation) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}