package models

import (
    "encoding/json"
    "errors"
)

// Presence Model
type Presence struct {
    // Presence ID.
    Id string `json:"$id"`
    // Presence sequence ID.
    Sequence string `json:"$sequence"`
    // Presence creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Presence update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Presence permissions. [Learn more about
    // permissions](https://appwrite.io/docs/permissions).
    Permissions []string `json:"$permissions"`
    // User internal ID.
    UserInternalId string `json:"userInternalId"`
    // User ID.
    UserId string `json:"userId"`
    // Presence status.
    Status string `json:"status"`
    // Presence source.
    Source string `json:"source"`
    // Presence expiry date in ISO 8601 format.
    ExpiresAt string `json:"expiresAt"`

    // Used by Decode() method
    data []byte
}

func (model Presence) New(data []byte) *Presence {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *Presence) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}