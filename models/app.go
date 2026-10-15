package models

import (
	"encoding/json"
	"errors"
)

// App Model
type App struct {
	// App ID.
	Id string `json:"$id"`
	// App creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// App update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Application name.
	Name string `json:"name"`
	// Application description shown to users during OAuth2 consent.
	Description string `json:"description"`
	// Application homepage URL shown to users during OAuth2 consent.
	ClientUri string `json:"clientUri"`
	// Application logo URL shown to users during OAuth2 consent.
	LogoUri string `json:"logoUri"`
	// Application privacy policy URL shown to users during OAuth2 consent.
	PrivacyPolicyUrl string `json:"privacyPolicyUrl"`
	// Application terms of service URL shown to users during OAuth2 consent.
	TermsUrl string `json:"termsUrl"`
	// Application support or security contact emails.
	Contacts []string `json:"contacts"`
	// Application tagline shown to users during OAuth2 consent.
	Tagline string `json:"tagline"`
	// Application tags shown to users during OAuth2 consent.
	Tags []string `json:"tags"`
	// Application labels. Read-only for clients; only a server SDK using a
	// project API key can update them.
	Labels []string `json:"labels"`
	// Application image URLs shown to users during OAuth2 consent.
	Images []string `json:"images"`
	// Application support URL shown to users during OAuth2 consent.
	SupportUrl string `json:"supportUrl"`
	// Application data deletion URL shown to users during OAuth2 consent.
	DataDeletionUrl string `json:"dataDeletionUrl"`
	// List of authorized redirect URIs. These URIs can be used to redirect users
	// after they authenticate.
	RedirectUris []string `json:"redirectUris"`
	// List of authorized post-logout redirect URIs for OpenID Connect
	// RP-Initiated Logout. The logout endpoint only redirects users to URIs in
	// this list after ending their session.
	PostLogoutRedirectUris []string `json:"postLogoutRedirectUris"`
	// Whether the app is enabled or not.
	Enabled bool `json:"enabled"`
	// OAuth2 client type. `public` for SPAs, mobile, and native apps that cannot
	// keep a client secret (PKCE required); `confidential` for server-side
	// clients that authenticate with a client secret.
	Type string `json:"type"`
	// Whether this client may use the OAuth2 Device Authorization Grant (RFC
	// 8628).
	DeviceFlow bool `json:"deviceFlow"`
	// ID of team that owns the application, if owned by team. Otherwise, user ID
	// will be used.
	TeamId string `json:"teamId"`
	// ID of user who owns the application, if owned by user. Otherwise, team ID
	// will be used.
	UserId string `json:"userId"`
	// Scopes the application requests when installed on a team.
	// Organization-level and project-level scopes only.
	InstallationScopes []string `json:"installationScopes"`
	// URL users are redirected to after creating or updating an installation of
	// this application. Empty for no redirect.
	InstallationRedirectUrl string `json:"installationRedirectUrl"`
	// List of application secrets.
	Secrets []AppSecret `json:"secrets"`

	// Used by Decode() method
	data []byte
}

func (model App) New(data []byte) *App {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *App) UnmarshalJSON(data []byte) error {
	type alias App
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                      json.RawMessage `json:"$id"`
		CreatedAt               json.RawMessage `json:"$createdAt"`
		UpdatedAt               json.RawMessage `json:"$updatedAt"`
		Name                    json.RawMessage `json:"name"`
		Description             json.RawMessage `json:"description"`
		ClientUri               json.RawMessage `json:"clientUri"`
		LogoUri                 json.RawMessage `json:"logoUri"`
		PrivacyPolicyUrl        json.RawMessage `json:"privacyPolicyUrl"`
		TermsUrl                json.RawMessage `json:"termsUrl"`
		Contacts                json.RawMessage `json:"contacts"`
		Tagline                 json.RawMessage `json:"tagline"`
		Tags                    json.RawMessage `json:"tags"`
		Labels                  json.RawMessage `json:"labels"`
		Images                  json.RawMessage `json:"images"`
		SupportUrl              json.RawMessage `json:"supportUrl"`
		DataDeletionUrl         json.RawMessage `json:"dataDeletionUrl"`
		RedirectUris            json.RawMessage `json:"redirectUris"`
		PostLogoutRedirectUris  json.RawMessage `json:"postLogoutRedirectUris"`
		Enabled                 json.RawMessage `json:"enabled"`
		Type                    json.RawMessage `json:"type"`
		DeviceFlow              json.RawMessage `json:"deviceFlow"`
		TeamId                  json.RawMessage `json:"teamId"`
		UserId                  json.RawMessage `json:"userId"`
		InstallationScopes      json.RawMessage `json:"installationScopes"`
		InstallationRedirectUrl json.RawMessage `json:"installationRedirectUrl"`
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
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Description == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Description); err != nil {
			return err
		}
	}
	if present.ClientUri == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ClientUri); err != nil {
			return err
		}
	}
	if present.LogoUri == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.LogoUri); err != nil {
			return err
		}
	}
	if present.PrivacyPolicyUrl == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.PrivacyPolicyUrl); err != nil {
			return err
		}
	}
	if present.TermsUrl == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.TermsUrl); err != nil {
			return err
		}
	}
	if present.Contacts == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Contacts); err != nil {
			return err
		}
	}
	if present.Tagline == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Tagline); err != nil {
			return err
		}
	}
	if present.Tags == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Tags); err != nil {
			return err
		}
	}
	if present.Labels == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Labels); err != nil {
			return err
		}
	}
	if present.Images == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Images); err != nil {
			return err
		}
	}
	if present.SupportUrl == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SupportUrl); err != nil {
			return err
		}
	}
	if present.DataDeletionUrl == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DataDeletionUrl); err != nil {
			return err
		}
	}
	if present.RedirectUris == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.RedirectUris); err != nil {
			return err
		}
	}
	if present.PostLogoutRedirectUris == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.PostLogoutRedirectUris); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"confidential\""), &decoded.Type); err != nil {
			return err
		}
	}
	if present.DeviceFlow == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.DeviceFlow); err != nil {
			return err
		}
	}
	if present.TeamId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.TeamId); err != nil {
			return err
		}
	}
	if present.UserId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserId); err != nil {
			return err
		}
	}
	if present.InstallationScopes == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.InstallationScopes); err != nil {
			return err
		}
	}
	if present.InstallationRedirectUrl == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.InstallationRedirectUrl); err != nil {
			return err
		}
	}

	*model = App(decoded)
	return nil
}

func (model *App) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
