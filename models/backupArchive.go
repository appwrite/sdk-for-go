package models

import (
	"encoding/json"
	"errors"
)

// Archive Model
type BackupArchive struct {
	// Archive ID.
	Id string `json:"$id"`
	// Archive creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Archive update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Archive policy ID.
	PolicyId string `json:"policyId"`
	// Archive size in bytes. Null until the size is known.
	Size *int `json:"size"`
	// The status of the archive creation. Possible values: pending, processing,
	// uploading, completed, failed, skipped.
	Status string `json:"status"`
	// The backup start time. Null until the backup starts.
	StartedAt *string `json:"startedAt"`
	// Migration ID.
	MigrationId string `json:"migrationId"`
	// The services that are backed up by this archive.
	Services []string `json:"services"`
	// The resources that are backed up by this archive.
	Resources []string `json:"resources"`
	// The resource ID to backup. Set only if this archive should backup a single
	// resource.
	ResourceId *string `json:"resourceId"`
	// The resource type to backup. Set only if this archive should backup a
	// single resource.
	ResourceType *string `json:"resourceType"`

	// Used by Decode() method
	data []byte
}

func (model BackupArchive) New(data []byte) *BackupArchive {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *BackupArchive) UnmarshalJSON(data []byte) error {
	type alias BackupArchive
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id          json.RawMessage `json:"$id"`
		CreatedAt   json.RawMessage `json:"$createdAt"`
		UpdatedAt   json.RawMessage `json:"$updatedAt"`
		PolicyId    json.RawMessage `json:"policyId"`
		Status      json.RawMessage `json:"status"`
		MigrationId json.RawMessage `json:"migrationId"`
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
	if present.UpdatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UpdatedAt); err != nil {
			return err
		}
	}
	if present.PolicyId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.PolicyId); err != nil {
			return err
		}
	}
	if present.Status == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Status); err != nil {
			return err
		}
	}
	if present.MigrationId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.MigrationId); err != nil {
			return err
		}
	}

	*model = BackupArchive(decoded)
	return nil
}

func (model *BackupArchive) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
