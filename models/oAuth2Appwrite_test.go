package models

import (
	"encoding/json"
	"testing"
)

func TestOAuth2AppwriteModel(t *testing.T) {
	model := OAuth2Appwrite{Id: "github", Enabled: true, ClientId: "6a42000000000000b5a0", ClientSecret: "b86afd000000000000000000000000000000000000000000000000000ced5f93"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result OAuth2Appwrite
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
