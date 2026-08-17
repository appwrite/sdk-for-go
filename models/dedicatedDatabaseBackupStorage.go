package models

import (
	"encoding/json"
	"errors"
)

// BackupStorageConfig Model
type DedicatedDatabaseBackupStorage struct {
	// Storage provider. Possible values: s3 (Amazon S3 or S3-compatible), gcs
	// (Google Cloud Storage), azure (Azure Blob Storage).
	Provider string `json:"provider"`
	// Storage bucket or container name.
	Bucket string `json:"bucket"`
	// Storage region.
	Region string `json:"region"`
	// Object key prefix for backups.
	Prefix string `json:"prefix"`
	// Custom endpoint for S3-compatible storage.
	Endpoint string `json:"endpoint"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabaseBackupStorage) New(data []byte) *DedicatedDatabaseBackupStorage {
	model.data = data
	return &model
}

func (model *DedicatedDatabaseBackupStorage) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
