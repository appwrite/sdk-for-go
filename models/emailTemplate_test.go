package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestEmailTemplateModel(t *testing.T) {
	model := EmailTemplate{TemplateId: "verification", Locale: "en_us", Message: "Click on the link to verify your account.", SenderName: "My User", SenderEmail: "mail@appwrite.io", ReplyToEmail: "emails@appwrite.io", ReplyToName: "Support Team", Subject: "Please verify your email address"}

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
	}
}

func TestEmailTemplateModelDefaults(t *testing.T) {
	var fresh EmailTemplate
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.TemplateId)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing templateId to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Locale)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing locale to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Message)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing message to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SenderName)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing senderName to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SenderEmail)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing senderEmail to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ReplyToEmail)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing replyToEmail to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ReplyToName)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing replyToName to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Subject)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing subject to default to %v, got %v", want, got)
		}
	}

	reused := EmailTemplate{TemplateId: "verification", Locale: "en_us", Message: "Click on the link to verify your account.", SenderName: "My User", SenderEmail: "mail@appwrite.io", ReplyToEmail: "emails@appwrite.io", ReplyToName: "Support Team", Subject: "Please verify your email address"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
