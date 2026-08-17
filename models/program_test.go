package models

import (
	"encoding/json"
	"testing"
)

func TestProgramModel(t *testing.T) {
	model := Program{Id: "string", Title: "string", Description: "string", Tag: "string", Icon: "string", Url: "string", Active: true, External: true, BillingPlanId: "string"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Program
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != model.Id {
		t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
	}
	if result.Title != model.Title {
		t.Errorf("Expected Title %v, got %v", model.Title, result.Title)
	}
	if result.Description != model.Description {
		t.Errorf("Expected Description %v, got %v", model.Description, result.Description)
	}
	if result.Tag != model.Tag {
		t.Errorf("Expected Tag %v, got %v", model.Tag, result.Tag)
	}
	if result.Icon != model.Icon {
		t.Errorf("Expected Icon %v, got %v", model.Icon, result.Icon)
	}
	if result.Url != model.Url {
		t.Errorf("Expected Url %v, got %v", model.Url, result.Url)
	}
	if result.Active != model.Active {
		t.Errorf("Expected Active %v, got %v", model.Active, result.Active)
	}
	if result.External != model.External {
		t.Errorf("Expected External %v, got %v", model.External, result.External)
	}
	if result.BillingPlanId != model.BillingPlanId {
		t.Errorf("Expected BillingPlanId %v, got %v", model.BillingPlanId, result.BillingPlanId)
	}
}
