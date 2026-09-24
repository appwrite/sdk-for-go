package models

import (
	"encoding/json"
	"errors"
)

// DomainTransferStatus Model
type DomainTransferStatus struct {
	// Transfer status.
	Status string `json:"status"`
	// Additional transfer status information.
	Reason string `json:"reason"`
	// Transfer status timestamp in ISO 8601 format.
	Timestamp string `json:"timestamp"`

	// Used by Decode() method
	data []byte
}

func (model DomainTransferStatus) New(data []byte) *DomainTransferStatus {
	model.data = data
	return &model
}

func (model *DomainTransferStatus) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
