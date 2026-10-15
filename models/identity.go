package models

import (
	"encoding/json"
	"errors"
)

// Identity Model
type Identity struct {
	// Identity ID.
	Id string `json:"$id"`
	// Identity creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Identity update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// User ID.
	UserId string `json:"userId"`
	// Identity Provider.
	Provider string `json:"provider"`
	// ID of the User in the Identity Provider.
	ProviderUid string `json:"providerUid"`
	// Email of the User in the Identity Provider.
	ProviderEmail string `json:"providerEmail"`
	// Identity Provider Access Token.
	ProviderAccessToken string `json:"providerAccessToken"`
	// The date of when the access token expires in ISO 8601 format.
	ProviderAccessTokenExpiry string `json:"providerAccessTokenExpiry"`
	// Identity Provider Refresh Token.
	ProviderRefreshToken string `json:"providerRefreshToken"`
	// Identity Provider ID token (JWT) from the most recent native sign-in. Empty
	// for identities created through the browser OAuth2 flow.
	ProviderIdToken string `json:"providerIdToken"`

	// Used by Decode() method
	data []byte
}

func (model Identity) New(data []byte) *Identity {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Identity) UnmarshalJSON(data []byte) error {
	type alias Identity
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                        json.RawMessage `json:"$id"`
		CreatedAt                 json.RawMessage `json:"$createdAt"`
		UpdatedAt                 json.RawMessage `json:"$updatedAt"`
		UserId                    json.RawMessage `json:"userId"`
		Provider                  json.RawMessage `json:"provider"`
		ProviderUid               json.RawMessage `json:"providerUid"`
		ProviderEmail             json.RawMessage `json:"providerEmail"`
		ProviderAccessToken       json.RawMessage `json:"providerAccessToken"`
		ProviderAccessTokenExpiry json.RawMessage `json:"providerAccessTokenExpiry"`
		ProviderRefreshToken      json.RawMessage `json:"providerRefreshToken"`
		ProviderIdToken           json.RawMessage `json:"providerIdToken"`
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
	if present.UserId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserId); err != nil {
			return err
		}
	}
	if present.Provider == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Provider); err != nil {
			return err
		}
	}
	if present.ProviderUid == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderUid); err != nil {
			return err
		}
	}
	if present.ProviderEmail == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderEmail); err != nil {
			return err
		}
	}
	if present.ProviderAccessToken == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderAccessToken); err != nil {
			return err
		}
	}
	if present.ProviderAccessTokenExpiry == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderAccessTokenExpiry); err != nil {
			return err
		}
	}
	if present.ProviderRefreshToken == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderRefreshToken); err != nil {
			return err
		}
	}
	if present.ProviderIdToken == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderIdToken); err != nil {
			return err
		}
	}

	*model = Identity(decoded)
	return nil
}

func (model *Identity) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
