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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Transaction) UnmarshalJSON(data []byte) error {
	type alias Transaction
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id         json.RawMessage `json:"$id"`
		CreatedAt  json.RawMessage `json:"$createdAt"`
		UpdatedAt  json.RawMessage `json:"$updatedAt"`
		Status     json.RawMessage `json:"status"`
		Operations json.RawMessage `json:"operations"`
		ExpiresAt  json.RawMessage `json:"expiresAt"`
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
	if present.Status == nil {
		if err := json.Unmarshal([]byte("\"pending\""), &decoded.Status); err != nil {
			return err
		}
	}
	if present.Operations == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Operations); err != nil {
			return err
		}
	}
	if present.ExpiresAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ExpiresAt); err != nil {
			return err
		}
	}

	*model = Transaction(decoded)
	return nil
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
