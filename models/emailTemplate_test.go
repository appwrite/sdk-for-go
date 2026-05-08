package models

import (
    "encoding/json"
    "testing"
)

func TestEmailTemplateModel(t *testing.T) {
    model := EmailTemplate{        TemplateId: "verification",        Locale: "en_us",        Message: "Click on the link to verify your account.",        SenderName: "My User",        SenderEmail: "mail@appwrite.io",        ReplyToEmail: "emails@appwrite.io",        ReplyToName: "Support Team",        Subject: "Please verify your email address",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result EmailTemplate
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.TemplateId != model.TemplateId {
        t.Errorf("Expected TemplateId %v, got %v", model.TemplateId, result.TemplateId)
    }
    if result.Locale != model.Locale {
        t.Errorf("Expected Locale %v, got %v", model.Locale, result.Locale)
    }
    if result.Message != model.Message {
        t.Errorf("Expected Message %v, got %v", model.Message, result.Message)
    }
    if result.SenderName != model.SenderName {
        t.Errorf("Expected SenderName %v, got %v", model.SenderName, result.SenderName)
    }
    if result.SenderEmail != model.SenderEmail {
        t.Errorf("Expected SenderEmail %v, got %v", model.SenderEmail, result.SenderEmail)
    }
    if result.ReplyToEmail != model.ReplyToEmail {
        t.Errorf("Expected ReplyToEmail %v, got %v", model.ReplyToEmail, result.ReplyToEmail)
    }
    if result.ReplyToName != model.ReplyToName {
        t.Errorf("Expected ReplyToName %v, got %v", model.ReplyToName, result.ReplyToName)
    }
    if result.Subject != model.Subject {
        t.Errorf("Expected Subject %v, got %v", model.Subject, result.Subject)
    }}
