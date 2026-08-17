package models

import (
	"encoding/json"
	"testing"
)

func TestOauth2TokenModel(t *testing.T) {
	model := Oauth2Token{AccessToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9...", TokenType: "Bearer", ExpiresIn: 3600, RefreshToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9...", Scope: "openid email profile"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Oauth2Token
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.AccessToken != model.AccessToken {
		t.Errorf("Expected AccessToken %v, got %v", model.AccessToken, result.AccessToken)
	}
	if result.TokenType != model.TokenType {
		t.Errorf("Expected TokenType %v, got %v", model.TokenType, result.TokenType)
	}
	if result.ExpiresIn != model.ExpiresIn {
		t.Errorf("Expected ExpiresIn %v, got %v", model.ExpiresIn, result.ExpiresIn)
	}
	if result.RefreshToken != model.RefreshToken {
		t.Errorf("Expected RefreshToken %v, got %v", model.RefreshToken, result.RefreshToken)
	}
	if result.Scope != model.Scope {
		t.Errorf("Expected Scope %v, got %v", model.Scope, result.Scope)
	}
}
