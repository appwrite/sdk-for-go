package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Grant Model
type Oauth2Grant struct {
	// Grant ID.
	Id string `json:"$id"`
	// Grant creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Grant update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// ID of the user the grant belongs to.
	UserId string `json:"userId"`
	// ID of the OAuth2 client (app) the grant was requested for.
	AppId string `json:"appId"`
	// Requested OAuth2 scopes the user is being asked to consent to.
	Scopes []string `json:"scopes"`
	// Requested RFC 8707 resource indicators the user is being asked to consent
	// to.
	Resources []string `json:"resources"`
	// Requested authorization_details the user is being asked to consent to, as a
	// JSON string. Each entry has a `type` plus project-defined fields.
	AuthorizationDetails string `json:"authorizationDetails"`
	// OIDC prompt directive the consent screen should honor. Space-separated list
	// of: login, consent, select_account.
	Prompt string `json:"prompt"`
	// Redirect URI the user will be sent to after the flow completes.
	RedirectUri string `json:"redirectUri"`
	// Unix timestamp of when the user last authenticated.
	AuthTime int `json:"authTime"`
	// Grant expiration time in ISO 8601 format.
	Expire string `json:"expire"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2Grant) New(data []byte) *Oauth2Grant {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Oauth2Grant) UnmarshalJSON(data []byte) error {
	type alias Oauth2Grant
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
		Scopes               json.RawMessage `json:"scopes"`
		Resources            json.RawMessage `json:"resources"`
		AuthorizationDetails json.RawMessage `json:"authorizationDetails"`
		Prompt               json.RawMessage `json:"prompt"`
		RedirectUri          json.RawMessage `json:"redirectUri"`
		AuthTime             json.RawMessage `json:"authTime"`
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
	if present.Prompt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Prompt); err != nil {
			return err
		}
	}
	if present.RedirectUri == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.RedirectUri); err != nil {
			return err
		}
	}
	if present.AuthTime == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.AuthTime); err != nil {
			return err
		}
	}
	if present.Expire == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Expire); err != nil {
			return err
		}
	}

	*model = Oauth2Grant(decoded)
	return nil
}

func (model *Oauth2Grant) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
