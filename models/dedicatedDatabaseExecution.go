package models

import (
	"encoding/json"
	"errors"
)

// Execution Model
type DedicatedDatabaseExecution struct {
	// Result rows as a list of column-name => value maps. Empty for non-returning
	// statements.
	Rows []interface{} `json:"rows"`
	// Number of rows returned (for SELECT) or affected (for
	// INSERT/UPDATE/DELETE).
	RowCount int `json:"rowCount"`
	// Column metadata in result-set order.
	Columns []DedicatedDatabaseExecutionColumn `json:"columns"`
	// Server-side execution time in milliseconds.
	DurationMs int `json:"durationMs"`
	// True when the configured row or byte cap was hit and the result was
	// truncated.
	Truncated bool `json:"truncated"`
	// Serialised payload size in bytes.
	Bytes int `json:"bytes"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabaseExecution) New(data []byte) *DedicatedDatabaseExecution {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DedicatedDatabaseExecution) UnmarshalJSON(data []byte) error {
	type alias DedicatedDatabaseExecution
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Rows       json.RawMessage `json:"rows"`
		RowCount   json.RawMessage `json:"rowCount"`
		DurationMs json.RawMessage `json:"durationMs"`
		Truncated  json.RawMessage `json:"truncated"`
		Bytes      json.RawMessage `json:"bytes"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Rows == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Rows); err != nil {
			return err
		}
	}
	if present.RowCount == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.RowCount); err != nil {
			return err
		}
	}
	if present.DurationMs == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.DurationMs); err != nil {
			return err
		}
	}
	if present.Truncated == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Truncated); err != nil {
			return err
		}
	}
	if present.Bytes == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Bytes); err != nil {
			return err
		}
	}

	*model = DedicatedDatabaseExecution(decoded)
	return nil
}

func (model *DedicatedDatabaseExecution) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
