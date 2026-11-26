package models

import (
	"encoding/json"
	"errors"
)

// EscalationStep Model
type EscalationStep struct {
	// Seconds to wait before notifying this step: after the alert fires for the
	// first step, after the previous step for the others.
	Delay int `json:"delay"`
	// Organization roles whose confirmed members this step notifies.
	Roles []string `json:"roles"`
	// IDs of the organization's members this step notifies, whatever their role.
	UserIds []string `json:"userIds"`

	// Used by Decode() method
	data []byte
}

func (model EscalationStep) New(data []byte) *EscalationStep {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *EscalationStep) UnmarshalJSON(data []byte) error {
	type alias EscalationStep
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Delay   json.RawMessage `json:"delay"`
		Roles   json.RawMessage `json:"roles"`
		UserIds json.RawMessage `json:"userIds"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Delay == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Delay); err != nil {
			return err
		}
	}
	if present.Roles == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Roles); err != nil {
			return err
		}
	}
	if present.UserIds == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.UserIds); err != nil {
			return err
		}
	}

	*model = EscalationStep(decoded)
	return nil
}

func (model *EscalationStep) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
