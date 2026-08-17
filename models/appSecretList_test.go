package models

import (
	"encoding/json"
	"testing"
)

func TestAppSecretListModel(t *testing.T) {
	model := AppSecretList{Total: 5, Secrets: []AppSecret{AppSecret{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", AppId: "5e5ea5c16897e", Secret: "string", Hint: "f5c6c7", CreatedById: "5e5ea5c16897e", CreatedByName: "Walter White"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result AppSecretList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}
