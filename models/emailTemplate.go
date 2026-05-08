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