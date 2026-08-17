package models

import (
	"encoding/json"
	"testing"
)

func TestOauth2ProjectModel(t *testing.T) {
	model := Oauth2Project{Id: "5e5ea5c16897e", Region: "fra", Endpoint: "https://fra.cloud.appwrite.io/v1"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Oauth2Project
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != model.Id {
		t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
	}
	if result.Region != model.Region {
		t.Errorf("Expected Region %v, got %v", model.Region, result.Region)
	}
	if result.Endpoint != model.Endpoint {
		t.Errorf("Expected Endpoint %v, got %v", model.Endpoint, result.Endpoint)
	}
}
