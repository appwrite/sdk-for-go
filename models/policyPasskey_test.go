package models

import (
	"encoding/json"
	"testing"
)

func TestPolicyPasskeyModel(t *testing.T) {
	model := PolicyPasskey{Id: "password-dictionary", RpId: "example.com", Origins: []string{"test"}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result PolicyPasskey
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != model.Id {
		t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
	}
	if result.RpId != model.RpId {
		t.Errorf("Expected RpId %v, got %v", model.RpId, result.RpId)
	}
}
