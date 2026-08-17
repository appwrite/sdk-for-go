package models

import (
	"encoding/json"
	"testing"
)

func TestOauth2ProjectListModel(t *testing.T) {
	model := Oauth2ProjectList{Total: 5, Projects: []Oauth2Project{Oauth2Project{Id: "5e5ea5c16897e", Region: "fra", Endpoint: "https://fra.cloud.appwrite.io/v1"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Oauth2ProjectList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}
