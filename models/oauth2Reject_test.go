package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOauth2RejectModel(t *testing.T) {
	model := Oauth2Reject{RedirectUrl: "https://example.com/callback?error=access_denied&state=fghij"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Oauth2Reject
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.RedirectUrl != model.RedirectUrl {
		t.Errorf("Expected RedirectUrl %v, got %v", model.RedirectUrl, result.RedirectUrl)
	}
}

func TestOauth2RejectModelDefaults(t *testing.T) {
	var fresh Oauth2Reject
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

	reused := Oauth2Reject{RedirectUrl: "https://example.com/callback?error=access_denied&state=fghij"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
