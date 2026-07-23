package models

import (
    "encoding/json"
    "errors"
)

// OAuth2DeviceAuthorization Model
type Oauth2DeviceAuthorization struct {
    // Device verification code used by the client to poll the token endpoint.
    DeviceCode string `json:"device_code"`
    // Short code the end user enters on the verification page.
    UserCode string `json:"user_code"`
    // URL where the end user enters the user code.
    VerificationUri string `json:"verification_uri"`
    // Verification URL with the user code prefilled as a query parameter.
    VerificationUriComplete string `json:"verification_uri_complete"`
    // Lifetime of the device code and user code in seconds.
    ExpiresIn int `json:"expires_in"`
    // Minimum polling interval for the token endpoint in seconds.
    Interval int `json:"interval"`

    // Used by Decode() method
    data []byte
}

func (model Oauth2DeviceAuthorization) New(data []byte) *Oauth2DeviceAuthorization {
    model.data = data
    return &model
}

func (model *Oauth2DeviceAuthorization) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}