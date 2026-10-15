package models

import (
	"encoding/json"
	"errors"
)

// PolicyPasskey Model
type PolicyPasskey struct {
	// Policy ID.
	Id string `json:"$id"`
	// Relying party ID passkeys are bound to. Empty until configured.
	RpId string `json:"rpId"`
	// Web origins passkeys work on, from the project's platforms on the relying
	// party ID and a localhost web platform. Read-only.
	Origins []string `json:"origins"`

	// Used by Decode() method
	data []byte
}

func (model PolicyPasskey) New(data []byte) *PolicyPasskey {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *PolicyPasskey) UnmarshalJSON(data []byte) error {
	type alias PolicyPasskey
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id      json.RawMessage `json:"$id"`
		RpId    json.RawMessage `json:"rpId"`
		Origins json.RawMessage `json:"origins"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.RpId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.RpId); err != nil {
			return err
		}
	}
	if present.Origins == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Origins); err != nil {
			return err
		}
	}

	*model = PolicyPasskey(decoded)
	return nil
}

func (model *PolicyPasskey) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
