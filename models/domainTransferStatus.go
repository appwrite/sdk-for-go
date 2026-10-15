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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DomainTransferStatus) UnmarshalJSON(data []byte) error {
	type alias DomainTransferStatus
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Status    json.RawMessage `json:"status"`
		Reason    json.RawMessage `json:"reason"`
		Timestamp json.RawMessage `json:"timestamp"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Status == nil {
		if err := json.Unmarshal([]byte("\"pending_registry\""), &decoded.Status); err != nil {
			return err
		}
	}
	if present.Reason == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Reason); err != nil {
			return err
		}
	}
	if present.Timestamp == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Timestamp); err != nil {
			return err
		}
	}

	*model = DomainTransferStatus(decoded)
	return nil
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
