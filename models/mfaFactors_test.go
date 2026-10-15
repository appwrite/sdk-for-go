package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMfaFactorsModel(t *testing.T) {
	model := MfaFactors{Totp: true, Phone: true, Email: true, RecoveryCode: true, Custom: true}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result MfaFactors
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Totp != model.Totp {
		t.Errorf("Expected Totp %v, got %v", model.Totp, result.Totp)
	}
	if result.Phone != model.Phone {
		t.Errorf("Expected Phone %v, got %v", model.Phone, result.Phone)
	}
	if result.Email != model.Email {
		t.Errorf("Expected Email %v, got %v", model.Email, result.Email)
	}
	if result.RecoveryCode != model.RecoveryCode {
		t.Errorf("Expected RecoveryCode %v, got %v", model.RecoveryCode, result.RecoveryCode)
	}
	if result.Custom != model.Custom {
		t.Errorf("Expected Custom %v, got %v", model.Custom, result.Custom)
	}
}

func TestMfaFactorsModelDefaults(t *testing.T) {
	var fresh MfaFactors
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Totp)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing totp to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Phone)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing phone to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Email)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing email to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.RecoveryCode)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing recoveryCode to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Custom)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing custom to default to %v, got %v", want, got)
		}
	}

	reused := MfaFactors{Totp: true, Phone: true, Email: true, RecoveryCode: true, Custom: true}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
