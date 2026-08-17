package models

import (
	"encoding/json"
	"testing"
)

func TestOauth2PARModel(t *testing.T) {
	model := Oauth2PAR{RequestUri: "urn:appwrite:oauth2:request:5e5ea5c16897e", ExpiresIn: 600}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Oauth2PAR
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.RequestUri != model.RequestUri {
		t.Errorf("Expected RequestUri %v, got %v", model.RequestUri, result.RequestUri)
	}
	if result.ExpiresIn != model.ExpiresIn {
		t.Errorf("Expected ExpiresIn %v, got %v", model.ExpiresIn, result.ExpiresIn)
	}
}
