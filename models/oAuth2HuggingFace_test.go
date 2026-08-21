package models

import (
	"encoding/json"
	"testing"
)

func TestOAuth2HuggingFaceModel(t *testing.T) {
	model := OAuth2HuggingFace{Id: "github", Enabled: true, ClientId: "2ab9cff9-d711-40ad-a91e-b08a49c42d24", ClientSecret: "oauth_app_secret_wcLhRtl000000000000000000000xbNdLt"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result OAuth2HuggingFace
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
