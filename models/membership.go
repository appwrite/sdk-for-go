package models

import (
    "encoding/json"
    "errors"
)

// Membership Model
type Membership struct {
    // Membership ID.
    Id string `json:"$id"`
    // Membership creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Membership update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // User ID.
    UserId string `json:"userId"`
    // User name. Hide this attribute by toggling membership privacy in the
    // Console.
    UserName string `json:"userName"`
    // User email address. Hide this attribute by toggling membership privacy in
    // the Console.
    UserEmail string `json:"userEmail"`
    // Team ID.
    TeamId string `json:"teamId"`
    // Team name.
    TeamName string `json:"teamName"`
    // Date, the user has been invited to join the team in ISO 8601 format.
    Invited string `json:"invited"`
    // Date, the user has accepted the invitation to join the team in ISO 8601
    // format.
    Joined string `json:"joined"`
    // User confirmation status, true if the user has joined the team or false
    // otherwise.
    Confirm bool `json:"confirm"`
    // Multi factor authentication status, true if the user has MFA enabled or
    // false otherwise. Hide this attribute by toggling membership privacy in the
    // Console.
    Mfa bool `json:"mfa"`
    // User list of roles
    Roles []string `json:"roles"`

    // Used by Decode() method
    data []byte
}

func (model Membership) New(data []byte) *Membership {
    model.data = data
    return &model
}

func (model *Membership) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}