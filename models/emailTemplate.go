package models

import (
	"encoding/json"
	"errors"
)

// EmailTemplate Model
type EmailTemplate struct {
	// Template type
	TemplateId string `json:"templateId"`
	// Template locale
	Locale string `json:"locale"`
	// Template message
	Message string `json:"message"`
	// Name of the sender
	SenderName string `json:"senderName"`
	// Email of the sender
	SenderEmail string `json:"senderEmail"`
	// Reply to email address
	ReplyToEmail string `json:"replyToEmail"`
	// Reply to name
	ReplyToName string `json:"replyToName"`
	// Email subject
	Subject string `json:"subject"`

	// Used by Decode() method
	data []byte
}

func (model EmailTemplate) New(data []byte) *EmailTemplate {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *EmailTemplate) UnmarshalJSON(data []byte) error {
	type alias EmailTemplate
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		TemplateId   json.RawMessage `json:"templateId"`
		Locale       json.RawMessage `json:"locale"`
		Message      json.RawMessage `json:"message"`
		SenderName   json.RawMessage `json:"senderName"`
		SenderEmail  json.RawMessage `json:"senderEmail"`
		ReplyToEmail json.RawMessage `json:"replyToEmail"`
		ReplyToName  json.RawMessage `json:"replyToName"`
		Subject      json.RawMessage `json:"subject"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.TemplateId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.TemplateId); err != nil {
			return err
		}
	}
	if present.Locale == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Locale); err != nil {
			return err
		}
	}
	if present.Message == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Message); err != nil {
			return err
		}
	}
	if present.SenderName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SenderName); err != nil {
			return err
		}
	}
	if present.SenderEmail == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SenderEmail); err != nil {
			return err
		}
	}
	if present.ReplyToEmail == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ReplyToEmail); err != nil {
			return err
		}
	}
	if present.ReplyToName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ReplyToName); err != nil {
			return err
		}
	}
	if present.Subject == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Subject); err != nil {
			return err
		}
	}

	*model = EmailTemplate(decoded)
	return nil
}

func (model *EmailTemplate) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
