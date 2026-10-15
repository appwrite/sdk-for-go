package models

import (
	"encoding/json"
	"errors"
)

// Connections Model
type DatabaseStatusConnections struct {
	// Current number of active connections.
	Current int `json:"current"`
	// The engine's own max_connections. On a pooled database this is the backend
	// limit the pooler multiplexes onto, not the ceiling a client pool may reach
	// — that is networkMaxConnections on the database resource.
	Max int `json:"max"`

	// Used by Decode() method
	data []byte
}

func (model DatabaseStatusConnections) New(data []byte) *DatabaseStatusConnections {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DatabaseStatusConnections) UnmarshalJSON(data []byte) error {
	type alias DatabaseStatusConnections
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Current json.RawMessage `json:"current"`
		Max     json.RawMessage `json:"max"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Current == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Current); err != nil {
			return err
		}
	}
	if present.Max == nil {
		if err := json.Unmarshal([]byte("100"), &decoded.Max); err != nil {
			return err
		}
	}

	*model = DatabaseStatusConnections(decoded)
	return nil
}

func (model *DatabaseStatusConnections) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
