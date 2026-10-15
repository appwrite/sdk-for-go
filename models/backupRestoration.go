package models

import (
	"encoding/json"
	"errors"
)

// Restoration Model
type BackupRestoration struct {
	// Restoration ID.
	Id string `json:"$id"`
	// Restoration creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Restoration update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Backup archive ID.
	ArchiveId string `json:"archiveId"`
	// Backup policy ID.
	PolicyId string `json:"policyId"`
	// The status of the restoration. Possible values: pending, downloading,
	// processing, completed, failed.
	Status string `json:"status"`
	// The restoration start time. Null until the restoration starts.
	StartedAt *string `json:"startedAt"`
	// Migration ID.
	MigrationId string `json:"migrationId"`
	// The services that are backed up by this policy.
	Services []string `json:"services"`
	// The resources that are backed up by this policy.
	Resources []string `json:"resources"`
	// Restoration options as a key-value object.
	Options interface{} `json:"options"`

	// Used by Decode() method
	data []byte
}

func (model BackupRestoration) New(data []byte) *BackupRestoration {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *BackupRestoration) UnmarshalJSON(data []byte) error {
	type alias BackupRestoration
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id          json.RawMessage `json:"$id"`
		CreatedAt   json.RawMessage `json:"$createdAt"`
		UpdatedAt   json.RawMessage `json:"$updatedAt"`
		ArchiveId   json.RawMessage `json:"archiveId"`
		PolicyId    json.RawMessage `json:"policyId"`
		Status      json.RawMessage `json:"status"`
		MigrationId json.RawMessage `json:"migrationId"`
		Options     json.RawMessage `json:"options"`
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
	if present.ArchiveId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ArchiveId); err != nil {
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
	if present.Options == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Options); err != nil {
			return err
		}
	}

	*model = BackupRestoration(decoded)
	return nil
}

func (model *BackupRestoration) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
