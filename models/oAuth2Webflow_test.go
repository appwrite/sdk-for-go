package models

import (
	"encoding/json"
	"testing"
)

func TestOAuth2WebflowModel(t *testing.T) {
	model := OAuth2Webflow{Id: "github", Enabled: true, ClientId: "8bb20000000000000000000000000000000000000000000000000000000040dd", ClientSecret: "59bf00000000000000000000000000000000000000000000000000000000fe59"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result OAuth2Webflow
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
	if result.ClientId != model.ClientId {
		t.Errorf("Expected ClientId %v, got %v", model.ClientId, result.ClientId)
	}
	if result.ClientSecret != model.ClientSecret {
		t.Errorf("Expected ClientSecret %v, got %v", model.ClientSecret, result.ClientSecret)
	}
}
