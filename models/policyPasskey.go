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

func (model *PolicyPasskey) UnmarshalJSON(data []byte) error {
	type alias PolicyPasskey
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"rpId\":\"\",\"origins\":[]}"), &fields); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	merged, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	return json.Unmarshal(merged, (*alias)(model))
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
