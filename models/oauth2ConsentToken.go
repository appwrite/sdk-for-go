package models

import (
	"encoding/json"
	"errors"
)

// OAuth2ConsentToken Model
type Oauth2ConsentToken struct {
	// Token family ID.
	Id string `json:"$id"`
	// Token creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Token update date in ISO 8601 format. Refreshing the token family updates
	// this.
	UpdatedAt string `json:"$updatedAt"`
	// ID of the consent the token family was issued under.
	ConsentId string `json:"consentId"`
	// ID of the user the token family belongs to.
	UserId string `json:"userId"`
	// ID of the registered app the token family was issued to. Empty for URL-form
	// (CIMD) clients.
	AppId string `json:"appId"`
	// Client ID metadata document URL of the client the token family was issued
	// to. Empty for registered apps.
	CimdUrl string `json:"cimdUrl"`
	// OAuth2 scopes granted on the token family.
	Scopes []string `json:"scopes"`
	// RFC 8707 resource indicators granted on the token family.
	Resources []string `json:"resources"`
	// Authorization details granted on the token family, as a JSON string. Each
	// entry has a `type` plus project-defined fields.
	AuthorizationDetails string `json:"authorizationDetails"`
	// Expiration time of the current access token of this family in ISO 8601
	// format.
	Expire string `json:"expire"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2ConsentToken) New(data []byte) *Oauth2ConsentToken {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Oauth2ConsentToken) UnmarshalJSON(data []byte) error {
	type alias Oauth2ConsentToken
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                   json.RawMessage `json:"$id"`
		CreatedAt            json.RawMessage `json:"$createdAt"`
		UpdatedAt            json.RawMessage `json:"$updatedAt"`
		ConsentId            json.RawMessage `json:"consentId"`
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
	if present.ConsentId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ConsentId); err != nil {
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

	*model = Oauth2ConsentToken(decoded)
	return nil
}

func (model *Oauth2ConsentToken) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
