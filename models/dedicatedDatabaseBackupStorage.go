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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DedicatedDatabaseBackupStorage) UnmarshalJSON(data []byte) error {
	type alias DedicatedDatabaseBackupStorage
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Provider json.RawMessage `json:"provider"`
		Bucket   json.RawMessage `json:"bucket"`
		Region   json.RawMessage `json:"region"`
		Prefix   json.RawMessage `json:"prefix"`
		Endpoint json.RawMessage `json:"endpoint"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Provider == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Provider); err != nil {
			return err
		}
	}
	if present.Bucket == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Bucket); err != nil {
			return err
		}
	}
	if present.Region == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Region); err != nil {
			return err
		}
	}
	if present.Prefix == nil {
		if err := json.Unmarshal([]byte("\"backups/\""), &decoded.Prefix); err != nil {
			return err
		}
	}
	if present.Endpoint == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Endpoint); err != nil {
			return err
		}
	}

	*model = DedicatedDatabaseBackupStorage(decoded)
	return nil
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
