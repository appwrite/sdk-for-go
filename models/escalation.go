package models

import (
	"encoding/json"
	"errors"
)

// Escalation Model
type Escalation struct {
	// Escalation ID.
	Id string `json:"$id"`
	// Escalation creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Escalation update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Escalation name.
	Name string `json:"name"`
	// Who is notified of a firing alert, in order, until a member acknowledges
	// it.
	Steps []EscalationStep `json:"steps"`
	// How many more times the steps run from the first, a monitor's cooldown
	// after the last step, while nobody has acknowledged the alert.
	Repeat int `json:"repeat"`

	// Used by Decode() method
	data []byte
}

func (model Escalation) New(data []byte) *Escalation {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Escalation) UnmarshalJSON(data []byte) error {
	type alias Escalation
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id        json.RawMessage `json:"$id"`
		CreatedAt json.RawMessage `json:"$createdAt"`
		UpdatedAt json.RawMessage `json:"$updatedAt"`
		Name      json.RawMessage `json:"name"`
		Repeat    json.RawMessage `json:"repeat"`
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
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Repeat == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Repeat); err != nil {
			return err
		}
	}

	*model = Escalation(decoded)
	return nil
}

func (model *Escalation) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
