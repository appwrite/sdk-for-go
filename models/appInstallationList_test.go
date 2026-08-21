package models

import (
	"encoding/json"
	"testing"
)

func TestAppInstallationListModel(t *testing.T) {
	model := AppInstallationList{Total: 5, Installations: []AppInstallation{AppInstallation{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", AppId: "5e5ea5c16897e", TeamId: "5e5ea5c16897e", Scopes: []string{"test"}, AuthorizationDetails: []interface{}{}, CreatedById: "5e5ea5c16897e", CreatedByName: "Walter White"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result AppInstallationList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}
