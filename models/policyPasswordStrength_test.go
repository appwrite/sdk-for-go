package models

import (
	"encoding/json"
	"testing"
)

func TestPolicyPasswordStrengthModel(t *testing.T) {
	model := PolicyPasswordStrength{Id: "password-dictionary", Min: 12, Uppercase: true, Lowercase: true, Number: true, Symbols: true}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result PolicyPasswordStrength
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != model.Id {
		t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
	}
	if result.Min != model.Min {
		t.Errorf("Expected Min %v, got %v", model.Min, result.Min)
	}
	if result.Uppercase != model.Uppercase {
		t.Errorf("Expected Uppercase %v, got %v", model.Uppercase, result.Uppercase)
	}
	if result.Lowercase != model.Lowercase {
		t.Errorf("Expected Lowercase %v, got %v", model.Lowercase, result.Lowercase)
	}
	if result.Number != model.Number {
		t.Errorf("Expected Number %v, got %v", model.Number, result.Number)
	}
	if result.Symbols != model.Symbols {
		t.Errorf("Expected Symbols %v, got %v", model.Symbols, result.Symbols)
	}
}
