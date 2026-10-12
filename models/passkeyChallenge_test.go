package models

import (
	"encoding/json"
	"testing"
)

func TestPasskeyChallengeModel(t *testing.T) {
	model := PasskeyChallenge{Id: "bb8ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", PasskeyId: "5e5ea5c16897e", Expire: "2020-10-15T06:38:00.000+00:00", PublicKey: map[string]interface{}{}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result PasskeyChallenge
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != model.Id {
		t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
	}
	if result.CreatedAt != model.CreatedAt {
		t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
	}
	if result.PasskeyId != model.PasskeyId {
		t.Errorf("Expected PasskeyId %v, got %v", model.PasskeyId, result.PasskeyId)
	}
	if result.Expire != model.Expire {
		t.Errorf("Expected Expire %v, got %v", model.Expire, result.Expire)
	}
}
