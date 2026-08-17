package models

import (
	"encoding/json"
	"errors"
)

// BackupList Model
type DedicatedDatabaseBackupList struct {
	// Total number of backups.
	Total int `json:"total"`
	// List of backups.
	Backups []DedicatedDatabaseBackup `json:"backups"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabaseBackupList) New(data []byte) *DedicatedDatabaseBackupList {
	model.data = data
	return &model
}

func (model *DedicatedDatabaseBackupList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
