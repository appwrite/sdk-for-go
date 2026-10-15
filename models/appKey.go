package models

import (
	"encoding/json"
	"errors"
)

// AppKey Model
type AppKey struct {
	// App key ID.
	Id string `json:"$id"`
	// App key creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// App key update time in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Application ID this app key belongs to.
	AppId string `json:"appId"`
	// App key secret.
	Secret string `json:"secret"`
	// Last few characters of the app key secret, used to help identify it.
	Hint string `json:"hint"`
	// ID of the user who created the app key.
	CreatedById string `json:"createdById"`
	// Name of the user who created the app key.
	CreatedByName string `json:"createdByName"`
	// Time the app key was last used for authentication in ISO 8601 format. Null
	// if never used.
	LastAccessedAt *string `json:"lastAccessedAt"`

	// Used by Decode() method
	data []byte
}

func (model AppKey) New(data []byte) *AppKey {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AppKey) UnmarshalJSON(data []byte) error {
	type alias AppKey
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id            json.RawMessage `json:"$id"`
		CreatedAt     json.RawMessage `json:"$createdAt"`
		UpdatedAt     json.RawMessage `json:"$updatedAt"`
		AppId         json.RawMessage `json:"appId"`
		Secret        json.RawMessage `json:"secret"`
		Hint          json.RawMessage `json:"hint"`
		CreatedById   json.RawMessage `json:"createdById"`
		CreatedByName json.RawMessage `json:"createdByName"`
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
	if present.UpdatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UpdatedAt); err != nil {
			return err
		}
	}
	if present.AppId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AppId); err != nil {
			return err
		}
	}
	if present.Secret == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Secret); err != nil {
			return err
		}
	}
	if present.Hint == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Hint); err != nil {
			return err
		}
	}
	if present.CreatedById == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedById); err != nil {
			return err
		}
	}
	if present.CreatedByName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedByName); err != nil {
			return err
		}
	}

	*model = AppKey(decoded)
	return nil
}

func (model *AppKey) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
