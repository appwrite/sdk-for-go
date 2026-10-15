package models

import (
	"encoding/json"
	"errors"
)

// Project Model
type Project struct {
	// Project ID.
	Id string `json:"$id"`
	// Project creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Project update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Project name.
	Name string `json:"name"`
	// Project team ID.
	TeamId string `json:"teamId"`
	// Project region
	Region string `json:"region"`
	// Status for custom SMTP
	SmtpEnabled bool `json:"smtpEnabled"`
	// SMTP sender name
	SmtpSenderName string `json:"smtpSenderName"`
	// SMTP sender email
	SmtpSenderEmail string `json:"smtpSenderEmail"`
	// SMTP reply to name
	SmtpReplyToName string `json:"smtpReplyToName"`
	// SMTP reply to email
	SmtpReplyToEmail string `json:"smtpReplyToEmail"`
	// SMTP server host name
	SmtpHost string `json:"smtpHost"`
	// SMTP server port
	SmtpPort int `json:"smtpPort"`
	// SMTP server username
	SmtpUsername string `json:"smtpUsername"`
	// SMTP server password. This property is write-only and always returned
	// empty.
	SmtpPassword string `json:"smtpPassword"`
	// SMTP server secure protocol
	SmtpSecure string `json:"smtpSecure"`
	// Number of times the ping was received for this project.
	PingCount int `json:"pingCount"`
	// Last ping datetime in ISO 8601 format.
	PingedAt string `json:"pingedAt"`
	// Labels for the project.
	Labels []string `json:"labels"`
	// Project status
	Status string `json:"status"`
	// Stage progress (completed or skipped) with timestamps and actor types,
	// keyed by stage id.
	Onboarding interface{} `json:"onboarding"`
	// List of auth methods.
	AuthMethods []ProjectAuthMethod `json:"authMethods"`
	// List of services.
	Services []ProjectService `json:"services"`
	// List of protocols.
	Protocols []ProjectProtocol `json:"protocols"`
	// Project blocks information
	Blocks []Block `json:"blocks"`
	// Last time the project was accessed via console. Used with plan's
	// projectInactivityDays to determine if project is paused.
	ConsoleAccessedAt string `json:"consoleAccessedAt"`
	// Whether WAF enforcement is enabled for the project.
	WafEnabled *bool `json:"wafEnabled"`
	// ID of the project this environment belongs to. Null when the project is not
	// an environment.
	ParentId *string `json:"parentId"`
	// Billing limits reached
	BillingLimits BillingLimits `json:"billingLimits"`
	// First time the project received a non-console API request. Empty until the
	// project is used.
	FirstAccessedAt *string `json:"firstAccessedAt"`
	// Last time the project was accessed through the MCP server. Empty if it was
	// never accessed via MCP.
	McpAccessedAt *string `json:"mcpAccessedAt"`
	// OAuth2 server status
	OAuth2ServerEnabled *bool `json:"oAuth2ServerEnabled"`
	// OAuth2 server authorization URL
	OAuth2ServerAuthorizationUrl *string `json:"oAuth2ServerAuthorizationUrl"`
	// OAuth2 server allowed scopes
	OAuth2ServerScopes []string `json:"oAuth2ServerScopes"`
	// OAuth2 server scopes used when an authorization request omits the scope
	// parameter
	OAuth2ServerDefaultScopes []string `json:"oAuth2ServerDefaultScopes"`
	// Scopes an application may request when installed on a team
	OAuth2ServerInstallationScopes []string `json:"oAuth2ServerInstallationScopes"`
	// OAuth2 server accepted RFC 9396 authorization_details types
	OAuth2ServerAuthorizationDetailsTypes []string `json:"oAuth2ServerAuthorizationDetailsTypes"`
	// OAuth2 server access token duration in seconds for confidential clients
	OAuth2ServerAccessTokenDuration *int `json:"oAuth2ServerAccessTokenDuration"`
	// OAuth2 server refresh token duration in seconds for confidential clients
	OAuth2ServerRefreshTokenDuration *int `json:"oAuth2ServerRefreshTokenDuration"`
	// OAuth2 server access token duration in seconds for public clients (SPAs,
	// mobile, native)
	OAuth2ServerPublicAccessTokenDuration *int `json:"oAuth2ServerPublicAccessTokenDuration"`
	// OAuth2 server refresh token duration in seconds for public clients (SPAs,
	// mobile, native)
	OAuth2ServerPublicRefreshTokenDuration *int `json:"oAuth2ServerPublicRefreshTokenDuration"`
	// OAuth2 server access token duration in seconds for app installation access
	// tokens
	OAuth2ServerInstallationAccessTokenDuration *int `json:"oAuth2ServerInstallationAccessTokenDuration"`
	// When enabled, PKCE is required for confidential clients (server-side flows
	// using client_secret). PKCE is always required for public clients regardless
	// of this setting.
	OAuth2ServerConfidentialPkce *bool `json:"oAuth2ServerConfidentialPkce"`
	// URL to your application page where users enter the device flow user code.
	// Empty when the Device Authorization Grant is not configured.
	OAuth2ServerVerificationUrl *string `json:"oAuth2ServerVerificationUrl"`
	// Number of characters in the device flow user code, excluding the formatting
	// separator.
	OAuth2ServerUserCodeLength *int `json:"oAuth2ServerUserCodeLength"`
	// Character set for device flow user codes: `numeric`, `alphabetic`, or
	// `alphanumeric`.
	OAuth2ServerUserCodeFormat *string `json:"oAuth2ServerUserCodeFormat"`
	// Lifetime in seconds of device flow device codes and user codes.
	OAuth2ServerDeviceCodeDuration *int `json:"oAuth2ServerDeviceCodeDuration"`
	// OAuth2 server discovery URL
	OAuth2ServerDiscoveryUrl *string `json:"oAuth2ServerDiscoveryUrl"`

	// Used by Decode() method
	data []byte
}

