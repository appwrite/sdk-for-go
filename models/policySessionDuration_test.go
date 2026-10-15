package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPolicySessionDurationModel(t *testing.T) {
	model := PolicySessionDuration{Id: "password-dictionary", Duration: 3600}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result PolicySessionDuration
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != model.Id {
		t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
	}
	if result.Duration != model.Duration {
		t.Errorf("Expected Duration %v, got %v", model.Duration, result.Duration)
	}
}

func TestPolicySessionDurationModelDefaults(t *testing.T) {
	var fresh PolicySessionDuration
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Id)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing $id to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("31536000"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Duration)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing duration to default to %v, got %v", want, got)
		}
	}

	reused := PolicySessionDuration{Id: "password-dictionary", Duration: 3600}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
