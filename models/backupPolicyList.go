package models

import (
	"encoding/json"
	"errors"
)

// BackupPolicyList Model
type BackupPolicyList struct {
	// Total number of policies that matched your query.
	Total int `json:"total"`
	// List of policies.
	Policies []BackupPolicy `json:"policies"`

	// Used by Decode() method
	data []byte
}

func (model BackupPolicyList) New(data []byte) *BackupPolicyList {
	model.data = data
	return &model
}

func (model *BackupPolicyList) UnmarshalJSON(data []byte) error {
	type alias BackupPolicyList
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"total\":0}"), &fields); err != nil {
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

func (model *BackupPolicyList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
