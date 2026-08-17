package models

import (
	"encoding/json"
	"testing"
)

func TestOauth2ConsentListModel(t *testing.T) {
	model := Oauth2ConsentList{Total: 5, Consents: []Oauth2Consent{Oauth2Consent{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", UserId: "5e5ea5c16897e", AppId: "5e5ea5c16897e", CimdUrl: "https://example.com/.well-known/client-metadata.json", Scopes: []string{"test"}, Resources: []string{"test"}, AuthorizationDetails: "[{\"type\":\"calendar\",\"identifier\":\"primary\",\"actions\":[\"read_events\",\"create_event\"]}]", Expire: "2020-10-15T06:38:00.000+00:00"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Oauth2ConsentList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}
