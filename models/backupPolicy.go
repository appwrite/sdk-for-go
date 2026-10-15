package models

import (
	"encoding/json"
	"errors"
)

// Backup Model
type BackupPolicy struct {
	// Backup policy ID.
	Id string `json:"$id"`
	// Backup policy name.
	Name string `json:"name"`
	// Policy creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Policy update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// The services that are backed up by this policy.
	Services []string `json:"services"`
	// The resources that are backed up by this policy.
	Resources []string `json:"resources"`
	// The resource ID to backup. Set only if this policy should backup a single
	// resource.
	ResourceId *string `json:"resourceId"`
	// The resource type to backup. Set only if this policy should backup a single
	// resource.
	ResourceType *string `json:"resourceType"`
	// How many days to keep the backup before it will be automatically deleted.
	Retention int `json:"retention"`
	// Policy backup schedule in CRON format.
	Schedule string `json:"schedule"`
	// Backup type. Possible values: full (complete database snapshot),
	// incremental (changes since last backup).
	Type string `json:"type"`
	// Is this policy enabled.
	Enabled bool `json:"enabled"`

	// Used by Decode() method
	data []byte
}

func (model BackupPolicy) New(data []byte) *BackupPolicy {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *BackupPolicy) UnmarshalJSON(data []byte) error {
	type alias BackupPolicy
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id        json.RawMessage `json:"$id"`
		Name      json.RawMessage `json:"name"`
		CreatedAt json.RawMessage `json:"$createdAt"`
		UpdatedAt json.RawMessage `json:"$updatedAt"`
		Retention json.RawMessage `json:"retention"`
		Schedule  json.RawMessage `json:"schedule"`
		Type      json.RawMessage `json:"type"`
		Enabled   json.RawMessage `json:"enabled"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
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
	if present.Retention == nil {
		if err := json.Unmarshal([]byte("1"), &decoded.Retention); err != nil {
			return err
		}
	}
	if present.Schedule == nil {
		if err := json.Unmarshal([]byte("\"0 * * * *\""), &decoded.Schedule); err != nil {
			return err
		}
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"full\""), &decoded.Type); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Enabled); err != nil {
			return err
		}
	}

	*model = BackupPolicy(decoded)
	return nil
}

func (model *BackupPolicy) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