func (model Project) New(data []byte) *Project {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Project) UnmarshalJSON(data []byte) error {
	type alias Project
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                json.RawMessage `json:"$id"`
		CreatedAt         json.RawMessage `json:"$createdAt"`
		UpdatedAt         json.RawMessage `json:"$updatedAt"`
		Name              json.RawMessage `json:"name"`
		TeamId            json.RawMessage `json:"teamId"`
		Region            json.RawMessage `json:"region"`
		SmtpEnabled       json.RawMessage `json:"smtpEnabled"`
		SmtpSenderName    json.RawMessage `json:"smtpSenderName"`
		SmtpSenderEmail   json.RawMessage `json:"smtpSenderEmail"`
		SmtpReplyToName   json.RawMessage `json:"smtpReplyToName"`
		SmtpReplyToEmail  json.RawMessage `json:"smtpReplyToEmail"`
		SmtpHost          json.RawMessage `json:"smtpHost"`
		SmtpPort          json.RawMessage `json:"smtpPort"`
		SmtpUsername      json.RawMessage `json:"smtpUsername"`
		SmtpPassword      json.RawMessage `json:"smtpPassword"`
		SmtpSecure        json.RawMessage `json:"smtpSecure"`
		PingCount         json.RawMessage `json:"pingCount"`
		PingedAt          json.RawMessage `json:"pingedAt"`
		Labels            json.RawMessage `json:"labels"`
		Status            json.RawMessage `json:"status"`
		Onboarding        json.RawMessage `json:"onboarding"`
		ConsoleAccessedAt json.RawMessage `json:"consoleAccessedAt"`
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
	if present.TeamId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.TeamId); err != nil {
			return err
		}
	}
	if present.Region == nil {
		if err := json.Unmarshal([]byte("\"fra\""), &decoded.Region); err != nil {
			return err
		}
	}
	if present.SmtpEnabled == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.SmtpEnabled); err != nil {
			return err
		}
	}
	if present.SmtpSenderName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SmtpSenderName); err != nil {
			return err
		}
	}
	if present.SmtpSenderEmail == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SmtpSenderEmail); err != nil {
			return err
		}
	}
	if present.SmtpReplyToName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SmtpReplyToName); err != nil {
			return err
		}
	}
	if present.SmtpReplyToEmail == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SmtpReplyToEmail); err != nil {
			return err
		}
	}
	if present.SmtpHost == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SmtpHost); err != nil {
			return err
		}
	}
	if present.SmtpPort == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.SmtpPort); err != nil {
			return err
		}
	}
	if present.SmtpUsername == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SmtpUsername); err != nil {
			return err
		}
	}
	if present.SmtpPassword == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SmtpPassword); err != nil {
			return err
		}
	}
	if present.SmtpSecure == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SmtpSecure); err != nil {
			return err
		}
	}
	if present.PingCount == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.PingCount); err != nil {
			return err
		}
	}
	if present.PingedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.PingedAt); err != nil {
			return err
		}
	}
	if present.Labels == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Labels); err != nil {
			return err
		}
	}
	if present.Status == nil {
		if err := json.Unmarshal([]byte("\"active\""), &decoded.Status); err != nil {
			return err
		}
	}
	if present.Onboarding == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Onboarding); err != nil {
			return err
		}
	}
	if present.ConsoleAccessedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ConsoleAccessedAt); err != nil {
			return err
		}
	}

	*model = Project(decoded)
	return nil
}

func (model *Project) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
