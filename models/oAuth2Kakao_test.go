package models

import (
	"encoding/json"
	"testing"
)

func TestOAuth2KakaoModel(t *testing.T) {
	model := OAuth2Kakao{Id: "github", Enabled: true, ClientId: "839ff5000000000000000000013206de", ClientSecret: "jLNVOK00000000000000000000yJebea"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result OAuth2Kakao
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
