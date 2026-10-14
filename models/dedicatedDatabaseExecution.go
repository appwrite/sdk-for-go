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

func (model *DedicatedDatabaseExecution) UnmarshalJSON(data []byte) error {
	type alias DedicatedDatabaseExecution
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"rows\":[],\"rowCount\":0,\"durationMs\":0,\"truncated\":false,\"bytes\":0}"), &fields); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	merged, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	return json.Unmarshal(merged, (*alias)(model))
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
