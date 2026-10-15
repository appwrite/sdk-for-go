package models

import (
	"encoding/json"
	"errors"
)

// AppInstallation Model
type AppInstallation struct {
	// Installation ID.
	Id string `json:"$id"`
	// Installation creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Installation update time in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// ID of the installed application.
	AppId string `json:"appId"`
	// ID of the team the application is installed on.
	TeamId string `json:"teamId"`
	// Scopes granted to the application. Snapshot of the application's
	// installation scopes taken when the installation was created or last
	// updated.
	Scopes []string `json:"scopes"`
	// Authorization details granted to the application. Rich authorization
	// request (RFC 9396) style entries; the Appwrite Console stores authorized
	// project IDs here.
	AuthorizationDetails []interface{} `json:"authorizationDetails"`
	// ID of the user who created the installation.
	CreatedById string `json:"createdById"`
	// Name of the user who created the installation.
	CreatedByName string `json:"createdByName"`
	// Time an access token was last issued for the installation in ISO 8601
	// format. Null if never used.
	LastAccessedAt *string `json:"lastAccessedAt"`

	// Used by Decode() method
	data []byte
}

func (model AppInstallation) New(data []byte) *AppInstallation {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AppInstallation) UnmarshalJSON(data []byte) error {
	type alias AppInstallation
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                   json.RawMessage `json:"$id"`
		CreatedAt            json.RawMessage `json:"$createdAt"`
		UpdatedAt            json.RawMessage `json:"$updatedAt"`
		AppId                json.RawMessage `json:"appId"`
		TeamId               json.RawMessage `json:"teamId"`
		Scopes               json.RawMessage `json:"scopes"`
		AuthorizationDetails json.RawMessage `json:"authorizationDetails"`
		CreatedById          json.RawMessage `json:"createdById"`
		CreatedByName        json.RawMessage `json:"createdByName"`
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
	if present.AppId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AppId); err != nil {
			return err
		}
	}
	if present.TeamId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.TeamId); err != nil {
			return err
		}
	}
	if present.Scopes == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Scopes); err != nil {
			return err
		}
	}
	if present.AuthorizationDetails == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.AuthorizationDetails); err != nil {
			return err
		}
	}
	if present.CreatedById == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedById); err != nil {
			return err
		}
	}
	if present.CreatedByName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedByName); err != nil {
			return err
		}
	}

	*model = AppInstallation(decoded)
	return nil
}

func (model *AppInstallation) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
