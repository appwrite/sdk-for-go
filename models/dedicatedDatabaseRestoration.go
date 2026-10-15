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
	SourceDatabaseId *string `json:"sourceDatabaseId"`
	// Project ID.
	ProjectId string `json:"projectId"`
	// Backup ID used for restoration (null for PITR).
	BackupId *string `json:"backupId"`
	// Restoration type. Possible values: backup (restore from a specific backup
	// snapshot), pitr (point-in-time recovery to a specific timestamp).
	Type string `json:"type"`
	// Where the backup was restored. Possible values: in_place (replayed over the
	// database that took the backup), clone (recovered onto a second workload
	// beside the database, which keeps serving until the clone is promoted onto
	// its endpoint), copy (recovered onto a separate database with its own
	// endpoint).
	RestoreMode string `json:"restoreMode"`
	// Restoration status. Possible values: pending (queued for processing),
	// running (currently in progress), completed (successfully finished), failed
	// (encountered an error).
	Status string `json:"status"`
	// Target time for PITR restoration in ISO 8601 format.
	TargetTime *string `json:"targetTime"`
	// Restoration start time in ISO 8601 format.
	StartedAt *string `json:"startedAt"`
	// Restoration completion time in ISO 8601 format.
	CompletedAt *string `json:"completedAt"`
	// Error message if restoration failed.
	Error string `json:"error"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabaseRestoration) New(data []byte) *DedicatedDatabaseRestoration {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DedicatedDatabaseRestoration) UnmarshalJSON(data []byte) error {
	type alias DedicatedDatabaseRestoration
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id          json.RawMessage `json:"$id"`
		CreatedAt   json.RawMessage `json:"$createdAt"`
		DatabaseId  json.RawMessage `json:"databaseId"`
		ProjectId   json.RawMessage `json:"projectId"`
		Type        json.RawMessage `json:"type"`
		RestoreMode json.RawMessage `json:"restoreMode"`
		Status      json.RawMessage `json:"status"`
		Error       json.RawMessage `json:"error"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.CreatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedAt); err != nil {
			return err
		}
	}
	if present.DatabaseId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DatabaseId); err != nil {
			return err
		}
	}
	if present.ProjectId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProjectId); err != nil {
			return err
		}
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"backup\""), &decoded.Type); err != nil {
			return err
		}
	}
	if present.RestoreMode == nil {
		if err := json.Unmarshal([]byte("\"in_place\""), &decoded.RestoreMode); err != nil {
			return err
		}
	}
	if present.Status == nil {
		if err := json.Unmarshal([]byte("\"pending\""), &decoded.Status); err != nil {
			return err
		}
	}
	if present.Error == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Error); err != nil {
			return err
		}
	}

	*model = DedicatedDatabaseRestoration(decoded)
	return nil
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
