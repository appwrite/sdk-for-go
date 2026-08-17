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
