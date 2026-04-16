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
    // Project description.
    Description string `json:"description"`
    // Project team ID.
    TeamId string `json:"teamId"`
    // Project logo file ID.
    Logo string `json:"logo"`
    // Project website URL.
    Url string `json:"url"`
    // Company legal name.
    LegalName string `json:"legalName"`
    // Country code in [ISO 3166-1](http://en.wikipedia.org/wiki/ISO_3166-1)
    // two-character format.
    LegalCountry string `json:"legalCountry"`
    // State name.
    LegalState string `json:"legalState"`
    // City name.
    LegalCity string `json:"legalCity"`
    // Company Address.
    LegalAddress string `json:"legalAddress"`
    // Company Tax ID.
    LegalTaxId string `json:"legalTaxId"`
    // Session duration in seconds.
    AuthDuration int `json:"authDuration"`
    // Max users allowed. 0 is unlimited.
    AuthLimit int `json:"authLimit"`
    // Max sessions allowed per user. 100 maximum.
    AuthSessionsLimit int `json:"authSessionsLimit"`
    // Max allowed passwords in the history list per user. Max passwords limit
    // allowed in history is 20. Use 0 for disabling password history.
    AuthPasswordHistory int `json:"authPasswordHistory"`
    // Whether or not to check user's password against most commonly used
    // passwords.
    AuthPasswordDictionary bool `json:"authPasswordDictionary"`
    // Whether or not to check the user password for similarity with their
    // personal data.
    AuthPersonalDataCheck bool `json:"authPersonalDataCheck"`
    // Whether or not to disallow disposable email addresses during signup and
    // email updates.
    AuthDisposableEmails bool `json:"authDisposableEmails"`
    // Whether or not to require canonical email addresses during signup and email
    // updates.
    AuthCanonicalEmails bool `json:"authCanonicalEmails"`
    // Whether or not to disallow free email addresses during signup and email
    // updates.
    AuthFreeEmails bool `json:"authFreeEmails"`
    // An array of mock numbers and their corresponding verification codes (OTPs).
    AuthMockNumbers []MockNumber `json:"authMockNumbers"`
    // Whether or not to send session alert emails to users.
    AuthSessionAlerts bool `json:"authSessionAlerts"`
    // Whether or not to show user names in the teams membership response.
    AuthMembershipsUserName bool `json:"authMembershipsUserName"`
    // Whether or not to show user emails in the teams membership response.
    AuthMembershipsUserEmail bool `json:"authMembershipsUserEmail"`
    // Whether or not to show user MFA status in the teams membership response.
    AuthMembershipsMfa bool `json:"authMembershipsMfa"`
    // Whether or not all existing sessions should be invalidated on password
    // change
    AuthInvalidateSessions bool `json:"authInvalidateSessions"`
    // List of Auth Providers.
    OAuthProviders []AuthProvider `json:"oAuthProviders"`
    // List of Platforms.
    Platforms []interface{} `json:"platforms"`
    // List of Webhooks.
    Webhooks []Webhook `json:"webhooks"`
    // List of API Keys.
    Keys []Key `json:"keys"`
    // List of dev keys.
    DevKeys []DevKey `json:"devKeys"`
    // Status for custom SMTP
    SmtpEnabled bool `json:"smtpEnabled"`
    // SMTP sender name
    SmtpSenderName string `json:"smtpSenderName"`
    // SMTP sender email
    SmtpSenderEmail string `json:"smtpSenderEmail"`
    // SMTP reply to email
    SmtpReplyTo string `json:"smtpReplyTo"`
    // SMTP server host name
    SmtpHost string `json:"smtpHost"`
    // SMTP server port
    SmtpPort int `json:"smtpPort"`
    // SMTP server username
    SmtpUsername string `json:"smtpUsername"`
    // SMTP server password
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
    // Email/Password auth method status
    AuthEmailPassword bool `json:"authEmailPassword"`
    // Magic URL auth method status
    AuthUsersAuthMagicURL bool `json:"authUsersAuthMagicURL"`
    // Email (OTP) auth method status
    AuthEmailOtp bool `json:"authEmailOtp"`
    // Anonymous auth method status
    AuthAnonymous bool `json:"authAnonymous"`
    // Invites auth method status
    AuthInvites bool `json:"authInvites"`
    // JWT auth method status
    AuthJWT bool `json:"authJWT"`
    // Phone auth method status
    AuthPhone bool `json:"authPhone"`
    // Account service status
    ServiceStatusForAccount bool `json:"serviceStatusForAccount"`
    // Avatars service status
    ServiceStatusForAvatars bool `json:"serviceStatusForAvatars"`
    // Databases (legacy) service status
    ServiceStatusForDatabases bool `json:"serviceStatusForDatabases"`
    // TablesDB service status
    ServiceStatusForTablesdb bool `json:"serviceStatusForTablesdb"`
    // Locale service status
    ServiceStatusForLocale bool `json:"serviceStatusForLocale"`
    // Health service status
    ServiceStatusForHealth bool `json:"serviceStatusForHealth"`
    // Project service status
    ServiceStatusForProject bool `json:"serviceStatusForProject"`
    // Storage service status
    ServiceStatusForStorage bool `json:"serviceStatusForStorage"`
    // Teams service status
    ServiceStatusForTeams bool `json:"serviceStatusForTeams"`
    // Users service status
    ServiceStatusForUsers bool `json:"serviceStatusForUsers"`
    // VCS service status
    ServiceStatusForVcs bool `json:"serviceStatusForVcs"`
    // Sites service status
    ServiceStatusForSites bool `json:"serviceStatusForSites"`
    // Functions service status
    ServiceStatusForFunctions bool `json:"serviceStatusForFunctions"`
    // Proxy service status
    ServiceStatusForProxy bool `json:"serviceStatusForProxy"`
    // GraphQL service status
    ServiceStatusForGraphql bool `json:"serviceStatusForGraphql"`
    // Migrations service status
    ServiceStatusForMigrations bool `json:"serviceStatusForMigrations"`
    // Messaging service status
    ServiceStatusForMessaging bool `json:"serviceStatusForMessaging"`
    // REST protocol status
    ProtocolStatusForRest bool `json:"protocolStatusForRest"`
    // GraphQL protocol status
    ProtocolStatusForGraphql bool `json:"protocolStatusForGraphql"`
    // Websocket protocol status
    ProtocolStatusForWebsocket bool `json:"protocolStatusForWebsocket"`
    // Project region
    Region string `json:"region"`
    // Billing limits reached
    BillingLimits BillingLimits `json:"billingLimits"`
    // Project blocks information
    Blocks []Block `json:"blocks"`
    // Last time the project was accessed via console. Used with plan's
    // projectInactivityDays to determine if project is paused.
    ConsoleAccessedAt string `json:"consoleAccessedAt"`

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