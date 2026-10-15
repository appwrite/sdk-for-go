package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Consent Model
type Oauth2Consent struct {
	// Consent ID.
	Id string `json:"$id"`
	// Consent creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Consent update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// ID of the user the consent belongs to.
	UserId string `json:"userId"`
	// ID of the registered app the consent was given to. Empty for URL-form
	// (CIMD) clients.
	AppId string `json:"appId"`
	// Client ID metadata document URL of the client the consent was given to.
	// Empty for registered apps.
	CimdUrl string `json:"cimdUrl"`
	// OAuth2 scopes the user consented to.
	Scopes []string `json:"scopes"`
	// RFC 8707 resource indicators the user consented to.
	Resources []string `json:"resources"`
	// Authorization details the user consented to, as a JSON string. Each entry
	// has a `type` plus project-defined fields.
	AuthorizationDetails string `json:"authorizationDetails"`
	// Consent expiration time in ISO 8601 format. Empty when the consent has no
	// token-bound expiry yet.
	Expire string `json:"expire"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2Consent) New(data []byte) *Oauth2Consent {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Oauth2Consent) UnmarshalJSON(data []byte) error {
	type alias Oauth2Consent
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                   json.RawMessage `json:"$id"`
		CreatedAt            json.RawMessage `json:"$createdAt"`
		UpdatedAt            json.RawMessage `json:"$updatedAt"`
		UserId               json.RawMessage `json:"userId"`
		AppId                json.RawMessage `json:"appId"`
		CimdUrl              json.RawMessage `json:"cimdUrl"`
		Scopes               json.RawMessage `json:"scopes"`
		Resources            json.RawMessage `json:"resources"`
		AuthorizationDetails json.RawMessage `json:"authorizationDetails"`
		Expire               json.RawMessage `json:"expire"`
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
	if present.AppId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AppId); err != nil {
			return err
		}
	}
	if present.CimdUrl == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CimdUrl); err != nil {
			return err
		}
	}
	if present.Scopes == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Scopes); err != nil {
			return err
		}
	}
	if present.Resources == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Resources); err != nil {
			return err
		}
	}
	if present.AuthorizationDetails == nil {
		if err := json.Unmarshal([]byte("\"[]\""), &decoded.AuthorizationDetails); err != nil {
			return err
		}
	}
	if present.Expire == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Expire); err != nil {
			return err
		}
	}

	*model = Oauth2Consent(decoded)
	return nil
}

func (model *Oauth2Consent) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
