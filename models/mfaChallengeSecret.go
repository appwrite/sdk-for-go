package models

import (
	"encoding/json"
	"errors"
)

// MFAChallengeSecret Model
type MfaChallengeSecret struct {
	// Token ID.
	Id string `json:"$id"`
	// Token creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// User ID.
	UserId string `json:"userId"`
	// Token expiration date in ISO 8601 format.
	Expire string `json:"expire"`
	// Challenge code to be delivered to the end user through a custom channel.
	Code string `json:"code"`

	// Used by Decode() method
	data []byte
}

func (model MfaChallengeSecret) New(data []byte) *MfaChallengeSecret {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *MfaChallengeSecret) UnmarshalJSON(data []byte) error {
	type alias MfaChallengeSecret
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id        json.RawMessage `json:"$id"`
		CreatedAt json.RawMessage `json:"$createdAt"`
		UserId    json.RawMessage `json:"userId"`
		Expire    json.RawMessage `json:"expire"`
		Code      json.RawMessage `json:"code"`
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
	if present.UserId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserId); err != nil {
			return err
		}
	}
	if present.Expire == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Expire); err != nil {
			return err
		}
	}
	if present.Code == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Code); err != nil {
			return err
		}
	}

	*model = MfaChallengeSecret(decoded)
	return nil
}

func (model *MfaChallengeSecret) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
