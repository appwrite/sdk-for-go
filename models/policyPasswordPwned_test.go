package models

import (
	"encoding/json"
	"testing"
)

func TestPolicyPasswordPwnedModel(t *testing.T) {
	model := PolicyPasswordPwned{Id: "password-dictionary", Enabled: true, Sessions: true, Users: true}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result PolicyPasswordPwned
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != model.Id {
		t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
	}
	if result.Enabled != model.Enabled {
		t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
	}
	if result.Sessions != model.Sessions {
		t.Errorf("Expected Sessions %v, got %v", model.Sessions, result.Sessions)
	}
	if result.Users != model.Users {
		t.Errorf("Expected Users %v, got %v", model.Users, result.Users)
	}
}
