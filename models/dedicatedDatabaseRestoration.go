package models

import (
	"encoding/json"
	"errors"
)

// Restoration Model
type DedicatedDatabaseRestoration struct {
	// Restoration ID.
	Id string `json:"$id"`
	// Restoration creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Database ID being restored into.
	DatabaseId string `json:"databaseId"`
	// Source database ID when restoring a backup into another database.
	SourceDatabaseId string `json:"sourceDatabaseId"`
	// Project ID.
	ProjectId string `json:"projectId"`
	// Backup ID used for restoration (null for PITR).
	BackupId string `json:"backupId"`
	// Restoration type. Possible values: backup (restore from a specific backup
	// snapshot), pitr (point-in-time recovery to a specific timestamp).
	Type string `json:"type"`
	// Restoration status. Possible values: pending (queued for processing),
	// running (currently in progress), completed (successfully finished), failed
	// (encountered an error).
	Status string `json:"status"`
	// Target time for PITR restoration in ISO 8601 format.
	TargetTime string `json:"targetTime"`
	// Restoration start time in ISO 8601 format.
	StartedAt string `json:"startedAt"`
	// Restoration completion time in ISO 8601 format.
	CompletedAt string `json:"completedAt"`
	// Error message if restoration failed.
	Error string `json:"error"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabaseRestoration) New(data []byte) *DedicatedDatabaseRestoration {
	model.data = data
	return &model
}

func (model *DedicatedDatabaseRestoration) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
