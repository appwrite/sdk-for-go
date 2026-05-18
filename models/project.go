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
    // List of auth methods.
    AuthMethods []ProjectAuthMethod `json:"authMethods"`
    // List of services.
    Services []ProjectService `json:"services"`
    // List of protocols.
    Protocols []ProjectProtocol `json:"protocols"`
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