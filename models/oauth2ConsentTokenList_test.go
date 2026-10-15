package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOauth2ConsentTokenListModel(t *testing.T) {
	model := Oauth2ConsentTokenList{Total: 5, Tokens: []Oauth2ConsentToken{Oauth2ConsentToken{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", ConsentId: "5e5ea5c16897e", UserId: "5e5ea5c16897e", AppId: "5e5ea5c16897e", CimdUrl: "https://example.com/.well-known/client-metadata.json", Scopes: []string{"test"}, Resources: []string{"test"}, AuthorizationDetails: "[{\"type\":\"calendar\",\"identifier\":\"primary\",\"actions\":[\"read_events\",\"create_event\"]}]", Expire: "2020-10-15T06:38:00.000+00:00"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Oauth2ConsentTokenList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestOauth2ConsentTokenListModelDefaults(t *testing.T) {
	var fresh Oauth2ConsentTokenList
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Total)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing total to default to %v, got %v", want, got)
		}
	}

	reused := Oauth2ConsentTokenList{Total: 5, Tokens: []Oauth2ConsentToken{Oauth2ConsentToken{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", ConsentId: "5e5ea5c16897e", UserId: "5e5ea5c16897e", AppId: "5e5ea5c16897e", CimdUrl: "https://example.com/.well-known/client-metadata.json", Scopes: []string{"test"}, Resources: []string{"test"}, AuthorizationDetails: "[{\"type\":\"calendar\",\"identifier\":\"primary\",\"actions\":[\"read_events\",\"create_event\"]}]", Expire: "2020-10-15T06:38:00.000+00:00"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
