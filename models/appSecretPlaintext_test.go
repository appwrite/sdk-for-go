package models

import (
	"encoding/json"
	"testing"
)

func TestAppSecretPlaintextModel(t *testing.T) {
	model := AppSecretPlaintext{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", AppId: "5e5ea5c16897e", Secret: "5f3c8d2a1b9e4f7a6c8b2d1e9f4a7b3c5d8e1f2a9b4c7d6e3f5a8b1c4d7e2f9a", Hint: "f5c6c7", CreatedById: "5e5ea5c16897e", CreatedByName: "Walter White"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result AppSecretPlaintext
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
	if result.UpdatedAt != model.UpdatedAt {
		t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
	}
	if result.AppId != model.AppId {
		t.Errorf("Expected AppId %v, got %v", model.AppId, result.AppId)
	}
	if result.Secret != model.Secret {
		t.Errorf("Expected Secret %v, got %v", model.Secret, result.Secret)
	}
	if result.Hint != model.Hint {
		t.Errorf("Expected Hint %v, got %v", model.Hint, result.Hint)
	}
	if result.CreatedById != model.CreatedById {
		t.Errorf("Expected CreatedById %v, got %v", model.CreatedById, result.CreatedById)
	}
	if result.CreatedByName != model.CreatedByName {
		t.Errorf("Expected CreatedByName %v, got %v", model.CreatedByName, result.CreatedByName)
	}
}
