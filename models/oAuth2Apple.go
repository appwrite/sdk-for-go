package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Apple Model
type OAuth2Apple struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// Apple OAuth2 service ID.
	ServiceId string `json:"serviceId"`
	// Apple OAuth2 key ID.
	KeyId string `json:"keyId"`
	// Apple OAuth2 team ID.
	TeamId string `json:"teamId"`
	// Apple OAuth2 .p8 private key file contents. The secret key wrapped by the
	// PEM markers is 200 characters long.
	P8File string `json:"p8File"`
	// Native Sign in with Apple is active and can be used to create sessions from
	// an ID token. Independent of enabled, which only controls the browser-based
	// flow.
	NativeEnabled bool `json:"nativeEnabled"`
	// App bundle IDs accepted as ID token audiences for native Sign in with
	// Apple, next to the Services ID.
	NativeClientIds []string `json:"nativeClientIds"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Apple) New(data []byte) *OAuth2Apple {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *OAuth2Apple) UnmarshalJSON(data []byte) error {
	type alias OAuth2Apple
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id              json.RawMessage `json:"$id"`
		Enabled         json.RawMessage `json:"enabled"`
		ServiceId       json.RawMessage `json:"serviceId"`
		KeyId           json.RawMessage `json:"keyId"`
		TeamId          json.RawMessage `json:"teamId"`
		P8File          json.RawMessage `json:"p8File"`
		NativeEnabled   json.RawMessage `json:"nativeEnabled"`
		NativeClientIds json.RawMessage `json:"nativeClientIds"`
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
	if present.ServiceId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ServiceId); err != nil {
			return err
		}
	}
	if present.KeyId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.KeyId); err != nil {
			return err
		}
	}
	if present.TeamId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.TeamId); err != nil {
			return err
		}
	}
	if present.P8File == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.P8File); err != nil {
			return err
		}
	}
	if present.NativeEnabled == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.NativeEnabled); err != nil {
			return err
		}
	}
	if present.NativeClientIds == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.NativeClientIds); err != nil {
			return err
		}
	}

	*model = OAuth2Apple(decoded)
	return nil
}

func (model *OAuth2Apple) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
