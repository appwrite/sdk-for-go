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
    // Deprecated since 1.9.5: List of dev keys.
    DevKeys []DevKey `json:"devKeys"`
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
    // Billing limits reached
    BillingLimits BillingLimits `json:"billingLimits"`
    // OAuth2 server status
    OAuth2ServerEnabled bool `json:"oAuth2ServerEnabled"`
    // OAuth2 server authorization URL
    OAuth2ServerAuthorizationUrl string `json:"oAuth2ServerAuthorizationUrl"`
    // OAuth2 server allowed scopes
    OAuth2ServerScopes []string `json:"oAuth2ServerScopes"`
    // OAuth2 server scopes used when an authorization request omits the scope
    // parameter
    OAuth2ServerDefaultScopes []string `json:"oAuth2ServerDefaultScopes"`
    // OAuth2 server accepted RFC 9396 authorization_details types
    OAuth2ServerAuthorizationDetailsTypes []string `json:"oAuth2ServerAuthorizationDetailsTypes"`
    // OAuth2 server access token duration in seconds for confidential clients
    OAuth2ServerAccessTokenDuration int `json:"oAuth2ServerAccessTokenDuration"`
    // OAuth2 server refresh token duration in seconds for confidential clients
    OAuth2ServerRefreshTokenDuration int `json:"oAuth2ServerRefreshTokenDuration"`
    // OAuth2 server access token duration in seconds for public clients (SPAs,
    // mobile, native)
    OAuth2ServerPublicAccessTokenDuration int `json:"oAuth2ServerPublicAccessTokenDuration"`
    // OAuth2 server refresh token duration in seconds for public clients (SPAs,
    // mobile, native)
    OAuth2ServerPublicRefreshTokenDuration int `json:"oAuth2ServerPublicRefreshTokenDuration"`
    // When enabled, PKCE is required for confidential clients (server-side flows
    // using client_secret). PKCE is always required for public clients regardless
    // of this setting.
    OAuth2ServerConfidentialPkce bool `json:"oAuth2ServerConfidentialPkce"`
    // URL to your application page where users enter the device flow user code.
    // Empty when the Device Authorization Grant is not configured.
    OAuth2ServerVerificationUrl string `json:"oAuth2ServerVerificationUrl"`
    // Number of characters in the device flow user code, excluding the formatting
    // separator.
    OAuth2ServerUserCodeLength int `json:"oAuth2ServerUserCodeLength"`
    // Character set for device flow user codes: `numeric`, `alphabetic`, or
    // `alphanumeric`.
    OAuth2ServerUserCodeFormat string `json:"oAuth2ServerUserCodeFormat"`
    // Lifetime in seconds of device flow device codes and user codes.
    OAuth2ServerDeviceCodeDuration int `json:"oAuth2ServerDeviceCodeDuration"`
    // OAuth2 server discovery URL
    OAuth2ServerDiscoveryUrl string `json:"oAuth2ServerDiscoveryUrl"`

    // Used by Decode() method
    data []byte
}

func (model Project) New(data []byte) *Project {
    model.data = data
    return &model
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