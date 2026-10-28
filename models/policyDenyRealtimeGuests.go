package models

import (
	"encoding/json"
	"errors"
)

// PolicyDenyRealtimeGuests Model
type PolicyDenyRealtimeGuests struct {
	// Policy ID.
	Id string `json:"$id"`
	// Whether guests are stopped from subscribing to realtime channels.
	Enabled bool `json:"enabled"`

	// Used by Decode() method
	data []byte
}

func (model PolicyDenyRealtimeGuests) New(data []byte) *PolicyDenyRealtimeGuests {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *PolicyDenyRealtimeGuests) UnmarshalJSON(data []byte) error {
	type alias PolicyDenyRealtimeGuests
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id      json.RawMessage `json:"$id"`
		Enabled json.RawMessage `json:"enabled"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Enabled); err != nil {
			return err
		}
	}

	*model = PolicyDenyRealtimeGuests(decoded)
	return nil
}

func (model *PolicyDenyRealtimeGuests) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
