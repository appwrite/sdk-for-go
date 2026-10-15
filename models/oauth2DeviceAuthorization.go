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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Oauth2DeviceAuthorization) UnmarshalJSON(data []byte) error {
	type alias Oauth2DeviceAuthorization
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		DeviceCode              json.RawMessage `json:"device_code"`
		UserCode                json.RawMessage `json:"user_code"`
		VerificationUri         json.RawMessage `json:"verification_uri"`
		VerificationUriComplete json.RawMessage `json:"verification_uri_complete"`
		ExpiresIn               json.RawMessage `json:"expires_in"`
		Interval                json.RawMessage `json:"interval"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.DeviceCode == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DeviceCode); err != nil {
			return err
		}
	}
	if present.UserCode == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserCode); err != nil {
			return err
		}
	}
	if present.VerificationUri == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.VerificationUri); err != nil {
			return err
		}
	}
	if present.VerificationUriComplete == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.VerificationUriComplete); err != nil {
			return err
		}
	}
	if present.ExpiresIn == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ExpiresIn); err != nil {
			return err
		}
	}
	if present.Interval == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Interval); err != nil {
			return err
		}
	}

	*model = Oauth2DeviceAuthorization(decoded)
	return nil
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
