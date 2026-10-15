package models

import (
	"encoding/json"
	"errors"
)

// Team Model
type Team struct {
	// Team ID.
	Id string `json:"$id"`
	// Team creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Team update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Team name.
	Name string `json:"name"`
	// Total number of team members.
	Total int `json:"total"`
	// Team preferences as a key-value object
	Prefs Preferences `json:"prefs"`

	// Used by Decode() method
	data []byte
}

func (model Team) New(data []byte) *Team {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Team) UnmarshalJSON(data []byte) error {
	type alias Team
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id        json.RawMessage `json:"$id"`
		CreatedAt json.RawMessage `json:"$createdAt"`
		UpdatedAt json.RawMessage `json:"$updatedAt"`
		Name      json.RawMessage `json:"name"`
		Total     json.RawMessage `json:"total"`
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
	if present.Total == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Total); err != nil {
			return err
		}
	}

	*model = Team(decoded)
	return nil
}

func (model *Team) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
