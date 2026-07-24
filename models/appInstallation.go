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
    AuthorizationDetails interface{} `json:"authorizationDetails"`
    // ID of the user who created the installation.
    CreatedById string `json:"createdById"`
    // Name of the user who created the installation.
    CreatedByName string `json:"createdByName"`
    // Time an access token was last issued for the installation in ISO 8601
    // format. Null if never used.
    LastAccessedAt string `json:"lastAccessedAt"`

    // Used by Decode() method
    data []byte
}

func (model AppInstallation) New(data []byte) *AppInstallation {
    model.data = data
    return &model
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