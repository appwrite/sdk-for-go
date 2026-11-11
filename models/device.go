package models

import (
	"encoding/json"
	"errors"
)

// Device Model
type Device struct {
	// Device ID.
	Id string `json:"$id"`
	// Device creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Device update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Device expiration date in ISO 8601 format.
	Expire string `json:"expire"`
	// ID of the session that is currently active on the device. Empty when no
	// session is active, for example right after the active session was deleted.
	ActiveSessionId string `json:"activeSessionId"`
	// List of sessions the device holds. Switch between them to change the active
	// account.
	Sessions []Session `json:"sessions"`
	// Secret used to authenticate the device. Carried by the device cookie; only
	// included in the response if the request was made with an API key.
	Secret string `json:"secret"`

	// Used by Decode() method
	data []byte
}

func (model Device) New(data []byte) *Device {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Device) UnmarshalJSON(data []byte) error {
	type alias Device
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id              json.RawMessage `json:"$id"`
		CreatedAt       json.RawMessage `json:"$createdAt"`
		UpdatedAt       json.RawMessage `json:"$updatedAt"`
		Expire          json.RawMessage `json:"expire"`
		ActiveSessionId json.RawMessage `json:"activeSessionId"`
		Secret          json.RawMessage `json:"secret"`
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
	if present.Expire == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Expire); err != nil {
			return err
		}
	}
	if present.ActiveSessionId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ActiveSessionId); err != nil {
			return err
		}
	}
	if present.Secret == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Secret); err != nil {
			return err
		}
	}

	*model = Device(decoded)
	return nil
}

func (model *Device) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
