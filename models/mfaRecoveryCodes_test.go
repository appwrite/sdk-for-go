package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMfaRecoveryCodesModel(t *testing.T) {
	model := MfaRecoveryCodes{RecoveryCodes: []string{"test"}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result MfaRecoveryCodes
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMfaRecoveryCodesModelDefaults(t *testing.T) {
	var fresh MfaRecoveryCodes
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("[]"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.RecoveryCodes)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing recoveryCodes to default to %v, got %v", want, got)
		}
	}

	reused := MfaRecoveryCodes{RecoveryCodes: []string{"test"}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
