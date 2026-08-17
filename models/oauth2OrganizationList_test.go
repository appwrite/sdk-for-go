package models

import (
	"encoding/json"
	"testing"
)

func TestOauth2OrganizationListModel(t *testing.T) {
	model := Oauth2OrganizationList{Total: 5, Organizations: []Oauth2Organization{Oauth2Organization{Id: "5e5ea5c16897e"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Oauth2OrganizationList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}
