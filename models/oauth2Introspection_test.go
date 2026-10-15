package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOauth2IntrospectionModel(t *testing.T) {
	model := Oauth2Introspection{Active: true}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Oauth2Introspection
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Active != model.Active {
		t.Errorf("Expected Active %v, got %v", model.Active, result.Active)
	}
}

func TestOauth2IntrospectionModelDefaults(t *testing.T) {
	var fresh Oauth2Introspection
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Active)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing active to default to %v, got %v", want, got)
		}
	}

	reused := Oauth2Introspection{Active: true}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
