package models

import (
	"encoding/json"
	"testing"
)

func TestOauth2OrganizationModel(t *testing.T) {
	model := Oauth2Organization{Id: "5e5ea5c16897e"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Oauth2Organization
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != model.Id {
		t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
	}
}
