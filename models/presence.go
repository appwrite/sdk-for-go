package models

import (
	"encoding/json"
	"errors"
)

// Presence Model
type Presence struct {
	// Presence ID.
	Id string `json:"$id"`
	// Presence creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Presence update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Presence permissions. [Learn more about
	// permissions](https://appwrite.io/docs/permissions).
	Permissions []string `json:"$permissions"`
	// User ID.
	UserId string `json:"userId"`
	// Presence status.
	Status *string `json:"status"`
	// Presence source.
	Source string `json:"source"`
	// Presence expiry date in ISO 8601 format.
	ExpiresAt *string `json:"expiresAt"`
	// Presence metadata.
	Metadata interface{} `json:"metadata"`

	// Used by Decode() method
	data []byte
}

func (model Presence) New(data []byte) *Presence {
	model.data = data
	return &model
}

func (model *Presence) UnmarshalJSON(data []byte) error {
	type alias Presence
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"$createdAt\":\"\",\"$updatedAt\":\"\",\"userId\":\"\",\"source\":\"\"}"), &fields); err != nil {
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

func (model *Presence) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
