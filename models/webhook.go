package models

import (
	"encoding/json"
	"errors"
)

// Webhook Model
type Webhook struct {
	// Webhook ID.
	Id string `json:"$id"`
	// Webhook creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Webhook update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Webhook name.
	Name string `json:"name"`
	// Optional description for the webhook.
	Description string `json:"description"`
	// Webhook URL endpoint.
	Url string `json:"url"`
	// Webhook trigger events.
	Events []string `json:"events"`
	// Indicates if SSL / TLS certificate verification is enabled.
	Tls bool `json:"tls"`
	// HTTP basic authentication username.
	AuthUsername string `json:"authUsername"`
	// HTTP basic authentication password. Write-only: always returned empty.
	AuthPassword string `json:"authPassword"`
	// Signature key which can be used to validate incoming webhook payloads. Only
	// returned on creation and secret rotation.
	Secret string `json:"secret"`
	// Indicates if this webhook is enabled.
	Enabled bool `json:"enabled"`
	// Webhook error logs from the most recent failure.
	Logs string `json:"logs"`
	// Number of consecutive failed webhook attempts.
	Attempts int `json:"attempts"`

	// Used by Decode() method
	data []byte
}

func (model Webhook) New(data []byte) *Webhook {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Webhook) UnmarshalJSON(data []byte) error {
	type alias Webhook
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id           json.RawMessage `json:"$id"`
		CreatedAt    json.RawMessage `json:"$createdAt"`
		UpdatedAt    json.RawMessage `json:"$updatedAt"`
		Name         json.RawMessage `json:"name"`
		Description  json.RawMessage `json:"description"`
		Url          json.RawMessage `json:"url"`
		Events       json.RawMessage `json:"events"`
		Tls          json.RawMessage `json:"tls"`
		AuthUsername json.RawMessage `json:"authUsername"`
		AuthPassword json.RawMessage `json:"authPassword"`
		Secret       json.RawMessage `json:"secret"`
		Enabled      json.RawMessage `json:"enabled"`
		Logs         json.RawMessage `json:"logs"`
		Attempts     json.RawMessage `json:"attempts"`
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
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Description == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Description); err != nil {
			return err
		}
	}
	if present.Url == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Url); err != nil {
			return err
		}
	}
	if present.Events == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Events); err != nil {
			return err
		}
	}
	if present.Tls == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Tls); err != nil {
			return err
		}
	}
	if present.AuthUsername == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AuthUsername); err != nil {
			return err
		}
	}
	if present.AuthPassword == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AuthPassword); err != nil {
			return err
		}
	}
	if present.Secret == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Secret); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.Logs == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Logs); err != nil {
			return err
		}
	}
	if present.Attempts == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Attempts); err != nil {
			return err
		}
	}

	*model = Webhook(decoded)
	return nil
}

func (model *Webhook) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
