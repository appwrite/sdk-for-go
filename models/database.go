package models

import (
	"encoding/json"
	"errors"
)

// Database Model
type Database struct {
	// Database ID.
	Id string `json:"$id"`
	// Database name.
	Name string `json:"name"`
	// Optional description for the database.
	Description string `json:"description"`
	// Database creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Database update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// If database is enabled. Can be 'enabled' or 'disabled'. When disabled, the
	// database is inaccessible to users, but remains accessible to Server SDKs
	// using API keys.
	Enabled bool `json:"enabled"`
	// Database type.
	Type string `json:"type"`
	// Dedicated database lifecycle status. Null when the database has no valid
	// dedicated backing.
	Status *string `json:"status"`
	// Underlying engine of the dedicated backing: postgresql, mysql, or mongodb.
	// A managed product (tablesdb, documentsdb, vectorsdb) reports the engine it
	// runs on, so its type and engine can differ. Null when the database has no
	// dedicated backing.
	Engine *string `json:"engine"`
	// Compute specification identifier of the dedicated backing, e.g.
	// s-2vcpu-2gb. Null when the database has no dedicated backing.
	Specification *string `json:"specification"`
	// Number of secondary high availability replicas, excluding the primary. Null
	// when backing configuration is unavailable.
	Replicas *int `json:"replicas"`
	// Error message when the dedicated backing failed. Null when the database has
	// no dedicated backing or has not failed.
	Error *string `json:"error"`
	// Container status of the dedicated backing: active, inactive, or failed (a
	// wake was declined because the backing namespace is gone). Null when the
	// database has no dedicated backing or the runtime has not reported one.
	ContainerStatus *string `json:"containerStatus"`
	// Idle-lifecycle state of the dedicated backing: active, warm, cold, or
	// hibernated. Null when the database has no dedicated backing or the runtime
	// has not reported one.
	LifecycleState *string `json:"lifecycleState"`
	// Database backup policies.
	Policies []BackupPolicy `json:"policies"`
	// Database backup archives.
	Archives []BackupArchive `json:"archives"`

	// Used by Decode() method
	data []byte
}

func (model Database) New(data []byte) *Database {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Database) UnmarshalJSON(data []byte) error {
	type alias Database
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id          json.RawMessage `json:"$id"`
		Name        json.RawMessage `json:"name"`
		Description json.RawMessage `json:"description"`
		CreatedAt   json.RawMessage `json:"$createdAt"`
		UpdatedAt   json.RawMessage `json:"$updatedAt"`
		Enabled     json.RawMessage `json:"enabled"`
		Type        json.RawMessage `json:"type"`
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
	if present.Description == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Description); err != nil {
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
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"legacy\""), &decoded.Type); err != nil {
			return err
		}
	}

	*model = Database(decoded)
	return nil
}

func (model *Database) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
