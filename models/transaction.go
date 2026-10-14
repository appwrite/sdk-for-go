package models

import (
	"encoding/json"
	"errors"
)

// Transaction Model
type Transaction struct {
	// Transaction ID.
	Id string `json:"$id"`
	// Transaction creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Transaction update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Current status of the transaction. One of: pending, committing, committed,
	// rolled_back, failed.
	Status string `json:"status"`
	// Number of operations in the transaction.
	Operations int `json:"operations"`
	// Expiration time in ISO 8601 format.
	ExpiresAt string `json:"expiresAt"`

	// Used by Decode() method
	data []byte
}

func (model Transaction) New(data []byte) *Transaction {
	model.data = data
	return &model
}

func (model *Transaction) UnmarshalJSON(data []byte) error {
	type alias Transaction
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"$createdAt\":\"\",\"$updatedAt\":\"\",\"status\":\"pending\",\"operations\":0,\"expiresAt\":\"\"}"), &fields); err != nil {
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

func (model *Transaction) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
