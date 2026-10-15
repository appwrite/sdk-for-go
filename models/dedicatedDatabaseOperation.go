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
	// Operation type, such as provision, update, credentials-update, restore,
	// pausing, resuming, failover, backup-create or cross-region-enable.
	Type string `json:"type"`
	// Operation status. Possible values: queued (accepted and waiting to resume),
	// running (in progress), completed (finished successfully), failed (ended in
	// an error).
	Status string `json:"status"`
	// Number of times this operation has been attempted.
	Attempts int `json:"attempts"`
	// Time the operation was requested, in ISO 8601 format.
	RequestedAt *string `json:"requestedAt"`
	// Time the operation started, in ISO 8601 format.
	StartedAt *string `json:"startedAt"`
	// Time the operation reached a terminal state, in ISO 8601 format.
	CompletedAt *string `json:"completedAt"`
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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DedicatedDatabaseOperation) UnmarshalJSON(data []byte) error {
	type alias DedicatedDatabaseOperation
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id           json.RawMessage `json:"$id"`
		CreatedAt    json.RawMessage `json:"$createdAt"`
		DatabaseId   json.RawMessage `json:"databaseId"`
		Type         json.RawMessage `json:"type"`
		Status       json.RawMessage `json:"status"`
		Attempts     json.RawMessage `json:"attempts"`
		ErrorCode    json.RawMessage `json:"errorCode"`
		ErrorMessage json.RawMessage `json:"errorMessage"`
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
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Type); err != nil {
			return err
		}
	}
	if present.Status == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Status); err != nil {
			return err
		}
	}
	if present.Attempts == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Attempts); err != nil {
			return err
		}
	}
	if present.ErrorCode == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ErrorCode); err != nil {
			return err
		}
	}
	if present.ErrorMessage == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ErrorMessage); err != nil {
			return err
		}
	}

	*model = DedicatedDatabaseOperation(decoded)
	return nil
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
