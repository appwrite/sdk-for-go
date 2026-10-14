package models

import (
	"encoding/json"
	"testing"
)

func TestPolicyPasswordRotationModel(t *testing.T) {
	model := PolicyPasswordRotation{Id: "password-dictionary", Enabled: true, MaxAge: 31536000}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result PolicyPasswordRotation
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
	if result.MaxAge != model.MaxAge {
		t.Errorf("Expected MaxAge %v, got %v", model.MaxAge, result.MaxAge)
	}
}
