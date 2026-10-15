package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOauth2ApproveModel(t *testing.T) {
	model := Oauth2Approve{RedirectUrl: "https://example.com/callback?code=abcde&state=fghij"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Oauth2Approve
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.RedirectUrl != model.RedirectUrl {
		t.Errorf("Expected RedirectUrl %v, got %v", model.RedirectUrl, result.RedirectUrl)
	}
}

func TestOauth2ApproveModelDefaults(t *testing.T) {
	var fresh Oauth2Approve
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.RedirectUrl)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing redirectUrl to default to %v, got %v", want, got)
		}
	}

	reused := Oauth2Approve{RedirectUrl: "https://example.com/callback?code=abcde&state=fghij"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
