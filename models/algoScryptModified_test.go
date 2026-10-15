package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAlgoScryptModifiedModel(t *testing.T) {
	model := AlgoScryptModified{Type: "scryptMod", Salt: "UxLMreBr6tYyjQ==", SaltSeparator: "Bw==", SignerKey: "XyEKE9RcTDeLEsL/RjwPDBv/RqDl8fb3gpYEOQaPihbxf1ZAtSOHCjuAAa7Q3oHpCYhXSN9tizHgVOwn6krflQ=="}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result AlgoScryptModified
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Type != model.Type {
		t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
	}
	if result.Salt != model.Salt {
		t.Errorf("Expected Salt %v, got %v", model.Salt, result.Salt)
	}
	if result.SaltSeparator != model.SaltSeparator {
		t.Errorf("Expected SaltSeparator %v, got %v", model.SaltSeparator, result.SaltSeparator)
	}
	if result.SignerKey != model.SignerKey {
		t.Errorf("Expected SignerKey %v, got %v", model.SignerKey, result.SignerKey)
	}
}

func TestAlgoScryptModifiedModelDefaults(t *testing.T) {
	var fresh AlgoScryptModified
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"scryptMod\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Type)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing type to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Salt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing salt to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SaltSeparator)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing saltSeparator to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SignerKey)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing signerKey to default to %v, got %v", want, got)
		}
	}

	reused := AlgoScryptModified{Type: "scryptMod", Salt: "UxLMreBr6tYyjQ==", SaltSeparator: "Bw==", SignerKey: "XyEKE9RcTDeLEsL/RjwPDBv/RqDl8fb3gpYEOQaPihbxf1ZAtSOHCjuAAa7Q3oHpCYhXSN9tizHgVOwn6krflQ=="}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
