package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Paypal Model
type OAuth2Paypal struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// PayPal OAuth2 client ID.
	ClientId string `json:"clientId"`
	// PayPal OAuth2 secret key.
	SecretKey string `json:"secretKey"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Paypal) New(data []byte) *OAuth2Paypal {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *OAuth2Paypal) UnmarshalJSON(data []byte) error {
	type alias OAuth2Paypal
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id        json.RawMessage `json:"$id"`
		Enabled   json.RawMessage `json:"enabled"`
		ClientId  json.RawMessage `json:"clientId"`
		SecretKey json.RawMessage `json:"secretKey"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.ClientId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ClientId); err != nil {
			return err
		}
	}
	if present.SecretKey == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SecretKey); err != nil {
			return err
		}
	}

	*model = OAuth2Paypal(decoded)
	return nil
}

func (model *OAuth2Paypal) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
