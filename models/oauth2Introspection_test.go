package models

import (
	"encoding/json"
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
