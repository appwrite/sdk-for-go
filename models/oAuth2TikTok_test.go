package models

import (
	"encoding/json"
	"testing"
)

func TestOAuth2TikTokModel(t *testing.T) {
	model := OAuth2TikTok{Id: "github", Enabled: true, ClientId: "awz000000000tyw0", ClientSecret: "6wXewM00000000000000000000yXnite"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result OAuth2TikTok
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
